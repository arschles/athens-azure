package censusai

import (
	"fmt"

	"github.com/arschles/athens-azure/pkg/cmd"
	"github.com/arschles/athens-azure/pkg/docker"
	"github.com/arschles/athens-azure/pkg/env"
	"github.com/magefile/mage/sh"
	"github.com/spf13/cobra"
)

func buildCmd() *cobra.Command {
	cmd := cmd.Skeleton("build", "Build the Docker image")
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		image, err := env.Check("CENSUSAI_IMAGE")
		if err != nil {
			return err
		}
		img, err := docker.Tag(image)
		if err != nil {
			return err
		}
		fmt.Println("Building", img)
		return sh.RunV("docker", "build", "-t", img, "./census")
	}
	return cmd
}
