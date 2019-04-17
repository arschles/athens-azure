package kube

import (
	"context"

	"github.com/ericchiang/k8s"
)

type Updater interface {
	Update(context.Context, *k8s.Client) error
}
