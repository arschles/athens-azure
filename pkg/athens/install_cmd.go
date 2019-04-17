package athens

import (
	"github.com/arschles/athens-azure/pkg/cmd"
	"github.com/arschles/athens-azure/pkg/helm"
	"github.com/spf13/cobra"
)

func installCmd() *cobra.Command {
	ret := cmd.Skeleton("install", "Install Athens")
	ret.RunE = func(cmd *cobra.Command, args []string) error {
		return helm.Install("./charts/athens", "athens", "athens", []helm.Set{
			{Name: "goGetWorkers", Val: "2"},
		})
	}
	return ret
}
