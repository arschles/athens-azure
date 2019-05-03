package conf

import "fmt"

type Web struct {
	fmt.Stringer
	Domain          string
	Replicas        int32
	Name            string
	Image           string
	Env             map[string]string
	HealthyHTTPPath string
	ReadyHTTPPath   string
	Port            int32
}

func (w Web) String() string {
	return fmt.Sprintf(
		"Web %s: %s with %d replicas",
		w.Name,
		w.Image,
		w.Replicas,
	)
}
