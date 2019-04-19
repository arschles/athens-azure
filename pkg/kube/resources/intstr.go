package resources

import (
	"github.com/ericchiang/k8s"
	intstr "github.com/ericchiang/k8s/util/intstr"
)

func newIntOrString(i int32) *intstr.IntOrString {
	return &intstr.IntOrString{
		IntVal: k8s.Int32(i),
	}
}
