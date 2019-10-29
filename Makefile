lint:
	go fmt && golint -set_exit_status

# Generate table of contents
# Keep docs up-to-date from local or remote code sources
docs:
	doctoc README.md && md-magic README.md

# Makefile will get confused if there are files and folders with the names of recipes
# Unless we mark them as 'PHONY'
# @see http://www.gnu.org/software/make/manual/make.html#Phony-Targets
.PHONY: lint docs
