package ingress

import (
	"github.com/arschles/athens-azure/pkg/cmd"
	"github.com/spf13/cobra"
)

const traefikReleaseName = "traefik"
const traefikNS = "traefik"

func Root(ctx cmd.Context) *cobra.Command {
	ret := cmd.Skeleton("ingress", "Install, update and modify ingress controllers")
	ret.AddCommand(installCmd())
	ret.AddCommand(pubIPCmd(ctx))
	return ret

}
