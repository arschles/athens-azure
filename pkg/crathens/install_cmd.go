package crathens

import (
	"github.com/arschles/athens-azure/pkg/env"
	"github.com/arschles/athens-azure/pkg/kube"

	"github.com/arschles/athens-azure/pkg/cmd"
	"github.com/spf13/cobra"
)

const imgEnvVar = "CRATHENS_IMAGE"

func installCmd(ctx cmd.Context) *cobra.Command {
	ret := cmd.Skeleton("install", "Install Crathens")
	ret.RunE = func(cmd *cobra.Command, args []string) error {
		img, err := env.CheckOrArg(imgEnvVar, args, 0)
		if err != nil {
			return err
		}
		cl, err := kube.LoadClientFromDiskKubeConfig()
		if err != nil {
			return err
		}

		jobContainer := kube.NewContainer("crathens", img)
		job := crathensJob(kube.ContainerList{jobContainer})
		jobProfile := kube.NewLongRunningBatchProfile(job)

		ctx.Infof("Setting up and installing profile %s", jobProfile)
		if err := kube.SetupAndInstallProfile(
			ctx,
			cl,
			jobProfile,
			kube.ErrorStrategyContinue,
		); err != nil {
			return err
		}

		ctx.Infof("Done")
		return nil
	}
	return ret
}
