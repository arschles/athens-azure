package kube

import (
	"context"

	"github.com/ericchiang/k8s"
)

type Getter interface {
	Get(context.Context, *k8s.Client, string, string) error
}
