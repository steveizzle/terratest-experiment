#!/bin/bash

INTEGRATION_TESTS="vault_module eso"

echo "Installing dependencies"
go install github.com/gruntwork-io/terratest/cmd/terratest_log_parser@latest

echo "Running Tests"

for module in $INTEGRATION_TESTS; do 
	echo "Run tests for Module ${module}..."
	make tests -C .terraform/modules/${module} | tee test_output_${module}.log
	terratest_log_parser -testlog test_output_${module}.log -outputdir ./test-reports/${module}/
done

