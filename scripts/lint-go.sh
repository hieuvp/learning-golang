#!/usr/bin/env bash

set -eou pipefail

mapfile -t FILES < <(git ls-files | grep --extended-regexp '.*\.(go)$')

for file in "${FILES[@]}"; do
  echo -en "\e[33m"
  echo "Processing: $file"
  echo -en "\e[0m"

  go vet "$file"
  golint -set_exit_status "$file"
done
