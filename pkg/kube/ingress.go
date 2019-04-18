package kube

import (
	"context"
	"fmt"

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

// IngressFromCore creates a new ingress from a core ingress value
func IngressFromCore(ing *extv1beta1.Ingress) *Ingress {
	return &Ingress{
		core: ing,
	}
}

func (i *Ingress) Install(ctx context.Context, cl *k8s.Client) error {
	return cl.Create(ctx, i.core)
}
