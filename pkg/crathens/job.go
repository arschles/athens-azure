package crathens

import (
	"github.com/arschles/athens-azure/pkg/kube/resources"
)

func crathensJob(cl resources.ContainerList) *resources.Job {
	return resources.NewJob(name, namespace, cl)
}
