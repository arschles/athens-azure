package docker

import (
	"github.com/arschles/athens-azure/pkg/cmd"
	"github.com/spf13/cobra"
)

func Root(ctx cmd.Context) *cobra.Command {
	ret := cmd.Skeleton("docker", "Utilities for docker images")
	ret.AddCommand(tagCmd())
	return ret
}
