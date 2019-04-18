package athens

import (
	"github.com/arschles/athens-azure/pkg/cmd"
	"github.com/arschles/athens-azure/pkg/kube"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func upgradeCmd(ctx cmd.Context) *cobra.Command {
	ret := cmd.Skeleton("upgrade", "Upgrade the Athens image (WIP)")
	ret.RunE = func(cmd *cobra.Command, args []string) error {
		newImg, err := getImage(args)
		if err != nil {
			return err
		}
		cl, err := kube.LoadClientFromDiskKubeConfig()
		if err != nil {
			return errors.WithStack(err)
		}
		profile := newProfile(newImg)
		if err := profile.Update(ctx, cl, kube.ErrorStrategyContinue); err != nil {
			return err
		}
		ctx.Infof("Updated!")
		return nil
	}
	return ret
}
