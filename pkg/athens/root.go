package athens

import (
	"github.com/arschles/athens-azure/pkg/cmd"
	"github.com/spf13/cobra"
)

const (
	athensNS     = "athens"
	athensName   = "athens"
	crathensNS   = "crathens"
	crathensName = "crathens"
	lathensNS    = "lathens"
	lathensName  = "lathens"
)

// Root returns the tree under 'flanders athens'
func Root(ctx cmd.Context) *cobra.Command {
	ret := cmd.Skeleton("athens", "Install, update, and administrate Athens")
	ret.AddCommand(installCmd(ctx))
	ret.AddCommand(uninstallCmd(ctx))
	ret.AddCommand(updateCmd(ctx))
	ret.AddCommand(currentCmd(ctx))
	return ret
}
