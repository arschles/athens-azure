package advanced

import (
	"github.com/arschles/athens-azure/pkg/cmd"
	"github.com/spf13/cobra"
)

func resourcesCmd(ctx cmd.Context) *cobra.Command {
	ret := cmd.Skeleton("resources", "Get detailed info on all the Kubernetes resources your app is using (WIP)")
	ret.RunE = func(cmd *cobra.Command, args []string) error {
		ctx.Infof("In progress!")
		return nil
	}
	return ret
}
