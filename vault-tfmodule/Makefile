mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
current_dir := $(notdir $(patsubst %/,%,$(dir $(mkfile_path))))

tests:
	@echo "Running tests for $(current_dir)"
	(cd test/terratest && go test -v ./...)
