package helm

import (
	"fmt"

	"github.com/arschles/athens-azure/pkg/run"
)

type Set struct {
	Name string
	Val  string
}

// Install installs a helm chart
func Install(chartName, name, ns string, sets []Set) error {
	args := []string{"install", chartName, "--name", name, "--namespace", ns}
	for _, set := range sets {
		args = append(args, "--set", fmt.Sprintf(`"%s=%s"`, set.Name, set.Val))
	}
	_, err := run.Command("helm", args...)
	return err
}
