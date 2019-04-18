package advanced

import (
	"github.com/arschles/athens-azure/pkg/cmd"
	"github.com/spf13/cobra"
)

// Root returns the root of the 'flanders advanced' command
func Root(ctx cmd.Context) *cobra.Command {
	ret := cmd.Skeleton("advanced", "Advanced features for flanders")
	ret.AddCommand(ejectCmd(ctx))
	ret.AddCommand(resourcesCmd(ctx))
	return ret
}
