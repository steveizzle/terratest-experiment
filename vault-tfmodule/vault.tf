resource "helm_release" "vault" {
  name       = var.name

  repository = "https://helm.releases.hashicorp.com"
  chart      = "vault"

  values = [
    "${file("${path.module}/values.yaml")}"
  ]
}
