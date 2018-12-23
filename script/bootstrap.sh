#!/bin/bash

set -euo pipefail

go get -d github.com/smacker/go-tree-sitter
pushd $GOPATH/src/github.com/smacker/go-tree-sitter
make
popd

go get -d github.com/stretchr/testify/require
