# Automatically format source code
# Print out coding style mistakes
lint:
	cd 01-go-by-example-basics && go fmt
	cd 02-go-by-example-concurrency && go fmt && golint -set_exit_status

# Examine source code and report suspicious constructs
# xargs:
# - A command that takes output of a command and passes it as argument of another command
# - If no command is specified, "xargs" executes "echo" by default
vet:
	ls 01-go-by-example-basics/*.go | xargs -L 1 go vet
	ls 02-go-by-example-concurrency/*.go | xargs -L 1 go vet

# Generate table of contents
# Keep docs up-to-date from local or remote code sources
docs:
	cd 01-go-by-example-basics && doctoc README.md && md-magic README.md
	cd 02-go-by-example-concurrency && doctoc README.md && md-magic README.md
	cd 03-go-by-example-advance && doctoc README.md && md-magic README.md

# Makefile will get confused if there are files and folders with the names of recipes
# Unless we mark them as 'PHONY'
# @see http://www.gnu.org/software/make/manual/make.html#Phony-Targets
.PHONY: lint vet docs
