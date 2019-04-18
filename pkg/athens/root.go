package athens

import (
	"github.com/arschles/athens-azure/pkg/cmd"
	"github.com/arschles/athens-azure/pkg/env"
	"github.com/spf13/cobra"
)

const namespace = "athens"
const name = "athens"

func getImage(args []string) (string, error) {
	return env.CheckOrArg("ATHENS_IMAGE", args, 0)
}

// Root returns the tree under 'flanders athens'
func Root(ctx cmd.Context) *cobra.Command {
	ret := cmd.Skeleton("athens", "Install, update, and administrate Athens")
	ret.AddCommand(installCmd(ctx))
	ret.AddCommand(uninstallCmd(ctx))
	ret.AddCommand(upgradeCmd(ctx))
	ret.AddCommand(currentCmd(ctx))
	return ret
}
