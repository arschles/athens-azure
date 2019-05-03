package athens

import (
	"github.com/arschles/athens-azure/pkg/cmd"
	"github.com/arschles/athens-azure/pkg/kube"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func updateCmd(ctx cmd.Context) *cobra.Command {
	ret := cmd.Skeleton("update", "Upgrade the Athens image (WIP)")
	flags := ret.PersistentFlags()
	ret.RunE = func(cmd *cobra.Command, args []string) error {
		imgs, err := getImages(flags, ret.MarkFlagRequired)
		if err != nil {
			return err
		}
		cl, err := kube.LoadClientFromDiskKubeConfig()
		if err != nil {
			return errors.WithStack(err)
		}
		profile := newProfile(imgs)
		if err := profile.Update(ctx, cl, kube.ErrorStrategyContinue); err != nil {
			return err
		}
		ctx.Infof("Updated!")
		return nil
	}
	return ret
}
