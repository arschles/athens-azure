package athens

import (
	"github.com/arschles/athens-azure/pkg/cmd"
	"github.com/arschles/athens-azure/pkg/conf"
	"github.com/arschles/athens-azure/pkg/kube"
	"github.com/spf13/cobra"
)

func currentCmd(ctx cmd.Context) *cobra.Command {
	ret := cmd.Skeleton("current", "Get the current state of the running Athens")
	ret.RunE = func(cmd *cobra.Command, args []string) error {
		cfg, err := conf.Athens()
		if err != nil {
			return err
		}
		ctx.Debugf("Looking for web images %s", cfg.WebImages())
		ctx.Debugf("Looking for job images: %s", cfg.JobImages())
		cl, err := kube.LoadClientFromDiskKubeConfig()
		if err != nil {
			return err
		}
		prof := newProfile(cfg.Webs, cfg.Jobs)
		if err := prof.Status(ctx, cl); err != nil {
			ctx.Infof("Error! %s", err)
			return err
		}
		ctx.Infof("All is good :)")
		return nil
	}
	return ret
}
