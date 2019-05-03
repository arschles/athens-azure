package resources

// DeploymentFieldUpdater knows how to update a part of a deployment
type DeploymentFieldUpdater interface {
	Update(*Deployment, *Deployment) *Deployment
}

type containerUpdater struct {
}

// NewContainerFieldUpdater returns a DeploymentUpdater that knows
// only how to update the first container of any given deployment to
// newContainer
func NewContainerFieldUpdater() DeploymentFieldUpdater {
	return &containerUpdater{}
}

func (c *containerUpdater) Update(origDepl *Deployment, newDepl *Deployment) *Deployment {
	ret := *origDepl

	ret.core.Spec.Template.Spec.Containers = newDepl.core.Spec.Template.Spec.Containers
	return &ret
}

// note for later: user should be only able to update:
//
// - replicas
// - name
// - images
// - env
// - healthy HTTP path
// - ready HTTP path
// - port
