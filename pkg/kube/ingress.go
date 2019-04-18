package kube

import (
	"context"
	"fmt"

	"github.com/arschles/athens-azure/pkg/stringer"
	"github.com/ericchiang/k8s"
	corev1 "github.com/ericchiang/k8s/apis/core/v1"
	extv1beta1 "github.com/ericchiang/k8s/apis/extensions/v1beta1"
)

// Ingress is a convenience wrapper around the Kubernetes Ingress resource
type Ingress struct {
	fmt.Stringer
	core *extv1beta1.Ingress
	// Note: Ingresses implement Resource, but don't embed it here
	// because the compiler will not give errors if you don't implement
	// a method in Resource when you put it into a profile
}

// NewIngress creates a new ingress from some values with sensible defaults
//
// TODO: this is a mess, clean it up by factoring out some struct
func NewIngress(
	name,
	ns,
	host,
	path,
	svcName string,
	svcPort corev1.ServicePort,
) *Ingress {
	rule := &extv1beta1.IngressRule{
		Host: k8s.String(host),
		IngressRuleValue: &extv1beta1.IngressRuleValue{
			Http: &extv1beta1.HTTPIngressRuleValue{
				Paths: []*extv1beta1.HTTPIngressPath{
					&extv1beta1.HTTPIngressPath{
						Path: k8s.String("/"),
						Backend: &extv1beta1.IngressBackend{
							ServiceName: k8s.String(svcName),
							ServicePort: newIntOrString(*svcPort.Port),
						},
					},
				},
			},
		},
	}
	return &Ingress{
		core: &extv1beta1.Ingress{
			Metadata: objectMeta(name, ns),
			Spec: &extv1beta1.IngressSpec{
				Rules: []*extv1beta1.IngressRule{
					rule,
				},
			},
		},
	}
}

// Install implements Installer
func (i *Ingress) Install(ctx context.Context, cl *k8s.Client) error {
	return cl.Create(ctx, i.core)
}

// Delete implements Deleter
func (i *Ingress) Delete(ctx context.Context, cl *k8s.Client) error {
	return cl.Delete(ctx, i.core)
}

// DeletedCh implements DeletedWatcher
func (i *Ingress) DeletedCh(ctx context.Context, cl *k8s.Client) <-chan error {
	// TODO
	ret := make(chan error)
	close(ret)
	return ret
}

// ReadyCh implements ReadyWatcher
func (i *Ingress) ReadyCh(context.Context, *k8s.Client) <-chan error {
	// TODO
	ret := make(chan error)
	close(ret)
	return ret
}

// Get implemented Getter
func (i *Ingress) Get(ctx context.Context, cl *k8s.Client, name, ns string) error {
	return cl.Get(ctx, ns, name, i.core)
}

// Name implements Namer
func (i *Ingress) Name() string {
	return *i.core.Metadata.Name
}

// Namespace implements Namespacer
func (i *Ingress) Namespace() *Namespace {
	return NewNamespace(*i.core.Metadata.Namespace)
}

// Type implements Typer
func (i *Ingress) Type() string {
	return "Ingress"
}

// Update implements Updater
func (i *Ingress) Update(ctx context.Context, cl *k8s.Client) error {
	return cl.Update(ctx, i.core)
}

func (i *Ingress) String() string {
	return stringer.ToJSON(i.core, i.Type())
}
