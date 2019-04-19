package resources

import (
	"github.com/ericchiang/k8s"
	extv1beta1 "github.com/ericchiang/k8s/apis/extensions/v1beta1"
)

func newIngressRule(
	host,
	path,
	svcName string,
	svcPort int32,
) *extv1beta1.IngressRule {
	return &extv1beta1.IngressRule{
		Host: k8s.String(host),
		IngressRuleValue: &extv1beta1.IngressRuleValue{
			Http: &extv1beta1.HTTPIngressRuleValue{
				Paths: []*extv1beta1.HTTPIngressPath{
					&extv1beta1.HTTPIngressPath{
						Path: k8s.String("/"),
						Backend: &extv1beta1.IngressBackend{
							ServiceName: k8s.String(svcName),
							ServicePort: newIntOrString(svcPort),
						},
					},
				},
			},
		},
	}
}
