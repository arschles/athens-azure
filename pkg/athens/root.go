package athens

import (
	"github.com/arschles/athens-azure/pkg/cmd"
	"github.com/spf13/cobra"
)

func Root(ctx cmd.Context) *cobra.Command {
	ret := cmd.Skeleton("athens", "Install, update, and administrate Athens")
	ret.AddCommand(installCmd(ctx))
	return ret
}
