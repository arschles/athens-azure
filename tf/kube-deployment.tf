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
    replicas = 5
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
            name  = "PORT"
            value = ":3000"
          }
        }
      }
    }
  }
}
