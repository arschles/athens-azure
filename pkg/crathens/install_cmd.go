package crathens

import (
	"context"
	"fmt"

	"github.com/arschles/athens-azure/pkg/kube"

	"github.com/arschles/athens-azure/pkg/cmd"
	"github.com/arschles/athens-azure/pkg/env"
	"github.com/spf13/cobra"
)

const imgEnvVar = "CRATHENS_IMAGE"

func installCmd(ctx context.Context) *cobra.Command {
	ret := cmd.Skeleton("install", "Install Crathens")
	ret.RunE = func(cmd *cobra.Command, args []string) error {
		img, err := env.CheckOr(imgEnvVar, func(key string) (string, error) {
			if len(args) > 0 {
				return args[0], nil
			}
			return "", fmt.Errorf(
				"no image tag given on the command line and no %s env var set",
				imgEnvVar,
			)
		})
		if err != nil {
			return err
		}
		kubeCfg := kube.DiskKubeConfigPath()
		cl, err := kube.LoadClient(kubeCfg)
		if err != nil {
			return err
		}
		if err := kube.UpsertNamespace(ctx, cl, namespace); err != nil {
			return err
		}
		jobContainer := kube.NewContainer("crathens", img)
		job := kube.NewJob(
			"crathens-job",
			namespace,
			kube.ContainerList{jobContainer},
		)

		if err := job.Install(ctx, cl); err != nil {
			return err
		}
		fmt.Println("crathens", img, "created")
		return nil
	}
	return ret
}
