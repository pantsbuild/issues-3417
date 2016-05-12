#!/usr/bin/env bash

set -o errexit
set -o pipefail

GOPATH=$PWD/src/go go list -f '{{.Dir}} {{.Name}}' ./... | while read dir name
do
  if [[ "${name}" == "main" ]]
  then
    echo "go_binary()" > "${dir}/BUILD"
  else
    echo "go_library()" > "${dir}/BUILD"
  fi
done

mkdir -p 3rdparty/go
./pants buildgen.go
