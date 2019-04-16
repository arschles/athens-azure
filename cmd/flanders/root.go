package main

import (
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

	root.AddCommand(chartmuseum.Root())
	root.AddCommand(censusai.Root())
	root.AddCommand(docker.Root())
	root.AddCommand(ingress.Root())
	return root
}
