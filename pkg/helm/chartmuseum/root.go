package chartmuseum

import (
	"github.com/arschles/athens-azure/pkg/cmd"
	"github.com/arschles/athens-azure/pkg/env"
	"github.com/arschles/athens-azure/pkg/helm"
	"github.com/spf13/cobra"
)

const accountVarName = "AZURE_STORAGE_ACCOUNT"
const accessKeyVarName = "AZURE_STORAGE_ACCESS_KEY"

func Root(ctx cmd.Context) *cobra.Command {
	cmd := cmd.Skeleton("chartmuseum", "do stuff with chartmuseum")
	cmd.AddCommand(installCmd())
	return cmd
}

func installCmd() *cobra.Command {
	cmd := cmd.Skeleton("install", "install Chartmuseum")
	cmd.RunE = func(cmd *cobra.Command, args []string) error {

		account, err := env.Check(accountVarName)
		if err != nil {
			return err
		}
		accessKey, err := env.Check(accessKeyVarName)
		if err != nil {
			return err
		}
		sets := []helm.Set{
			{Name: "env.open.STORAGE", Val: "microsoft"},
			{Name: "env.open.STORAGE_MICROSOFT_CONTAINER", Val: "chartmuseum"},
			{Name: "env.open.STORAGE_MICROSOFT_PREFIX", Val: ""},
			{Name: "env.secret.AZURE_STORAGE_ACCOUNT", Val: account},
			{Name: "env.secret.AZURE_STORAGE_ACCESS_KEY", Val: accessKey},
		}
		return helm.Install("stable/chartmuseum", "chartmuseum", "chartmuseum", sets)
	}
	return cmd
}
