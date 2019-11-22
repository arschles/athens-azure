resource "kubernetes_service" "example" {
  metadata {
    name = "athens"
    namespace = var.namespace
  }
  spec {
    selector = {
      kind = "httpserver"
      app = "athens"
      #  TODO: make this come from the deployment. For example:
      # ${kubernetes_deployment.athens-server.spec.template.metadata.labels.app}
    }
    port {
      port        = 8080
      target_port = 3000
    }
    type = "LoadBalancer"
  }
}
