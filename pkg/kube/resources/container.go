package resources

import (
	"fmt"

	"github.com/ericchiang/k8s"
	corev1 "github.com/ericchiang/k8s/apis/core/v1"
)

type Container struct {
	core *corev1.Container
}

// NewContainer returns a new Container object from the given name and image.
//
// The container will also have a port specified on it if port > 0
func NewContainer(name, image string, port int32) *Container {
	ctr := &Container{
		core: &corev1.Container{
			Name:  k8s.String(name),
			Image: k8s.String(image),
		},
	}

	if port > 0 {
		ctr.core.Ports = []*corev1.ContainerPort{
			&corev1.ContainerPort{
				Name:          k8s.String(fmt.Sprintf("port-%d", port)),
				HostPort:      k8s.Int32(port),
				ContainerPort: k8s.Int32(port),
			},
		}
	}
	return ctr
}

func (c *Container) toCore() *corev1.Container {
	return c.core
}

// GetServicePorts returns a list of the core ServicePorts that correspond
// to the container ports on c
func (c *Container) GetServicePorts() []*corev1.ServicePort {
	if len(c.core.Ports) > 0 {
		prt := c.core.Ports[0]
		return []*corev1.ServicePort{
			&corev1.ServicePort{
				Port:       k8s.Int32(80),
				TargetPort: newIntOrString(*prt.ContainerPort),
			},
		}
	}
	return nil
}
