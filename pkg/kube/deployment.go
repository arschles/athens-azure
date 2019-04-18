package kube

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/ericchiang/k8s"
	appsv1 "github.com/ericchiang/k8s/apis/apps/v1"
	metav1 "github.com/ericchiang/k8s/apis/meta/v1"
)

// Deployment is a convenience wrapper around a k8s deployment object
type Deployment struct {
	core *appsv1.Deployment
	Resource
	fmt.Stringer
}

// NewDeployment creates a new deployment with sensible defaults
func NewDeployment(name, ns string, containers ContainerList) *Deployment {
	return &Deployment{
		core: &appsv1.Deployment{
			Metadata: objectMeta(name, ns),
			Spec: &appsv1.DeploymentSpec{
				Template: podTemplateSpec(name, ns, containers),
				Selector: &metav1.LabelSelector{
					MatchLabels: make(map[string]string),
				},
			},
		},
	}
}

// GetImage returns the image name for the idx'th container
func (d *Deployment) GetImage(idx int) (string, error) {
	con := containerFromPodTemplateSpec(d.core.Spec.Template, idx)
	if con == nil {
		return "", fmt.Errorf("container %d doesn't exist", idx)
	}
	return con.GetImage(), nil
}

// Type is the Typer implementation
func (d *Deployment) Type() string {
	return "Deployment"
}

// Name is the Namer implementation
func (d *Deployment) Name() string {
	return *d.core.Metadata.Name
}

// setReplicas _copies_ d, updates spec.replicas to num on the copy, and
// returns the copy with the updated value
//
// Since this function doesn't copy in place, you'll need to update
// your deployment to the return value of this function
func (d *Deployment) setReplicas(num int32) *Deployment {
	copy := *d
	copy.core.Spec.Replicas = k8s.Int32(num)
	return &copy
}

// setReplicas _copies_ d, updates spec.selector.matchLabels to num
// on the copy, and returns the copy
//
// Since this function doesn't copy in place, you'll need to update
// your deployment to the return value of this function
func (d *Deployment) setMatchLabels(m map[string]string) *Deployment {
	copy := *d
	copy.core.Spec.Selector.MatchLabels = m
	return &copy
}

// Install is the implementation of Installer
func (d *Deployment) Install(ctx context.Context, cl *k8s.Client) error {
	return cl.Create(ctx, d.core)
}

// Update is the implementation of Updater
func (d *Deployment) Update(ctx context.Context, cl *k8s.Client) error {
	return cl.Update(ctx, d.core)
}

// Delete implements Deleter
func (d *Deployment) Delete(ctx context.Context, cl *k8s.Client) error {
	return cl.Delete(ctx, d.core)
}

// Get is the implementation of Getter
func (d *Deployment) Get(ctx context.Context, cl *k8s.Client, name, ns string) error {
	return cl.Get(ctx, ns, name, d.core)
}

// ReadyCh is the ReadyWatcher implementation
func (d *Deployment) ReadyCh(ctx context.Context, cl *k8s.Client) <-chan error {
	// TODO
	ret := make(chan error)
	close(ret)
	return ret
}

// DeletedCh is the DeletedWatcher implementation
func (d *Deployment) DeletedCh(context.Context, *k8s.Client) <-chan error {
	// TODO
	ret := make(chan error)
	close(ret)
	return ret
}

// Namespace is the implementation of Namespacer
func (d *Deployment) Namespace() *Namespace {
	return NewNamespace(*d.core.Metadata.Namespace)
}

func (d *Deployment) String() string {
	b, err := json.Marshal(d.core)
	if err != nil {
		return fmt.Sprintf("error marshaling Deployment %s", d.Name())
	}
	var buf bytes.Buffer
	if err := json.Indent(&buf, b, "", "    "); err != nil {
		return fmt.Sprintf("error indenting JSON for job %s", d.Name())
	}
	return string(buf.Bytes())
}
