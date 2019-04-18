package kube

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
