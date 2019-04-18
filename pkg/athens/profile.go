package athens

import "github.com/arschles/athens-azure/pkg/kube"

func newProfile(img string) kube.Profile {
	return kube.NewWebServerProfile(
		name,
		namespace,
		"athens.azurefd.net",
		3,
		containerList(img),
	)
}

func athensDeployment(img string) *kube.Deployment {
	clist := containerList(img)
	return kube.NewDeployment(name, namespace, map[string]string{
		"app": "athens",
	}, clist)
}

func containerList(img string) kube.ContainerList {
	return kube.ContainerList{
		kube.NewContainer(name, img, 3000),
	}
}
