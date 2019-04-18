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

func NewService(
	name,
	ns string,
	svcType string,
	selector map[string]string,
	ports []*corev1.ServicePort,
) *Service {
	return &Service{
		core: &corev1.Service{
			Metadata: objectMeta(name, ns),
			Spec: &corev1.ServiceSpec{
				Ports:    ports,
				Type:     k8s.String(svcType),
				Selector: selector,
			},
		},
	}
}

func (s *Service) setType(t string) {
	s.core.Spec.Type = k8s.String(t)
}

func ServiceFromCore(svc *corev1.Service) *Service {
	return &Service{core: svc}
}

func (s *Service) Install(ctx context.Context, cl *k8s.Client) error {
	return cl.Create(ctx, s.core)
}
