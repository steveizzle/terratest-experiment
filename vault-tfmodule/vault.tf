resource "helm_release" "vault" {
  name       = "vault"

  repository = "https://helm.releases.hashicorp.com"
  chart      = "vault"

  values = [
    "${file("${path.module}/values.yaml")}"
  ]
}
