package kube

import (
	"context"
	"net/http"

	"github.com/ericchiang/k8s"
	corev1 "github.com/ericchiang/k8s/apis/core/v1"
)

type namespace struct {
	Installer
	core *corev1.Namespace
}

func newNamespace(name string) *namespace {
	return &namespace{
		core: &corev1.Namespace{
			Metadata: objectMeta(name, name),
		},
	}
}

func (n *namespace) Install(ctx context.Context, cl *k8s.Client) error {
	err := cl.Create(ctx, n.core)
	if apiErr := errToAPIErr(err); apiErr != nil {
		// 201 CREATED or 409 CONFLICT means it's already there
		if apiErr.Code != http.StatusCreated ||
			apiErr.Code != http.StatusConflict {
			return apiErr
		}
		return nil
	}
	return err
}

func UpsertNamespace(ctx context.Context, cl *k8s.Client, name string) error {
	ns := newNamespace(name)
	return ns.Install(ctx, cl)
}
