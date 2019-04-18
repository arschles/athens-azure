package kube

// NewLongRunningBatchProfile creates a new profile that contains all the
// kubernetes resources you need to launch or update a long running batch
// job
func NewLongRunningBatchProfile(j *Job) *Profile {
	return &Profile{
		resources: []Resource{j},
	}
}

// NewWebServerProfile creates a new profile that contains all the kubernetes
// resources you need to launch or update a web server
func NewWebServerProfile(
	name,
	ns string,
	replicas int32,
	containers ContainerList,
) *Profile {
	svcPorts := containers.toServicePorts()

	depl := NewDeployment(name, ns, containers)

	depl = depl.setReplicas(replicas)
	depl = depl.setMatchLabels(map[string]string{
		"app": name,
		// "release": rel
	})

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
		svc = svc.setType("ClusterIP")
	}

	// ing := NewIngress()
	return &Profile{
		resources: []Resource{depl},
	}
}
