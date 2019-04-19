package crathens

import (
	"github.com/arschles/athens-azure/pkg/cmd"
	"github.com/arschles/athens-azure/pkg/kube"
	"github.com/arschles/athens-azure/pkg/kube/resources"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func currentCmd(ctx cmd.Context) *cobra.Command {
	ret := cmd.Skeleton("current", "Get the current state of the running Crathens")

	ret.RunE = func(cmd *cobra.Command, args []string) error {
		cl, err := kube.LoadClientFromDiskKubeConfig()
		if err != nil {
			return errors.WithStack(err)
		}
		job := crathensJob(resources.ContainerList{})
		if err := job.Get(ctx, cl, name, namespace); err != nil {
			return errors.WithStack(err)
		}
		img, err := job.GetImage(0)
		if err != nil {
			return errors.WithStack(err)
		}
		ctx.Infof(img)
		return nil
	}
	return ret
}
