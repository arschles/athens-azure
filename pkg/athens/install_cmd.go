package athens

import (
	"context"
	"fmt"

	"github.com/arschles/athens-azure/pkg/cmd"
	"github.com/arschles/athens-azure/pkg/env"
	"github.com/arschles/athens-azure/pkg/kube"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

const namespace = "athens"

func installCmd(ctx context.Context) *cobra.Command {
	ret := cmd.Skeleton("install", "Install Athens")
	ret.RunE = func(cmd *cobra.Command, args []string) error {

		img, err := env.CheckOrArg("ATHENS_IMAGE", args, 0)
		if err != nil {
			return errors.WithStack(err)
		}
		cl, err := kube.LoadClientFromDiskKubeConfig()
		if err != nil {
			return errors.WithStack(err)
		}
		if err := kube.UpsertNamespace(ctx, cl, namespace); err != nil {
			return errors.WithStack(err)
		}
		fmt.Println("Created namespace", namespace)
		fmt.Println("Creating service")
		// service := athensService(port)
		// if err := service.Install(ctx, cl); err != nil {
		// 	return err
		// }

		fmt.Println("Deploying", img)
		// deployment := athensDeployment(img)
		// if err := deployment.Install(ctx, cl); err != nil {
		// 	return err
		// }

		// return helm.Install("./charts/athens", "athens", "athens", []helm.Set{
		// 	{Name: "goGetWorkers", Val: "2"},
		// })
		fmt.Println("Deployment created")

		fmt.Println("Creating ingress")
		// ingress :=  athensIngress(svc)
		// if err := ingress.Install(ctx, cl); err != nil {
		// 	return err
		// }
		fmt.Println("Ingress created")
		return nil
	}
	return ret
}

func athensDeployment(img string) *kube.Deployment {
	containerList := kube.ContainerList{
		kube.NewContainer("athens", img),
	}
	return kube.NewDeployment("athens", namespace, containerList)
}
