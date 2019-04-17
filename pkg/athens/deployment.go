package athens

import "github.com/arschles/athens-azure/pkg/kube"

func athensDeployment(img string) *kube.Deployment {
	containerList := kube.ContainerList{
		kube.NewContainer("athens", img),
	}
	return kube.NewDeployment("athens", namespace, containerList)
}
