package athens

import "github.com/arschles/athens-azure/pkg/kube"

func athensDeployment(img string) *kube.Deployment {
	containerList := kube.ContainerList{
		kube.NewContainer(name, img),
	}
	return kube.NewDeployment(name, namespace, containerList)
}
