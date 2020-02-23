#!/usr/bin/env bash

set -eoux pipefail

# go vet ./...
# golint -set_exit_status ./...

# cd 01-go-by-example-basics && go fmt && ls *.go | xargs -L 1 go vet
#	cd 02-go-by-example-concurrency && go fmt && golint -set_exit_status

# xargs:
# - Take output of a command and pass it as argument to another command
# - If no command is specified, "xargs" executes "echo" by default
