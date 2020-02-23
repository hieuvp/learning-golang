#!/usr/bin/env bash

set -eou pipefail

mapfile -t FILES < <(git ls-files | grep --extended-regexp '.*\.(go)$')

is_ignored() {
  local -r file=$1

  if [[ $file =~ ^unorganized\-.*\/.*$ ]]; then
    echo -en "\e[33m"
    echo "Skipping: $file"
    echo -en "\e[0m"

    return 0
  fi

  return 1
}

process() {
  local -r file=$1

  echo -en "\e[33m"
  echo "Processing: $file"
  echo -en "\e[0m"

  go vet "$file"
  golint -set_exit_status "$file"
}

for file in "${FILES[@]}"; do
  if is_ignored "$file"; then
    continue
  fi

  process "$file"
done
