package crathens

import (
	"github.com/arschles/athens-azure/pkg/cmd"
	"github.com/spf13/cobra"
)

const namespace = "crathens"

func Root(ctx cmd.Context) *cobra.Command {
	debug := false
	ret := cmd.Skeleton("crathens", "Install, update, and administrate Crathens")
	ret.AddCommand(installCmd(ctx, debug))
	return ret
}
