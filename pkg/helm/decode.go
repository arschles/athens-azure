package helm

// this is a work in progress to decode helm charts into kube convenience
// structures. I don't think it's necessary right now

// import (
// 	appsv1 "github.com/ericchiang/k8s/apis/apps/v1"
// 	corev1 "github.com/ericchiang/k8s/apis/core/v1"
// 	extv1beta1 "github.com/ericchiang/k8s/apis/extensions/v1beta1"
// )

// // render and decode all of the yaml files from the helm chart
// //
// // ... then deploy them to Kubernetes

// const valuesFile = "values.yaml"
// const deploymentFile = "templates/deployment.yaml"
// const serviceFile = "templates/service.yaml"
// const ingressFile = "templates/ingress.yaml"

// type decoded struct {
// 	deployment *appsv1.Deployment
// 	service    *corev1.Service
// 	ingress    *extv1beta1.Ingress
// 	vals       map[string]interface{}
// }

// func Decode(
// 	chartLoc string,
// 	values map[string]interface{},
// ) (*Decoded, error) {
// 	depl := new(appsv1.Deployment)
// 	svc := new(appsv1.Service)
// 	ing := new(extv1beta1.Ingress)
// 	return &Decoded{
// 		Deployment: kube.DeploymentFromCore(depl),
// 		Service:    kube.ServiceFromCore(svc),
// 		Ingress:    kube.IngressFromCore(ing),
// 	}, nil

// 	return nil, nil
// }

// type Decoded struct {
// 	Deployment *kube.Deployment
// 	Service    *kube.Service
// 	Ingress    *kube.Ingress
// }
