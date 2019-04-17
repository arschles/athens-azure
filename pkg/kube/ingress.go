package kube

import (
	"context"

	"github.com/ericchiang/k8s"
	extv1beta1 "github.com/ericchiang/k8s/apis/extensions/v1beta1"
)

type Ingress struct {
	Installer
	core *extv1beta1.Ingress
}

func IngressFromCore(ing *extv1beta1.Ingress) *Ingress {
	return &Ingress{
		core: ing,
	}
}

func (i *Ingress) Install(ctx context.Context, cl *k8s.Client) error {
	return cl.Create(ctx, i.core)
}
