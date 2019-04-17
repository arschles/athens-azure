package kube

import (
	corev1 "github.com/ericchiang/k8s/apis/core/v1"
)

func podTemplateSpec(name, ns string, containers ContainerList) *corev1.PodTemplateSpec {
	return &corev1.PodTemplateSpec{
		Metadata: objectMeta(name, ns),
		Spec: &corev1.PodSpec{
			Containers: containers.toCoreList(),
		},
	}
}
