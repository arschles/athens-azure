package docker

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/arschles/athens-azure/pkg/cmd"
	"github.com/spf13/cobra"
)

func Tag(image string) (string, error) {
	cmd := exec.Command("git", "rev-parse", "--short", "HEAD")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	tag := strings.TrimRight(string(out), "\n")
	return image + ":" + tag, nil
}

func tagCmd() *cobra.Command {
	ret := cmd.Skeleton("tag", "Output the fully tagged name of a given Docker image")
	ret.RunE = func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("No Docker image provided")
		}
		img := args[0]
		fmt.Println(Tag(img))
		return nil
	}
	return ret
}
