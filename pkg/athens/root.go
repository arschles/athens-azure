package athens

import (
	"context"

	"github.com/arschles/athens-azure/pkg/cmd"
	"github.com/spf13/cobra"
)

func Root() *cobra.Command {
	ctx := context.Background()
	ret := cmd.Skeleton("athens", "Install, update, and administrate Athens")
	ret.AddCommand(installCmd(ctx))
	return ret
}
