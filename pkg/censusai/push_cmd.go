package censusai

import (
	"github.com/arschles/athens-azure/pkg/cmd"
	"github.com/arschles/athens-azure/pkg/docker"
	"github.com/arschles/athens-azure/pkg/env"
	"github.com/magefile/mage/sh"
	"github.com/spf13/cobra"
)

func pushCmd() *cobra.Command {
	cmd := cmd.Skeleton("push", "Push the Docker image")
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		image, err := env.Check("CENSUSAI_IMAGE")
		if err != nil {
			return err
		}
		img, err := docker.Tag(image)
		if err != nil {
			return err
		}
		return sh.RunV("docker", "push", img)
	}
	return cmd
}
