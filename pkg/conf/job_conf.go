package conf

import "fmt"

type Job struct {
	fmt.Stringer
	Name  string
	Image string
}

func (j Job) String() string {
	return fmt.Sprintf("Job %s: %s", j.Name, j.Image)
}
