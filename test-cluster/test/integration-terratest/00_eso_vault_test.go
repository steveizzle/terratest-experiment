package test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/k8s"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/require"

	_ "github.com/gruntwork-io/terratest/modules/k8s"
)

func TestEsoVault(t *testing.T) {

	// Construct the terraform options with default retryable errors to handle the most common
	// retryable errors in terraform testing.
	terraformDir := os.Getenv("TF_DIR")
	if terraformDir == "" {
		terraformDir = "../.."
	}

	fmt.Printf("Applying %s\n", terraformDir)

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// Set the path to the Terraform code that will be tested.
		TerraformDir: terraformDir,
	})

	terraformCleanOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// Set the path to the Terraform code that will be tested.
		TerraformDir: terraformDir,
	})

	defer terraform.InitAndApply(t, terraformCleanOptions)

	secretName := "test-sync"

	// Run "terraform init" and "terraform apply". Fail the test if there are any errors.
	terraform.InitAndApply(t, terraformOptions)

	kubectlOptions := k8s.NewKubectlOptions("", "", "default")

	// Create Secret Store, Put Key and Value in Vault and create External Secret
	k8s.KubectlApply(t, kubectlOptions, "00_eso_vault_create.yaml")

	// Clean up Vault and all created Resources after test

	defer k8s.KubectlApply(t, kubectlOptions, "00_eso_vault_clean.yaml")

	defer k8s.KubectlDelete(t, kubectlOptions, "00_eso_vault_create.yaml")

	k8s.WaitUntilSecretAvailable(t, kubectlOptions, secretName, 10, time.Second*10)

	// Decode and Compare Secret. If successful a secret was created from the key in vault

	b64EncodedSecret := k8s.GetSecret(t, kubectlOptions, secretName)

	require.Equal(t, "world", string(b64EncodedSecret.Data["foobar"]))

}
