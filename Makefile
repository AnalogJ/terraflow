


# targets wrapping Terraform binary
plan: clean-state prepare
	@echo "plan '$(comp_name)' component in '$(env_name)' environment"
	terraform plan \
	-input=false \
	-refresh=true \
	-var-file=config/environments/$(env_name).tfvars \
	-var-file=config/components/$(comp_name).tfvars \
	-var='env_name=$(env_name)' \
	-var='comp_name=$(comp_name)' \
	components/$(comp_name)

apply: clean-state prepare
	@echo "create '$(comp_name)' component in '$(env_name)' environment"
	terraform apply \
	-input=false \
	-var-file=config/environments/$(env_name).tfvars \
	-var-file=config/components/$(comp_name).tfvars \
	-var='env_name=$(env_name)' \
	-var='comp_name=$(comp_name)' \
	-target='$(target_name)' \
	components/$(comp_name)

destroy: clean-state prepare
	@echo "destroy '$(comp_name)' component in '$(env_name)' environment"
	terraform destroy \
	-input=false \
	-var-file=config/environments/$(env_name).tfvars \
	-var-file=config/components/$(comp_name).tfvars \
	-var='env_name=$(env_name)' \
	-var='comp_name=$(comp_name)' \
	-target='$(target_name)' \
	components/$(comp_name)

# target to prepare the workspace
prepare:
	@echo "initialize terraform project"
	@test $${env_name?Please set environment variable "env_name" ("make ACTION env_name=envfoo")}
	@test $${comp_name?Please set environment variable "comp_name" ("make ACTION comp_name=compfoo")}
	@test -d components/$(comp_name) || (echo "Component '$(comp_name)' does not exist." && exit 1)
	terraform init \
	-backend=true \
	-force-copy \
	-get-plugins=true \
	-input=false \
	components/$(comp_name)


# targets to create Terraflow folder structure
project:
	@echo "init folder structure"
	mkdir -p config/{environments,components}
	mkdir -p components

component: project
	@test $${comp_name?Please set environment variable "comp_name" ("make component comp_name=compfoo")}
	mkdir -p components/$(comp_name)
	@test -f config/components/$(comp_name).tfvars || touch config/components/$(comp_name).tfvars
	@test -f components/$(comp_name)/main.tf || touch components/$(comp_name)/main.tf
	@test -f components/$(comp_name)/outputs.tf || touch components/$(comp_name)/outputs.tf
	@test -f components/$(comp_name)/secrets.tf || touch components/$(comp_name)/secrets.tf
	@test -f components/$(comp_name)/security.tf || touch components/$(comp_name)/security.tf
	@test -f components/$(comp_name)/variables.tf || touch components/$(comp_name)/variables.tf

environment: project
	@test $${env_name?Please set environment variable "env_name" ("make environment env_name=envfoo")}
	@test -f config/environments/$(env_name).tfvars || touch config/environments/$(env_name).tfvars


# helper targets
clean-state:
	@echo "cleaning directory"
	rm -rf .terraform
