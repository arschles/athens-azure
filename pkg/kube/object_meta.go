package kube

import (
	"github.com/ericchiang/k8s"
	metav1 "github.com/ericchiang/k8s/apis/meta/v1"
)

func objectMeta(name, ns string) *metav1.ObjectMeta {
	return &metav1.ObjectMeta{
		Name:      k8s.String(name),
		Namespace: k8s.String(ns),
		Labels:    make(map[string]string),
	}
}
