module "vault_module" {
  source = "git::ssh://git@github.com/steveizzle/vault-tfmodule.git?ref=v0.0.7-rc5"
  name = var.vault_name
}
