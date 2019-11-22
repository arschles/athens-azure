resource "kubernetes_namespace" "athens-azurefd" {
  metadata {
    name = var.namespace
  }
}
