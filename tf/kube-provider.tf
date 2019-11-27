variable "kube-host" {
  type = string
  description = "The kubernetes API server hostname"
}
variable "kube-cluster-ca" {
  type = string
  description = "The cluster certificate authority"
}

variable "kube-client-cert" {
  type = string
  description = "The certificate to use for Kubernetes"
}

variable "kube-client-key" {
  type = string
  description = "The private key to use for Kubernetes"
}

provider "kubernetes" {
  host = var.kube-host
  client_certificate = var.kube-client-cert
  client_key = var.kube-client-key
  cluster_ca_certificate = var.kube-cluster-ca
  version = "~> 1.10"
}
