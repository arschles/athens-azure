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

// setType _copies_ s, updates the type of the copied service, and
// returns the copy
//
// Since this function doesn't copy in place, you'll need to update
// your service to the return value of this function
func (s *Service) setType(t string) *Service {
	copy := *s
	copy.core.Spec.Type = k8s.String(t)
	return &copy
}

func ServiceFromCore(svc *corev1.Service) *Service {
	return &Service{core: svc}
}

func (s *Service) Install(ctx context.Context, cl *k8s.Client) error {
	return cl.Create(ctx, s.core)
}
