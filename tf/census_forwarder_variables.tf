variable "census_forwarder_image" {
  type        = string
  description = "The Docker image name of the OpenCensus to App Insights forwarder"
}

variable "census_forwarder_app_key" {
  type        = string
  description = "The application key for the OpenCensus to App Insights forwarder"
}
