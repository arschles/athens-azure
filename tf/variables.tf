variable "namespace" {
  type        = string
  description = "The namespace to put the Athens server & friends in"
  default     = "athens-server"
}

variable "server-replicas" {
  type        = string
  description = "The number of replicas of the Athens proxy server"
  default     = "5"
}

variable "mongo-conn-string" {
  type        = string
  description = "The connection string for Athens to connect to Mongo (CosmosDB in prod)"
}

variable "athens-go-get-workers" {
  type        = string
  description = "The number of Go processes that Athens will run in the background"
  default     = "3"
}

variable "lathens-image-tag" {
  type        = string
  description = "The tag of the lathens docker image. This is usually the Git SHA, which you can get with 'make git-sha'"
}
