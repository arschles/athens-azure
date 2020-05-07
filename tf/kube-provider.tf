variable "kube_host" {
  type        = string
  description = "The kubernetes API server hostname"
}
variable "kube_cluster_ca" {
  type        = string
  description = "The cluster certificate authority"
}

variable "kube_client_cert" {
  type        = string
  description = "The certificate to use for Kubernetes"
}

variable "kube_client_key" {
  type        = string
  description = "The private key to use for Kubernetes"
}

provider "kubernetes" {
  host                   = var.kube_host
  client_certificate     = var.kube_client_cert
  client_key             = var.kube_client_key
  cluster_ca_certificate = var.kube_cluster_ca
  version                = "~> 1.10"
}
