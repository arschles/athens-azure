resource "kubernetes_namespace" "athens-namespace" {
  metadata {
    name = var.namespace
  }
}
