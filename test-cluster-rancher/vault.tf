module "vault_module" {
  source = "git::ssh://git@github.com/steveizzle/terratest-experiment.git//vault-tfmodule?ref=main"
  name = var.vault_name
}
