#!/bin/bash

for module in $(ls .terraform/modules/); do 
	make tests -C .terraform/modules/$module/vault-tfmodule; 
done

