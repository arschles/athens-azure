package conf

import "fmt"

// web {
//     domain = "my.domain"
//     // optional, defaults to 3
//     Replicas = 123
//     Name = "Athens"
//     Image = "quay.io/gomods/athens:v0.3.1"
//     // optional
//     Env = {
//         "SOMETHING": "SOMETHING_ELSE"
//     }
//     // optional
//     HealthyHTTPPath = "/healthy"
//     // optional
//     ReadyHTTPPath = "/ready"
//     // optional, defaults to no ports exposed
//     Port = 8080
// }

// job {
//     Name = "Crathens"
//     Image = "quay.io/arschles/crathens:canary"
// }

// Root is the root of all configuration
type Root struct {
	fmt.Stringer
	Name string
	Webs []Web `json:"web"`
	Jobs []Job `json:"job"`
}

// JobImages gets a list of all the images in jobs
func (r *Root) JobImages() []string {
	ret := make([]string, len(r.Jobs))
	for i, job := range r.Jobs {
		ret[i] = job.Image
	}
	return ret
}

// WebImages gets a list of all the images in webs
func (r *Root) WebImages() []string {
	ret := make([]string, len(r.Webs))
	for i, web := range r.Webs {
		ret[i] = web.Image
	}
	return ret
}

func (r *Root) String() string {
	return fmt.Sprintf(
		"App %s with %d webs and %d jobs",
		r.Name,
		len(r.Webs),
		len(r.Jobs),
	)
}
