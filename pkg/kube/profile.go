package kube

import (
	"context"
	"fmt"

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
}

func NewLongRunningBatchProfile(j *Job) *Profile {
	return &Profile{
		resources: []Resource{j},
	}
}

func NewWebServerProfile(name, ns string, containers ContainerList) *Profile {
	depl := NewDeployment(name, ns, containers)
	return &Profile{
		resources: []Resource{depl},
	}
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

// func (c *Crudder) Read(ctx context.Context) error {
// 	return c.Resource.Get(ctx, c.Client)
// }

// Resource is a single Kubernetes object that you can do standard CRUD
// operations on
type Resource interface {
	Installer
	Updater
	Getter
	Deleter
	ReadyWatcher
	DeletedWatcher
}

type ErrorStrategy string

const (
	ErrorStrategyStop     ErrorStrategy = "stop"
	ErrorStrategyRollback ErrorStrategy = "rollback"
	ErrorStrategyIgnore   ErrorStrategy = "ignore"
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
