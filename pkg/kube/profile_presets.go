package kube

import (
	"github.com/ericchiang/k8s"
)

func NewLongRunningBatchProfile(j *Job) *Profile {
	return &Profile{
		resources: []Resource{j},
	}
}

func NewWebServerProfile(
	name,
	ns string,
	replicas int32,
	containers ContainerList,
) *Profile {
	svcPorts := containers.toServicePorts()

	depl := NewDeployment(name, ns, containers)

	depl.core.Spec.Replicas = k8s.Int32(replicas)
	depl.core.Spec.Selector.MatchLabels["app"] = name
	// depl.core.Spec.Selector.MatchLabels["release"] = rel

	depl.core.Spec.Template.Metadata.Labels["app"] = name

	svc := NewService(
		name,
		ns,
		"None",
		depl.core.Spec.Template.Metadata.Labels,
		svcPorts,
	)
	// if there were any ports exposed on the container, set the service
	// type to 'ClusterIP' so that it's network accessible from the
	// ingress controller
	if len(svcPorts) > 0 {
		svc.setType("ClusterIP")
	}

	// ing := NewIngress()
	return &Profile{
		resources: []Resource{depl},
	}
}
