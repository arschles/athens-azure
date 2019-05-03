package athens

import (
	"github.com/arschles/athens-azure/pkg/conf"
	"github.com/arschles/athens-azure/pkg/kube"
	"github.com/arschles/athens-azure/pkg/kube/resources"
)

func newProfile(webConfs []conf.Web, jobConfs []conf.Job) kube.Profile {
	profiles := []kube.Profile{}
	for _, webConf := range webConfs {
		containers := containerList(webConf.Name, webConf.Image, webConf.Port)
		profile := kube.NewWebServerProfile(
			webConf.Name,
			webConf.Name,
			"", // TODO: make sure ingresses don't set up a host if this is empty
			webConf.Replicas,
			containers,
		)
		profiles = append(profiles, profile)
	}

	for _, jobConf := range jobConfs {
		containers := containerList(jobConf.Name, jobConf.Image, -1)
		profile := kube.NewLongRunningBatchProfile(
			jobConf.Name,
			jobConf.Name,
			containers,
		)
		profiles = append(profiles, profile)
	}
	return kube.NewComposedProfile(profiles...)
}

func containerList(name, img string, port int32) resources.ContainerList {
	return resources.ContainerList{
		resources.NewContainer(name, img, 3000),
	}
}
