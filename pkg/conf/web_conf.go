package conf

type Web struct {
	Domain          string
	Replicas        int32
	Name            string
	Image           string
	Env             map[string]string
	HealthyHTTPPath string
	ReadyHTTPPath   string
	Port            int32
}
