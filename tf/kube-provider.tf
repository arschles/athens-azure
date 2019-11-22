provider "kubernetes" {
  config_path = var.kubeconfig-path
  config_context_auth_info = var.kube-auth
  config_context_cluster   = var.kube-cluster
}
