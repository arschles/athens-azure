package athens

import (
	"github.com/arschles/athens-azure/pkg/kube"
	"github.com/arschles/athens-azure/pkg/kube/resources"
)

func newProfile(img string) kube.Profile {
	return kube.NewWebServerProfile(
		name,
		namespace,
		"athens.azurefd.net",
		3,
		containerList(img),
	)
}

func containerList(img string) resources.ContainerList {
	return resources.ContainerList{
		resources.NewContainer(name, img, 3000),
	}
}
