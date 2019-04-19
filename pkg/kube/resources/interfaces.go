package resources

import (
	"context"

	"github.com/ericchiang/k8s"
)

// Resource is a single Kubernetes object that you can do standard CRUD
// operations on
type Resource interface {
	Installer
	Updater
	Getter
	Deleter
	ReadyWatcher
	DeletedWatcher
	Namespacer
	Namer
	Typer
}

// Installer installs the current in-memory resource into the cluster
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

// Getter gets the resource from the cluster and writes it into the
// local in-memory copy. It takes a name and a namespace (in that order)
// for the in-cluster resource
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
