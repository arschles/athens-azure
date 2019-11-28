variable "census-forwarder-image" {
  type        = string
  description = "The Docker image name of the OpenCensus to App Insights forwarder"
}

variable "census-forwarder-app-key" {
  type        = string
  description = "The application key for the OpenCensus to App Insights forwarder"
}
