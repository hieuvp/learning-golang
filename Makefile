# Automatically format source code
# Print out style mistakes
lint:
	cd 01-go-by-example-basics && go fmt && golint -set_exit_status
	cd 02-go-by-example-concurrency && go fmt && golint -set_exit_status
	cd 03-go-by-example-advance && go fmt && golint -set_exit_status

# Generate table of contents
# Keep docs up-to-date from local or remote code sources
docs:
	cd 01-go-by-example-basics && doctoc README.md && md-magic README.md
	cd 02-go-by-example-concurrency && doctoc README.md && md-magic README.md
	cd 03-go-by-example-advance && doctoc README.md && md-magic README.md

# Makefile will get confused if there are files and folders with the names of recipes
# Unless we mark them as 'PHONY'
# @see http://www.gnu.org/software/make/manual/make.html#Phony-Targets
.PHONY: lint docs
