package crathens

import (
	"fmt"

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
		if err := kube.UpsertNamespace(ctx, cl, namespace); err != nil {
			return err
		}

		jobContainer := kube.NewContainer("crathens", img)
		job := crathensJob(kube.ContainerList{jobContainer})

		ctx.Debugf("job:\n%s", job)

		if err := job.Install(ctx, cl); err != nil {
			return err
		}
		fmt.Println("crathens", img, "created")
		return nil
	}
	return ret
}
