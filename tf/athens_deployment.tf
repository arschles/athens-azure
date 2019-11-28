resource "kubernetes_deployment" "athens-server" {
  metadata {
    name      = "athens-server"
    namespace = var.namespace
    labels = {
      kind = "httpserver"
      app  = "athens"
    }
  }
  spec {
    replicas = var.server-replicas
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
            value = var.mongo-conn-string
          }
          env {
            name  = "ATHENS_GO_GET_WORKERS"
            value = var.athens-go-get-workers
          }
          env {
            name  = "GO_ENV"
            value = "development"
          }
          env {
            name  = "ATHENS_TRACE_EXPORTER_URL"
            value = "http://localhost:55678"
          }
        }
        container {
          image             = var.census-forwarder-image
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
            value = var.census-forwarder-app-key
          }
        }
      }
    }
  }
}
