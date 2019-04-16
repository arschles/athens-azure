package censusai

import (
	"fmt"

	"github.com/arschles/athens-azure/pkg/cmd"
	"github.com/arschles/athens-azure/pkg/docker"
	"github.com/arschles/athens-azure/pkg/env"
	"github.com/arschles/athens-azure/pkg/run"
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
		res, err := run.Command("docker", "push", img)
		if err != nil {
			return err
		}
		fmt.Println(res)
		return nil
	}
	return cmd
}
