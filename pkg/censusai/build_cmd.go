package censusai

import (
	"fmt"

	"github.com/arschles/athens-azure/pkg/cmd"
	"github.com/arschles/athens-azure/pkg/docker"
	"github.com/arschles/athens-azure/pkg/env"
	"github.com/arschles/athens-azure/pkg/run"
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
		res, err := run.Command("docker", "build", "-t", img, "./census")
		if err != nil {
			return err
		}
		fmt.Println(res)
		return nil
	}
	return cmd
}
