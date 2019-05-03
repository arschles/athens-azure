package athens

import (
	"github.com/arschles/athens-azure/pkg/cmd"
	"github.com/arschles/athens-azure/pkg/conf"
	"github.com/arschles/athens-azure/pkg/kube"
	"github.com/spf13/cobra"
)

func updateCmd(ctx cmd.Context) *cobra.Command {
	ret := cmd.Skeleton("update", "Upgrade the Athens image (WIP)")
	ret.RunE = func(cmd *cobra.Command, args []string) error {
		cfg, err := conf.Athens()
		if err != nil {
			return err
		}
		cl, err := kube.LoadClientFromDiskKubeConfig()
		if err != nil {
			return err
		}
		profile := newProfile(cfg.Webs, cfg.Jobs)
		if err := profile.Update(ctx, cl, kube.ErrorStrategyContinue); err != nil {
			return err
		}
		ctx.Infof("Updated!")
		return nil
	}
	return ret
}
