variable "namespace" {
  type        = string
  description = "The namespace to put the Athens server & friends in"
  default     = "athens-server"
}

variable "server_replicas" {
  type        = string
  description = "The number of replicas of the Athens proxy server"
  default     = "5"
}

variable "mongo_conn_string" {
  type        = string
  description = "The connection string for Athens to connect to Mongo (CosmosDB in prod)"
}

variable "athens_go_get_workers" {
  type        = string
  description = "The number of Go processes that Athens will run in the background"
  default     = "3"
}
