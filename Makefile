# go fmt - automatically format source code
# go vet - examine source code and report suspicious constructs
# golint - print out coding style mistakes
lint:
	cd 01-go-by-example-basics && go fmt && ls *.go | xargs -L 1 go vet
	cd 02-go-by-example-concurrency && go fmt && golint -set_exit_status

# xargs:
# - Take output of a command and pass it as argument to another command
# - If no command is specified, "xargs" executes "echo" by default

# Generate table of contents
# Keep docs up-to-date from local or remote sources
docs:
	cd 01-go-by-example-basics && doctoc README.md && md-magic README.md
	cd 02-go-by-example-concurrency && doctoc README.md && md-magic README.md
	cd 03-go-by-example-advance && doctoc README.md && md-magic README.md

# Makefile will get confused if there are files and folders with the names of recipes
# Unless we mark them as 'PHONY'
# @see http://www.gnu.org/software/make/manual/make.html#Phony-Targets
.PHONY: lint docs
