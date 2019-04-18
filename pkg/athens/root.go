package athens

import (
	"github.com/arschles/athens-azure/pkg/cmd"
	"github.com/spf13/cobra"
)

const namespace = "athens"
const name = "athens"

func Root(ctx cmd.Context) *cobra.Command {
	ret := cmd.Skeleton("athens", "Install, update, and administrate Athens")
	ret.AddCommand(installCmd(ctx))
	ret.AddCommand(upgradeCmd(ctx))
	ret.AddCommand(currentCmd(ctx))
	return ret
}
