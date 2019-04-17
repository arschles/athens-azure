package athens

import (
	"github.com/arschles/athens-azure/pkg/cmd"
	"github.com/arschles/athens-azure/pkg/env"
	"github.com/arschles/athens-azure/pkg/kube"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func upgradeCmd(ctx cmd.Context) *cobra.Command {
	ret := cmd.Skeleton("upgrade", "Upgrade the Athens image")
	ret.RunE = func(cmd *cobra.Command, args []string) error {
		img, err := env.CheckOrArg("ATHENS_IMAGE", args, 0)
		if err != nil {
			return errors.WithStack(err)
		}
		deployment := athensDeployment(img)
		cl, err := kube.LoadClientFromDiskKubeConfig()
		if err != nil {
			return errors.WithStack(err)
		}
		if err := deployment.Update(ctx, cl); err != nil {
			return errors.WithStack(err)
		}
		return nil
	}
	return ret
}
