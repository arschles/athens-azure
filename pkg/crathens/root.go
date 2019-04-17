package crathens

import (
	"context"

	"github.com/arschles/athens-azure/pkg/cmd"
	"github.com/spf13/cobra"
)

const namespace = "crathens"

func Root() *cobra.Command {
	ctx := context.Background()
	ret := cmd.Skeleton("crathens", "Install, update, and administrate Crathens")
	ret.AddCommand(installCmd(ctx))
	return ret
}
