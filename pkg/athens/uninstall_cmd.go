package athens

import (
	"github.com/arschles/athens-azure/pkg/cmd"
	"github.com/arschles/athens-azure/pkg/conf"
	"github.com/arschles/athens-azure/pkg/kube"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func uninstallCmd(ctx cmd.Context) *cobra.Command {
	ret := cmd.Skeleton("uninstall", "Uninstall Athens")
	ret.RunE = func(cmd *cobra.Command, args []string) error {
		cfg, err := conf.Athens()
		if err != nil {
			return err
		}
		cl, err := kube.LoadClientFromDiskKubeConfig()
		if err != nil {
			return errors.WithStack(err)
		}

		athensProfile := newProfile(cfg.Webs, cfg.Jobs)
		if ctx.IsDebug() {
			ctx.Debugf("Here are the resources that are going to be uninstalled:")
			resources := athensProfile.AllResources()
			for _, res := range resources {
				ctx.Debugf("%s %s/%s", res.Type(), res.Namespace().Name(), res.Name())
				ctx.Debugf("%s", res)
			}
		}
		ctx.Infof("Uninstalling %s", athensProfile)
		if err := athensProfile.Uninstall(
			ctx,
			cl,
			kube.ErrorStrategyContinue,
		); err != nil {
			return err
		}
		ctx.Infof("Done")

		return nil
	}
	return ret
}
