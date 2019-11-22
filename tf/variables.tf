variable "kube-auth" {
  type        = string
  description = "The Kubernetes user or auth info (from the kubeconfig file) to use"
  default     = "clusterUser_athens_flanders"
}

variable "kube-cluster" {
  type        = string
  description = "The Kubernetes cluster (from the kubeconfig file) to use"
  default     = "flanders"
}

variable "namespace" {
  type        = string
  description = "The namespace to put the Athens server & friends in"
  default     = "athens-server"
}
