package crathens

import (
	"github.com/arschles/athens-azure/pkg/kube"
)

func crathensJob(cl kube.ContainerList) *kube.Job {
	return kube.NewJob(name, namespace, cl)
}
