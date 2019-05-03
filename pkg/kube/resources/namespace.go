package resources

import (
	"context"
	"net/http"

	"github.com/ericchiang/k8s"
	corev1 "github.com/ericchiang/k8s/apis/core/v1"
)

type Namespace struct {
	core *corev1.Namespace
}

func NewNamespace(name string) *Namespace {
	meta := objectMeta(name, name)
	meta.Namespace = nil // namespaces aren't namespaced
	return &Namespace{
		core: &corev1.Namespace{
			Metadata: meta,
		},
	}
}

func (n *Namespace) Name() string {
	return *n.core.Metadata.Name
}

func (n *Namespace) Type() string {
	return "Namespace"
}

func (n *Namespace) Install(ctx context.Context, cl *k8s.Client) error {
	err := cl.Create(ctx, n.core)
	if apiErr := errToAPIErr(err); apiErr != nil {
		// 201 CREATED and 409 CONFLICT means it's already there
		if apiErr.Code != http.StatusCreated &&
			apiErr.Code != http.StatusConflict {
			return apiErr
		}
		return nil
	}
	return err
}

func (n *Namespace) ReadyCh(context.Context, *k8s.Client) <-chan error {
	// TODO
	ret := make(chan error)
	close(ret)
	return ret
}

func (n *Namespace) DeletedCh(context.Context, *k8s.Client) <-chan error {
	// TODO
	ret := make(chan error)
	close(ret)
	return ret
}
