package athens

import (
	"github.com/arschles/athens-azure/pkg/cmd"
	"github.com/arschles/athens-azure/pkg/conf"
	"github.com/arschles/athens-azure/pkg/kube"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func installCmd(ctx cmd.Context) *cobra.Command {
	ret := cmd.Skeleton("install", "Install Athens")
	flags := ret.PersistentFlags()
	var dryRun bool
	flags.BoolVar(&dryRun, "dryrun", false, "Do a dry run, and don't modify any Kubernetes resources")

	ret.RunE = func(cmd *cobra.Command, args []string) error {
		cfg, err := conf.Athens()
		if err != nil {
			return errors.WithStack(err)
		}
		cl, err := kube.LoadClientFromDiskKubeConfig()
		if err != nil {
			return errors.WithStack(err)
		}

		profile := newProfile(cfg.Webs, cfg.Jobs)
		if ctx.IsDebug() {
			ctx.Debugf("Here are the resources that are going to be installed:")
			resources := profile.AllResources()
			for _, res := range resources {
				ctx.Debugf("%s %s/%s", res.Type(), res.Namespace().Name(), res.Name())
				ctx.Debugf("%s", res)
			}
		}
		ctx.Infof("Setting up and installing:\n%s", profile)
		if dryRun {
			ctx.Infof("----> Not doing anything because this is a dry run")
		} else {
			if err := kube.SetupAndInstallProfile(
				ctx,
				cl,
				profile,
				kube.ErrorStrategyContinue,
			); err != nil {
				return err
			}
		}
		ctx.Infof("Done")
		return nil
	}
	return ret
}
