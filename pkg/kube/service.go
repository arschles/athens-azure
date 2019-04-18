package kube

import (
	"context"

	"github.com/arschles/athens-azure/pkg/stringer"
	"github.com/ericchiang/k8s"
	corev1 "github.com/ericchiang/k8s/apis/core/v1"
)

// Service is a convenience wrapper around a k8s deployment object.
// It implements Resource
type Service struct {
	core *corev1.Service
	// Note: Services implement Resource, but don't embed it here
	// because the compiler will not give errors if you don't implement
	// a method in Resource when you put it into a profile
}

// NewService creates a new Service with sensible defaults
func NewService(
	name,
	ns string,
	svcType string,
	selector map[string]string,
	ports []*corev1.ServicePort,
) *Service {
	return &Service{
		core: &corev1.Service{
			Metadata: objectMeta(name, ns),
			Spec: &corev1.ServiceSpec{
				Ports:    ports,
				Type:     k8s.String(svcType),
				Selector: selector,
			},
		},
	}
}

// setType _copies_ s, updates the type of the copied service, and
// returns the copy
//
// Since this function doesn't copy in place, you'll need to update
// your service to the return value of this function
func (s *Service) setType(t string) *Service {
	copy := *s
	copy.core.Spec.Type = k8s.String(t)
	return &copy
}

// ServiceFromCore creates a new service from a core service
func ServiceFromCore(svc *corev1.Service) *Service {
	return &Service{core: svc}
}

// Install implements Installer
func (s *Service) Install(ctx context.Context, cl *k8s.Client) error {
	return cl.Create(ctx, s.core)
}

// Delete implements Deleter
func (s *Service) Delete(ctx context.Context, cl *k8s.Client) error {
	return cl.Delete(ctx, s.core)
}

// Name implements Namer
func (s *Service) Name() string {
	return *s.core.Metadata.Name
}

// Namespace implements Namespacer
func (s *Service) Namespace() *Namespace {
	return NewNamespace(*s.core.Metadata.Namespace)
}

// DeletedCh implements DeletedWatcher
func (s *Service) DeletedCh(ctx context.Context, cl *k8s.Client) <-chan error {
	ret := make(chan error)
	close(ret)
	return ret
}

// ReadyCh implements ReadyWatcher
func (s *Service) ReadyCh(ctx context.Context, cl *k8s.Client) <-chan error {
	ret := make(chan error)
	close(ret)
	return ret
}

// Get implements Getter
func (s *Service) Get(ctx context.Context, cl *k8s.Client, name, ns string) error {
	return cl.Get(ctx, ns, name, s.core)
}

// Type implements Typer
func (s *Service) Type() string {
	return "Service"
}

// Update is the implementation of Updater
func (s *Service) Update(ctx context.Context, cl *k8s.Client) error {
	return cl.Update(ctx, s.core)
}

func (s *Service) String() string {
	return stringer.ToJSON(s.core, s.Type())
}
