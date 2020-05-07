resource "kubernetes_deployment" "athens" {
  metadata {
    name      = "athens-server"
    namespace = var.namespace
    labels = {
      kind = "httpserver"
      app  = "athens"
    }
  }
  spec {
    replicas = var.server_replicas
    selector {
      match_labels = {
        kind = "httpserver"
        app  = "athens"
      }
    }
    template {
      metadata {
        labels = {
          kind = "httpserver"
          app  = "athens"
        }
      }
      spec {
        container {
          image             = "gomods/athens:v0.7.0"
          name              = "athens-server"
          image_pull_policy = "IfNotPresent"
          port {
            container_port = 3000
          }
          env {
            name  = "ATHENS_PORT"
            value = ":3000"
          }
          env {
            name  = "ATHENS_STORAGE_TYPE"
            value = "mongo"
          }
          env {
            name  = "ATHENS_MONGO_STORAGE_URL"
            value = var.mongo_conn_string
          }
          env {
            name  = "ATHENS_GO_GET_WORKERS"
            value = var.athens_go_get_workers
          }
          env {
            name  = "GO_ENV"
            value = "development"
          }
          env {
            name  = "ATHENS_TRACE_EXPORTER_URL"
            value = "http://localhost:55678"
          }
          env {
            name  = "ATHENS_TRACE_EXPORTER"
            value = "jaeger"
          }
        }
        container {
          image             = var.census_forwarder_image
          name              = "census-forwarder"
          image_pull_policy = "IfNotPresent"
          port {
            container_port = 50001
          }
          port {
            container_port = 55678
          }
          env {
            name  = "APPINSIGHTS_INSTRUMENTATIONKEY"
            value = var.census_forwarder_app_key
          }
        }
      }
    }
  }
}
