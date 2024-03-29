package kube

import (
	"context"
	"fmt"
	"strings"

	"github.com/arschles/athens-azure/pkg/kube/resources"
	"github.com/ericchiang/k8s"
)

// Profile is a group of kubernetes resources that make up a specific kind of
// app
//
// For example, a web server would have a Deployment, Service and Ingress
// in most cases.
//
// A long-running background job would be just a Job in most cases.
//
// Profiles represent "sensible defaults", and they are represented
// internally as lists of Resources that are installed in order
// and uninstalled in reverse-order
type Profile interface {
	fmt.Stringer
	// Setup prepares Kubernetes to install the profile. This is doing things
	// like creating namespaces, etc...
	Setup(context.Context, *k8s.Client, ErrorStrategy) error
	// AllResources returns all the Kubernetes resources in this profile, in the
	// order they're stored
	AllResources() []resources.Resource
	Install(context.Context, *k8s.Client, ErrorStrategy) error
	// Uninstall calls Delete on all resources in the profile, in reverse order
	Uninstall(context.Context, *k8s.Client, ErrorStrategy) error
	// Update updates every resource in the profile
	Update(context.Context, *k8s.Client, ErrorStrategy) error
	// Status checks all the resources inside the profile and returns nil
	// if everything is installed properly.
	//
	// Otherwise, returns an error(s) indicating what's wrong
	Status(context.Context, *k8s.Client) error
}

type profile struct {
	resources []resources.Resource
}

func (p *profile) Setup(
	ctx context.Context,
	cl *k8s.Client,
	strat ErrorStrategy,
) error {
	return forEachResource(p.resources, strat, func(res resources.Resource) error {
		ns := res.Namespace()
		readyCh := ns.ReadyCh(ctx, cl)
		if err := ns.Install(ctx, cl); err != nil {
			return err
		}
		return chWait(ctx, readyCh)
	})
}

func (p *profile) Install(
	ctx context.Context,
	cl *k8s.Client,
	strat ErrorStrategy,
) error {
	return forEachResource(p.resources, strat, func(res resources.Resource) error {
		readyCh := res.ReadyCh(ctx, cl)
		if err := res.Install(ctx, cl); err != nil {
			return err
		}
		return chWait(ctx, readyCh)
	})
}

// Uninstall calls Delete on all resources in the profile, in reverse order
func (p *profile) Uninstall(
	ctx context.Context,
	cl *k8s.Client,
	strat ErrorStrategy,
) error {
	return forEachResourceReverse(p.resources, strat, func(res resources.Resource) error {
		deletedCh := res.DeletedCh(ctx, cl)
		if err := res.Delete(ctx, cl); err != nil {
			return err
		}
		return chWait(ctx, deletedCh)
	})
}

func (p *profile) Update(
	ctx context.Context,
	cl *k8s.Client,
	strat ErrorStrategy,
) error {
	return forEachResource(
		p.resources,
		strat,
		func(res resources.Resource) error {
			readyCh := res.ReadyCh(ctx, cl)
			if err := res.Update(ctx, cl); err != nil {
				return err
			}
			return chWait(ctx, readyCh)
		},
	)
}

func (p *profile) String() string {
	strs := make([]string, len(p.resources))
	if err := forEachResourceIdx(
		p.resources,
		ErrorStrategyContinue,
		func(i int, res resources.Resource) error {
			strs[i] = fmt.Sprintf(
				"%s: %s/%s",
				res.Type(),
				res.Namespace().Name(),
				res.Name(),
			)
			return nil
		},
	); err != nil {
		// There's not gonna be an error b/c we always return nil from the
		// closure
		return ""
	}
	return strings.Join(strs, "\n")
}

func (p *profile) AllResources() []resources.Resource {
	return p.resources
}

func (p *profile) Status(ctx context.Context, cl *k8s.Client) error {
	return forEachResource(
		p.resources,
		ErrorStrategyContinue,
		func(res resources.Resource) error {
			return res.Get(ctx, cl, res.Name(), res.Namespace().Name())
		},
	)
}

// SetupAndInstallProfile calls pr.Setup and then pr.Install according to
// strat
func SetupAndInstallProfile(
	ctx context.Context,
	cl *k8s.Client,
	pr Profile,
	strat ErrorStrategy,
) error {
	// TODO: error strategy
	if err := pr.Setup(ctx, cl, strat); err != nil {
		return err
	}
	if err := pr.Install(ctx, cl, strat); err != nil {
		return err
	}
	return nil
}

func chWait(ctx context.Context, ch <-chan error) error {
	select {
	case <-ctx.Done():
		return fmt.Errorf("Timeout waiting for channel")
	case err := <-ch:
		if err != nil {
			return fmt.Errorf("Error returned from wait channel (%s)", err)
		}
		return nil
	}
}
