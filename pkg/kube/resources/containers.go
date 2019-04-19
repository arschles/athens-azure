package resources

import (
	corev1 "github.com/ericchiang/k8s/apis/core/v1"
)

func containersFromPodTemplateSpec(pts *corev1.PodTemplateSpec) []*corev1.Container {
	return pts.Spec.Containers
}

func containerFromPodTemplateSpec(pts *corev1.PodTemplateSpec, idx int) *corev1.Container {
	containers := containersFromPodTemplateSpec(pts)
	if len(containers) < idx+1 {
		return nil
	}
	return containers[idx]
}

// ContainerList is a convenience wrapper around a list of Containers
type ContainerList []*Container

func (c ContainerList) toCoreList() []*corev1.Container {
	ret := make([]*corev1.Container, len(c))
	for i, ctr := range c {
		ret[i] = ctr.toCore()
	}
	return ret
}

// GetServicePorts returns a list of the core ServicePorts that correspond
// to all the container ports on c. This is similar to GetServicePorts on
// the Container type, except it aggregates all of the service ports from
// each container into one flat list
func (c ContainerList) GetServicePorts() []*corev1.ServicePort {
	ret := []*corev1.ServicePort{}
	for _, ctr := range c {
		ret = append(ret, ctr.GetServicePorts()...)
	}
	return ret
}
