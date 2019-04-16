package kube

import (
	"fmt"
	"io/ioutil"

	"github.com/ericchiang/k8s"
	"github.com/ghodss/yaml"
)

// LoadClient parses a kubeconfig from a file and returns a Kubernetes
// client. It does not support extensions or client auth providers.
//
// This code was taken from:
// // https://github.com/ericchiang/k8s/tree/68fb2168bedf77759577a56e44f2ccfaf7229673#creating-out-of-cluster-clients
func LoadClient(configPath string) (*k8s.Client, error) {
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("Could not read KubeConfig (%s)", err)
	}

	// Unmarshal YAML into a Kubernetes config object.
	var config k8s.Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("unmarshal kubeconfig: %v", err)
	}
	return k8s.NewClient(&config)
}
