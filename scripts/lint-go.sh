#!/usr/bin/env bash

set -eou pipefail

mapfile -t FILES < <(git ls-files | grep --extended-regexp '.*\.(go)$')

for file in "${FILES[@]}"; do
  echo -en "\033[33m"
  echo "Processing: $file"
  echo -en "\033[0m"

  go vet "$file"
  golint -set_exit_status "$file"
done
