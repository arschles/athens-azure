package resources

import (
	"context"
	"fmt"

	"github.com/arschles/athens-azure/pkg/stringer"
	"github.com/ericchiang/k8s"
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
	svcPort int32,
) *Ingress {
	rule := newIngressRule(host, path, svcName, svcPort)
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

// gets the first host in the list of ingress rules, or an error if there
// are no ingress rules
func (i *Ingress) getHost() (string, error) {
	rules := i.core.Spec.Rules
	if len(rules) <= 0 {
		return "", fmt.Errorf("No ingress rules are listed, so no host")
	}
	return *rules[0].Host, nil
}

// func (i *Ingress) setHost(h string)*Ingress {
// 	copy := *i
// 	copyPtr := &copy
// 	rules := copyPtr.core.Spec.Rules
// 	if len(rules) < 1 {
// 		newRule := newIngressRule(h, copyPtr.getPath())
// 		copyPtr.core.Spec.Rules = []*extv1beta1.IngressRule{
// 					rule,
// 				},
// 	}
// 	rules := i.core.Spec.Rules
// 	if len(rules) < 1 {
// 		i.core.Spec.Rules
// 	}
// }

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
	name := i.Name()
	ns := i.Namespace().Name()
	copy := *i
	copyPtr := &copy
	if err := copyPtr.Get(ctx, cl, name, ns); err != nil {
		return err
	}
	origHost, err := i.getHost()
	if err != nil {
		return err
	}
	newHost, err := copyPtr.getHost()
	if err != nil {
		return err
	}
	// no change, so bail
	if origHost == newHost {
		return nil
	}
	// TODO: set and update the host
	// if err := copyPtr.setHost(newHost).Update(ctx, cl); err != nil {
	// 	return err
	// }

	// user should only be able to update the domain
	return nil
}

func (i *Ingress) String() string {
	return stringer.ToJSON(i.core, i.Type())
}
