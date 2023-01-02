#!/bin/bash

go install github.com/gruntwork-io/terratest/cmd/terratest_log_parser@latest

for module in $(ls .terraform/modules/); do 
	make tests -C .terraform/modules/${module} | tee test_output_${module}.log
	terratest_log_parser -testlog test_output_${module}.log -outputdir ./test-reports/${module}/
done

