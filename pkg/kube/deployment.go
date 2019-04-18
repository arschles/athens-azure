package kube

import (
	"context"
	"fmt"

	"github.com/ericchiang/k8s"
	appsv1 "github.com/ericchiang/k8s/apis/apps/v1"
	corev1 "github.com/ericchiang/k8s/apis/core/v1"
)

// Deployment is a convenience wrapper around a k8s deployment object
type Deployment struct {
	core *appsv1.Deployment
	Installer
	Updater
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
	containers := d.containers()
	if len(containers) < idx+1 {
		return "", fmt.Errorf(
			"requested container %d, but there were only %d",
			idx+1,
			len(containers),
		)
	}
	return containers[idx].GetImage(), nil
}

func (d *Deployment) containers() []*corev1.Container {
	return d.core.Spec.Template.Spec.Containers
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
