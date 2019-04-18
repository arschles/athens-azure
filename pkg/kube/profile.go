package kube

import (
	"context"
	"fmt"
	"strings"

	"github.com/ericchiang/k8s"
	"github.com/souz9/errlist"
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
type Profile struct {
	resources []Resource
	fmt.Stringer
}

// NewManualProfile creates a new profile out of a list of resources.
// This profile has no presets
//
// TODO: maybe get rid of this in favor of the convenience functions
// above...
func NewManualProfile(resources []Resource) *Profile {
	return &Profile{resources: resources}
}

// Setup prepares Kubernetes to install the profile. This is doing things
// like creating namespaces, etc...
func (p *Profile) Setup(
	ctx context.Context,
	cl *k8s.Client,
	strat ErrorStrategy,
) error {
	errs := []error{}
	for _, res := range p.resources {
		ns := res.Namespace()
		if err := ns.Install(ctx, cl); err != nil {
			// TODO: strategy
			errs = append(errs, err)
		}
	}
	if len(errs) > 0 {
		return errlist.Error(errs)
	}
	return nil
}

func (p *Profile) Install(
	ctx context.Context,
	cl *k8s.Client,
	strat ErrorStrategy,
) error {
	errs := []error{}
	for _, res := range p.resources {
		readyCh := res.ReadyCh(ctx, cl)
		if err := res.Install(ctx, cl); err != nil {
			// TODO: strategy
			errs = append(errs, err)
		}
		if err := chWait(ctx, readyCh); err != nil {
			//TODO: strategy
		}
	}
	if len(errs) > 0 {
		return errlist.Error(errs)
	}
	return nil
}

// Uninstall calls Delete on all resources in the profile, in reverse order
func (p *Profile) Uninstall(
	ctx context.Context,
	cl *k8s.Client,
	strat ErrorStrategy,
) error {
	errs := []error{}
	for i := len(p.resources) - 1; i >= 0; i-- {
		res := p.resources[i]
		deletedCh := res.DeletedCh(ctx, cl)
		if err := res.Delete(ctx, cl); err != nil {
			// TODO: strategy
			errs = append(errs, err)
		}
		if err := chWait(ctx, deletedCh); err != nil {
			//TODO: strategy
		}
	}
	if len(errs) > 0 {
		return errlist.Error(errs)
	}
	return nil
}

func (p *Profile) Update(
	ctx context.Context,
	cl *k8s.Client,
	strat ErrorStrategy,
) error {
	errs := []error{}
	for _, res := range p.resources {
		readyCh := res.ReadyCh(ctx, cl)
		if err := res.Update(ctx, cl); err != nil {
			// TODO: strategy
			errs = append(errs, err)
		}
		if err := chWait(ctx, readyCh); err != nil {
			// TODO: strategy
		}
	}
	if len(errs) > 0 {
		return errlist.Error(errs)
	}
	return nil
}

func (p *Profile) String() string {
	strs := make([]string, len(p.resources))
	for i, res := range p.resources {
		strs[i] = fmt.Sprintf(
			"%s: %s/%s",
			res.Type(),
			res.Namespace().Name(),
			res.Name(),
		)
	}
	return strings.Join(strs, "\n")
}

// AllResources returns all the Kubernetes resources in this profile, in the
// order they're stored
func (p *Profile) AllResources() []Resource {
	return p.resources
}

// SetupAndInstallProfile calls pr.Setup and then pr.Install according to
// strat
func SetupAndInstallProfile(
	ctx context.Context,
	cl *k8s.Client,
	pr *Profile,
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

type ErrorStrategy string

const (
	ErrorStrategyStop     ErrorStrategy = "stop"
	ErrorStrategyRollback ErrorStrategy = "rollback"
	ErrorStrategyContinue ErrorStrategy = "continue"
)

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
