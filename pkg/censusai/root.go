package censusai

import (
	"github.com/arschles/athens-azure/pkg/cmd"
	"github.com/spf13/cobra"
)

func Root() *cobra.Command {
	cmd := cmd.Skeleton(
		"censusai",
		"Tools to build the OpenCensus Sidecar for Application Insights",
	)
	cmd.AddCommand(buildCmd())
	cmd.AddCommand(pushCmd())
	return cmd
}
