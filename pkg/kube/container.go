package kube

import (
	"github.com/ericchiang/k8s"
	corev1 "github.com/ericchiang/k8s/apis/core/v1"
)

type Container struct {
	*corev1.Container
}

func NewContainer(name, img string) *Container {
	return &Container{
		Container: &corev1.Container{
			Name:  k8s.String(name),
			Image: k8s.String(img),
		},
	}
}

func (c *Container) toCore() *corev1.Container {
	return c.Container
}

type ContainerList []*Container

func (c ContainerList) toCoreList() []*corev1.Container {
	ret := make([]*corev1.Container, len(c))
	for i, ctr := range c {
		ret[i] = ctr.toCore()
	}
	return ret
}
