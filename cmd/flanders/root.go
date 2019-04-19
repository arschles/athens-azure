package main

import (
	"context"

	"github.com/arschles/athens-azure/pkg/advanced"
	"github.com/arschles/athens-azure/pkg/athens"
	"github.com/arschles/athens-azure/pkg/censusai"
	"github.com/arschles/athens-azure/pkg/cmd"
	"github.com/arschles/athens-azure/pkg/docker"
	"github.com/arschles/athens-azure/pkg/helm/chartmuseum"
	"github.com/arschles/athens-azure/pkg/helm/ingress"
	"github.com/spf13/cobra"
)

// Root sets up the command for the top level a3e probram,
// adds all child commands, and does additional configuration.
// It only needs to be called once and then run by the main function
func Root() *cobra.Command {
	root := cmd.Skeleton("flanders", "Do container-ey and Kubernetes-ey things with less fuss")
	var debug bool
	flags := root.PersistentFlags()
	flags.BoolVarP(&debug, "debug", "d", false, "Turn on debug logging")
	c := context.Background()
	ctx := cmd.NewContext(c, debug)

	root.AddCommand(chartmuseum.Root(ctx))
	root.AddCommand(censusai.Root(ctx))
	root.AddCommand(docker.Root(ctx))
	root.AddCommand(ingress.Root(ctx))
	root.AddCommand(athens.Root(ctx))
	root.AddCommand(advanced.Root(ctx))
	return root
}
