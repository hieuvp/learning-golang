include .env

.PHONY: fmt
fmt:
	@printf "\n"
	$(SCRIPT_PATH)/fmt-shell.sh
	@printf "\n"

	@printf "\n"
	$(SCRIPT_PATH)/fmt-go.sh
	@printf "\n"

	@printf "\n"
	$(SCRIPT_PATH)/fmt-markdown.sh
	@printf "\n"

.PHONY: lint
lint:
	@printf "\n"
	$(SCRIPT_PATH)/lint-shell.sh
	@printf "\n"

	@printf "\n"
	$(SCRIPT_PATH)/lint-go.sh
	@printf "\n"

	cd 01-go-by-example-basics && go fmt && ls *.go | xargs -L 1 go vet
	cd 02-go-by-example-concurrency && go fmt && golint -set_exit_status

# xargs:
# - Take output of a command and pass it as argument to another command
# - If no command is specified, "xargs" executes "echo" by default

.PHONY: git-add
git-add: fmt lint
	@printf "\n"
	git add .
	@printf "\n"
