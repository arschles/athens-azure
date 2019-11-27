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

variable "kubeconfig-path" {
  type        = string
  description = "The path to the Kube config file"
}

variable "server-replicas" {
  type = string
  description = "The number of replicas of the Athens proxy server"
  default = "5"
}

variable "mongo-conn-string" {
  type = string
  description = "The connection string for Athens to connect to Mongo (CosmosDB in prod)"
}

variable "athens-go-get-workers" {
  type = string
  description = "The number of Go processes that Athens will run in the background"
  default = "3"
}
