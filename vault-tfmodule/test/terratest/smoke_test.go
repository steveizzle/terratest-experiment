package test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/gruntwork-io/terratest/modules/k8s"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestTerraformHelloWorldExample(t *testing.T) {

	vaultName := "vault-test"
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
		Vars: map[string]interface{}{
			"vault_name": vaultName,
		},
	})

	// Clean up resources with "terraform destroy" at the end of the test.
	// defer terraform.Destroy(t, terraformOptions)

	// Run "terraform init" and "terraform apply". Fail the test if there are any errors.
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the values of output variables and check they have the expected values.
	// output := terraform.Output(t, terraformOptions, "hello_world")
	k8s.WaitUntilPodAvailable(t, &k8s.KubectlOptions{Namespace: "Default"}, fmt.Sprintf("%s-0", vaultName), 4, time.Duration(time.Second*10))

	fmt.Println("Clean up")

	terraformOptions = terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// Set the path to the Terraform code that will be tested.
		TerraformDir: terraformDir,
	})

	terraform.InitAndApply(t, terraformOptions)
}
