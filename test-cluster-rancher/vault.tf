module "vault_module" {
  source = "git::ssh://git@github.com/steveizzle/vault-tfmodule.git?ref=v0.0.5"
  name = var.vault_name
}
