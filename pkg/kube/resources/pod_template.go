package resources

import (
	corev1 "github.com/ericchiang/k8s/apis/core/v1"
)

func podTemplateSpec(
	labels map[string]string,
	containers ContainerList,
) *corev1.PodTemplateSpec {
	return &corev1.PodTemplateSpec{
		Metadata: objectMetaWithLabels(emptyStr, emptyStr, labels),
		Spec: &corev1.PodSpec{
			Containers: containers.toCoreList(),
		},
	}
}
