package crathens

import (
	"github.com/arschles/athens-azure/pkg/cmd"
	"github.com/spf13/cobra"
)

const name = "crathens-job"
const namespace = "crathens"

func Root(ctx cmd.Context) *cobra.Command {
	ret := cmd.Skeleton("crathens", "Install, update, and administrate Crathens")
	ret.AddCommand(installCmd(ctx))
	ret.AddCommand(currentCmd(ctx))
	return ret
}
