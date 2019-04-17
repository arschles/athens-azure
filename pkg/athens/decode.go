package athens

import (
	appsv1 "github.com/ericchiang/k8s/apis/apps/v1"
	corev1 "github.com/ericchiang/k8s/apis/core/v1"
)

// render and decode all of the yaml files from the helm chart
//
// ... then deploy them to Kubernetes

const valuesFile = "values.yaml"
const deploymentFile = "templates/deployment.yaml"
const serviceFile = "templates/service.yaml"
const ingressFile = "templates/ingress.yaml"

type decoded struct {
	deployment *appsv1.Deployment
	service    *corev1.Service
	// more
}
