package cmd

import (
	"github.com/spf13/cobra"
)

func Skeleton(use, short string) *cobra.Command {
	// TODO: debug
	return &cobra.Command{
		Use:          use,
		Short:        short,
		SilenceUsage: true,
		//SilenceErrors: true,
	}
}
