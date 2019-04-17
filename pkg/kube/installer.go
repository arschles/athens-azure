package kube

import (
	"context"

	"github.com/ericchiang/k8s"
)

type Installer interface {
	Install(context.Context, *k8s.Client) error
}
