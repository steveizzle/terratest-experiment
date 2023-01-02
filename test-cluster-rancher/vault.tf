module "vault_module" {
  source = "git::ssh://git@github.com/steveizzle/terratest-experiment.git//vault-tfmodule?ref=main"
}
module "vault_module_2" {
  source = "git::ssh://git@github.com/steveizzle/terratest-experiment.git//vault-tfmodule?ref=main"
  name = "vault-2"
}
