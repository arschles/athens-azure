package athens

import (
	"github.com/arschles/athens-azure/pkg/kube"
	"github.com/arschles/athens-azure/pkg/kube/resources"
)

func newProfile(imgs *images) kube.Profile {
	athensContainers := containerList(athensName, imgs.athens)
	athensProfile := kube.NewWebServerProfile(
		athensName,
		athensNS,
		"", // TODO: make sure ingresses don't set up a host if this is empty
		3,
		athensContainers,
	)
	crathensContainers := containerList(crathensName, imgs.crathens)
	crathensProfile := kube.NewLongRunningBatchProfile(
		crathensName,
		crathensNS,
		crathensContainers,
	)

	lathensContainers := containerList(lathensName, imgs.lathens)
	lathensProfile := kube.NewWebServerProfile(
		lathensName,
		lathensNS,
		"athens.azurefd.net",
		3,
		lathensContainers,
	)
	return kube.NewComposedProfile(athensProfile, crathensProfile, lathensProfile)
}

func containerList(name, img string) resources.ContainerList {
	return resources.ContainerList{
		resources.NewContainer(name, img, 3000),
	}
}
