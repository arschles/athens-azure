package athens

import (
	"github.com/arschles/athens-azure/pkg/cmd"
	"github.com/arschles/athens-azure/pkg/kube"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func currentCmd(ctx cmd.Context) *cobra.Command {
	ret := cmd.Skeleton("current", "Get the current state of the running Athens")

	ret.RunE = func(cmd *cobra.Command, args []string) error {
		cl, err := kube.LoadClientFromDiskKubeConfig()
		if err != nil {
			return errors.WithStack(err)
		}
		depl := athensDeployment("null")
		if err := depl.Get(ctx, cl, name, namespace); err != nil {
			return errors.WithStack(err)
		}
		img, err := depl.GetImage(0)
		if err != nil {
			return errors.WithStack(err)
		}
		ctx.Infof(img)
		return nil
	}
	return ret
}
