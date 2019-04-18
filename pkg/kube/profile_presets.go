package kube

// NewLongRunningBatchProfile creates a new profile that contains all the
// kubernetes resources you need to launch or update a long running batch
// job
func NewLongRunningBatchProfile(j *Job) Profile {
	return &profile{
		resources: []Resource{j},
	}
}

// NewWebServerProfile creates a new profile that contains all the kubernetes
// resources you need to launch or update a web server
func NewWebServerProfile(
	name,
	ns,
	host string,
	replicas int32,
	containers ContainerList,
) Profile {
	res := []Resource{}

	svcPorts := containers.toServicePorts()

	// set up the deployment
	depl := NewDeployment(name, ns, map[string]string{
		"app": name,
	}, containers)

	depl = depl.setReplicas(replicas)
	// depl = depl.setMatchLabels(map[string]string{
	// 	"app": name,
	// 	// "release": rel
	// })

	depl = depl.setTplMetadataLabels(map[string]string{
		"app": name,
	})
	res = append(res, depl)

	// set up the service
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
	res = append(res, svc)

	// set up the ingress only if there were ports exposed on the service
	// only set up an ingres
	if len(svcPorts) > 0 {
		svcName := svc.Name()
		svcPort := svcPorts[0]
		ing := NewIngress(name, ns, host, "/", svcName, *svcPort)
		res = append(res, ing)
	}
	return &profile{
		resources: res,
	}
}
