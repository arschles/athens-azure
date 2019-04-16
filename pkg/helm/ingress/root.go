package ingress

import (
	"context"

	"github.com/arschles/athens-azure/pkg/cmd"
	"github.com/spf13/cobra"
)

const traefikReleaseName = "traefik"
const traefikNS = "traefik"

func Root() *cobra.Command {
	ctx := context.Background()
	ret := cmd.Skeleton("ingress", "Install, update and modify ingress controllers")
	ret.AddCommand(installCmd())
	ret.AddCommand(pubIPCmd(ctx))
	return ret

}
