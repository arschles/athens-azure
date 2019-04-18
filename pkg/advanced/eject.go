package advanced

import (
	"github.com/arschles/athens-azure/pkg/cmd"
	"github.com/spf13/cobra"
)

func ejectCmd(ctx cmd.Context) *cobra.Command {
	ret := cmd.Skeleton("eject", "Eject! Get a full Helm chart for your app, so you can customize everything (WIP)")
	ret.RunE = func(cmd *cobra.Command, args []string) error {
		ctx.Infof("In progress!")
		return nil
	}
	return ret
}
