package athens

import (
	"github.com/arschles/athens-azure/pkg/cmd"
	"github.com/spf13/cobra"
)

func upgradeCmd(ctx cmd.Context) *cobra.Command {
	ret := cmd.Skeleton("upgrade", "Upgrade the Athens image (WIP)")
	ret.RunE = func(cmd *cobra.Command, args []string) error {
		ctx.Infof("In progress!")
		return nil
	}
	return ret
}
