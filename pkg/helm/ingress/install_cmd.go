package ingress

import (
	"github.com/arschles/athens-azure/pkg/cmd"
	"github.com/arschles/athens-azure/pkg/helm"
	"github.com/spf13/cobra"
)

func installCmd() *cobra.Command {
	ret := cmd.Skeleton("install", "Install an ingress controller")
	ret.RunE = func(cmd *cobra.Command, args []string) error {
		return helm.Install("stable/traefik", traefikReleaseName, traefikNS, nil)
	}
	return ret
}
