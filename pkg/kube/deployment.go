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

// SetImage sets the image for the containerNum image in the deployment
func (d *Deployment) SetImage(containerNum int, img string) error {
	numContainers := len(d.core.Spec.Template.Spec.Containers)
	if numContainers < containerNum {
		return fmt.Errorf(
			"Only %d container(s) in the deployment, asked to set #%s",
			numContainers,
			containerNum,
		)
	}
	d.core.Spec.Template.Spec.Containers[containerNum].Image = k8s.String(img)
	return nil
}

func (d *Deployment) Install(ctx context.Context, cl *k8s.Client) error {
	return cl.Create(ctx, d.core)
}

func (d *Deployment) Update(ctx context.Context, cl *k8s.Client) error {
	return cl.Update(ctx, d.core)
}
