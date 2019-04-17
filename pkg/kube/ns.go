package kube

import (
	"context"
	"net/http"

	"github.com/ericchiang/k8s"
	corev1 "github.com/ericchiang/k8s/apis/core/v1"
	"github.com/pkg/errors"
)

type namespace struct {
	Installer
	core *corev1.Namespace
}

func newNamespace(name string) *namespace {
	meta := objectMeta(name, name)
	meta.Namespace = nil // namespaces aren't namespaced
	return &namespace{
		core: &corev1.Namespace{
			Metadata: meta,
		},
	}
}

func (n *namespace) Install(ctx context.Context, cl *k8s.Client) error {
	err := cl.Create(ctx, n.core)
	if apiErr := errToAPIErr(err); apiErr != nil {
		// 201 CREATED and 409 CONFLICT means it's already there
		if apiErr.Code != http.StatusCreated &&
			apiErr.Code != http.StatusConflict {
			return errors.WithStack(apiErr)
		}
		return nil
	}
	return errors.WithStack(err)
}

func UpsertNamespace(ctx context.Context, cl *k8s.Client, name string) error {
	ns := newNamespace(name)
	return ns.Install(ctx, cl)
}
