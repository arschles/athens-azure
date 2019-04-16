package ingress

import (
	"github.com/arschles/athens-azure/pkg/cmd"
	"github.com/spf13/cobra"
)

func Root() *cobra.Command {
	return cmd.Skeleton("ingress", "Install, update and modify ingress controllers")
}
