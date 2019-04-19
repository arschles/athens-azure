package resources

import (
	"context"
	"fmt"

	"github.com/arschles/athens-azure/pkg/stringer"
	"github.com/ericchiang/k8s"
	appsv1 "github.com/ericchiang/k8s/apis/apps/v1"
	metav1 "github.com/ericchiang/k8s/apis/meta/v1"
)

// Deployment is a convenience wrapper around a k8s deployment object
type Deployment struct {
	core *appsv1.Deployment
	fmt.Stringer
	// Note: Deployments implement Resource, but don't embed it here
	// because the compiler will not give errors if you don't implement
	// a method in Resource when you put it into a profile
}

// NewDeployment creates a new deployment with sensible defaults
func NewDeployment(
	name,
	ns string,
	selectorLabels map[string]string,
	containers ContainerList,
) *Deployment {
	return &Deployment{
		core: &appsv1.Deployment{
			Metadata: objectMetaWithLabels(name, ns, selectorLabels),
			Spec: &appsv1.DeploymentSpec{
				Replicas: k8s.Int32(3),
				Template: podTemplateSpec(selectorLabels, containers),
				Selector: &metav1.LabelSelector{
					MatchLabels: selectorLabels,
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

// WithReplicas _copies_ d, updates spec.replicas to num on the copy, and
// returns the copy with the updated value
//
// Since this function doesn't copy in place, you'll need to update
// your deployment to the return value of this function
func (d *Deployment) WithReplicas(num int32) *Deployment {
	copy := *d
	copy.core.Spec.Replicas = k8s.Int32(num)
	return &copy
}

// setMatchLabels _copies_ d, updates spec.selector.matchLabels to num
// on the copy, and returns the copy
//
// Since this function doesn't copy in place, you'll need to update
// your deployment to the return value of this function
func (d *Deployment) setMatchLabels(m map[string]string) *Deployment {
	copy := *d
	copy.core.Spec.Selector.MatchLabels = m
	return &copy
}

// WithTplMetadataLabels _copies_ d, updates spec.template.metadata.labels to m
// on the copy, and returns the copy
//
// Since this function doesn't copy in place, you'll need to update
// your deployment to the return value of this function
func (d *Deployment) WithTplMetadataLabels(m map[string]string) *Deployment {
	copy := *d
	copy.core.Spec.Template.Metadata.Labels = m
	return &copy
}

// GetTplMetadataLabels returns the value of spec.template.metadata.labels
func (d *Deployment) GetTplMetadataLabels() map[string]string {
	return d.core.Spec.Template.Metadata.Labels
}

// Install is the implementation of Installer
func (d *Deployment) Install(ctx context.Context, cl *k8s.Client) error {
	return cl.Create(ctx, d.core)
}

// Update is the implementation of Updater
func (d *Deployment) Update(ctx context.Context, cl *k8s.Client) error {
	// user should be only able to update:
	//
	// - replicas
	// - name
	// - image
	// - env
	// - healthy HTTP path
	// - ready HTTP path
	// - port

	return nil
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
	return stringer.ToJSON(d.core, d.Type())
}
