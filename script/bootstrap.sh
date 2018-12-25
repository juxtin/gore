#!/bin/bash

set -euo pipefail

echo "downloading and building go-tree-sitter"
go get -d github.com/smacker/go-tree-sitter
pushd $GOPATH/src/github.com/smacker/go-tree-sitter
git checkout 12d486660a848e0808308f02c65fab18a607298d
make
popd

echo "downloading gographviz"
go get -d github.com/awalterschulze/gographviz

echo "downloading testify"
go get -d github.com/stretchr/testify/require
