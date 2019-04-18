package athens

import (
	"github.com/arschles/athens-azure/pkg/cmd"
	"github.com/arschles/athens-azure/pkg/env"
	"github.com/arschles/athens-azure/pkg/kube"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func installCmd(ctx cmd.Context) *cobra.Command {
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

		athensProfile := newProfile(img)
		if ctx.IsDebug() {
			ctx.Debugf("Here are the resources that are going to be installed:")
			resources := athensProfile.AllResources()
			for _, res := range resources {
				ctx.Debugf("%s %s/%s", res.Type(), res.Namespace().Name(), res.Name())
				ctx.Debugf("%s", res)
			}
		}
		ctx.Infof("Setting up and installing %s", athensProfile)
		if err := kube.SetupAndInstallProfile(
			ctx,
			cl,
			athensProfile,
			kube.ErrorStrategyContinue,
		); err != nil {
			return err
		}
		ctx.Infof("Done")
		// if err := kube.UpsertNamespace(ctx, cl, namespace); err != nil {
		// 	return errors.WithStack(err)
		// }
		// ctx.Debugf("Created namespace %s", namespace)
		// ctx.Debugf("Creating service")
		// service := athensService(port)
		// if err := service.Install(ctx, cl); err != nil {
		// 	return err
		// }

		// ctx.Debugf("Deploying %s", img)
		// deployment := athensDeployment(img)
		// if err := deployment.Install(ctx, cl); err != nil {
		// 	return err
		// }

		// return helm.Install("./charts/athens", "athens", "athens", []helm.Set{
		// 	{Name: "goGetWorkers", Val: "2"},
		// })
		// ctx.Debugf("Deployment created")

		// ctx.Debugf("Creating ingress")
		// ingress :=  athensIngress(svc)
		// if err := ingress.Install(ctx, cl); err != nil {
		// 	return err
		// }
		// ctx.Debugf("Ingress created")
		return nil
	}
	return ret
}
