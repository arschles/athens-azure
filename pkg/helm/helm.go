package helm

import (
	"fmt"

	"github.com/magefile/mage/sh"
)

type Set struct {
	Name string
	Val  string
}

func Install(chartName, name, ns string, sets []Set) error {
	args := []string{"install", chartName, "--name", name, "--namespace", ns}
	for _, set := range sets {
		args = append(args, "--set", fmt.Sprintf(`"%s=%s"`, set.Name, set.Val))
	}
	return sh.Run("helm", args...)
}
