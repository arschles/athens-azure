package kube

import "github.com/arschles/athens-azure/pkg/kube/resources"

// NewLongRunningBatchProfile creates a new profile that contains all the
// kubernetes resources you need to launch or update a long running batch
// job
func NewLongRunningBatchProfile(
	name,
	ns string,
	cl resources.ContainerList,
) Profile {
	job := resources.NewJob(name, ns, cl)
	return &profile{
		resources: []resources.Resource{job},
	}
}

// NewWebServerProfile creates a new profile that contains all the kubernetes
// resources you need to launch or update a web server
func NewWebServerProfile(
	name,
	ns,
	host string,
	replicas int32,
	containers resources.ContainerList,
) Profile {
	res := []resources.Resource{}

	svcPorts := containers.GetServicePorts()

	// set up the deployment
	depl := resources.NewDeployment(name, ns, map[string]string{
		"app": name,
	}, containers, []resources.DeploymentFieldUpdater{
		resources.NewContainerFieldUpdater(),
	})

	depl = depl.WithReplicas(replicas)
	// depl = depl.setMatchLabels(map[string]string{
	// 	"app": name,
	// 	// "release": rel
	// })

	depl = depl.WithTplMetadataLabels(map[string]string{
		"app": name,
	})
	res = append(res, depl)

	// set up the service
	svc := resources.NewService(
		name,
		ns,
		"None",
		// depl.core.Spec.Template.Metadata.Labels,
		depl.GetTplMetadataLabels(),
		svcPorts,
	)
	// if there were any ports exposed on the container, set the service
	// type to 'ClusterIP' so that it's network accessible from the
	// ingress controller
	if len(svcPorts) > 0 {
		svc = svc.WithType("ClusterIP")
	}
	res = append(res, svc)

	// set up the ingress only if there were ports exposed on the service
	// only set up an ingres
	if len(svcPorts) > 0 {
		svcName := svc.Name()
		svcPort := svcPorts[0].Port
		ing := resources.NewIngress(name, ns, host, "/", svcName, *svcPort)
		res = append(res, ing)
	}
	return &profile{
		resources: res,
	}
}
