package kube

import (
	"os"
	"path/filepath"

	"github.com/arschles/athens-azure/pkg/env"
)

// DiskKubeConfigPath returns the path to a KubeConfig file on disk.
//
// This is meant for client-side access to Kubernetes, not in-cluster access
func DiskKubeConfigPath() string {
	kubeCfg, err := env.Check("KUBECONFIG")
	if err != nil {
		home := os.Getenv("HOME")
		kubeCfg = filepath.Join(home, ".kube", "config")
	}
	return kubeCfg
}
