package kube

import (
	"context"
	"fmt"

	"github.com/ericchiang/k8s"
	appsv1 "github.com/ericchiang/k8s/apis/apps/v1"
)

// Deployment is a convenience wrapper around a k8s deployment object
type Deployment struct {
	core *appsv1.Deployment
	Installer
	Updater
	Getter
}

func NewDeployment(name, ns string, containers ContainerList) *Deployment {
	return &Deployment{
		core: &appsv1.Deployment{
			Metadata: objectMeta(name, ns),
			Spec: &appsv1.DeploymentSpec{
				Template: podTemplateSpec(name, ns, containers),
			},
		},
	}
}

func (d *Deployment) GetImage(idx int) (string, error) {
	con := containerFromPodTemplateSpec(d.core.Spec.Template, idx)
	if con == nil {
		return "", fmt.Errorf("container %d doesn't exist", idx)
	}
	return con.GetImage(), nil
}

func (d *Deployment) Install(ctx context.Context, cl *k8s.Client) error {
	return cl.Create(ctx, d.core)
}

func (d *Deployment) Update(ctx context.Context, cl *k8s.Client) error {
	return cl.Update(ctx, d.core)
}

func (d *Deployment) Get(ctx context.Context, cl *k8s.Client, name, ns string) error {
	return cl.Get(ctx, ns, name, d.core)
}
