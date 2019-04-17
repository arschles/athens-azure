package kube

import (
	"context"

	"github.com/ericchiang/k8s"
	corev1 "github.com/ericchiang/k8s/apis/core/v1"
)

// Service is a convenience wrapper around a k8s deployment object
type Service struct {
	core *corev1.Service
	Installer
}

func ServiceFromCore(svc *corev1.Service) *Service {
	return &Service{core: svc}
}

func (s *Service) Install(ctx context.Context, cl *k8s.Client) error {
	return cl.Create(ctx, s.core)
}
