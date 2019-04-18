package kube

import (
	"context"

	"github.com/ericchiang/k8s"
)

type Installer interface {
	Install(context.Context, *k8s.Client) error
}

// Namer gets the current in-memory name of the resource
//
// This might be different from the name of the resource in the cluster
type Namer interface {
	Name() string
}

// Typer gets the current in-memory type of the resource
type Typer interface {
	Type() string
}

type Getter interface {
	Get(context.Context, *k8s.Client, string, string) error
}

type Updater interface {
	Update(context.Context, *k8s.Client) error
}

type Deleter interface {
	Delete(context.Context, *k8s.Client) error
}
type ReadyWatcher interface {
	ReadyCh(context.Context, *k8s.Client) <-chan error
}

type DeletedWatcher interface {
	DeletedCh(context.Context, *k8s.Client) <-chan error
}

type Namespacer interface {
	Namespace() *Namespace
}
