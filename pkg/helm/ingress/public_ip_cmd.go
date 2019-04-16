package ingress

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/arschles/athens-azure/pkg/cmd"
	"github.com/arschles/athens-azure/pkg/env"
	"github.com/arschles/athens-azure/pkg/kube"
	v1 "github.com/ericchiang/k8s/apis/core/v1"
	"github.com/spf13/cobra"
)

func pubIPCmd(ctx context.Context) *cobra.Command {
	ret := cmd.Skeleton("pubip", "Get the public IP of the ingress controller")
	ret.RunE = func(cmd *cobra.Command, args []string) error {
		kubeCfg, err := env.Check("KUBECONFIG")
		if err != nil {
			home := os.Getenv("HOME")
			kubeCfg = filepath.Join(home, ".kube", "config")
		}
		cl, err := kube.LoadClient(kubeCfg)
		if err != nil {
			return err
		}
		var svc v1.Service
		if err := cl.Get(ctx, traefikNS, traefikReleaseName, &svc); err != nil {
			return err
		}
		fmt.Print(svc)

		return nil
	}
	return ret
}