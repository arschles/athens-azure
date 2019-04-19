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
		imgs, err := getImages(args)
		if err != nil {
			return err
		}
		cl, err := kube.LoadClientFromDiskKubeConfig()
		if err != nil {
			return errors.WithStack(err)
		}
		prof := newProfile(imgs)
		if err := prof.Status(ctx, cl); err != nil {
			ctx.Infof("Error! %s", err)
			return err
		}
		ctx.Infof("All is good :)")
		return nil
	}
	return ret
}
