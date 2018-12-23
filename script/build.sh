#!/bin/bash

set -euo pipefail

output="bin/gore"

go build -o $output pkg/main/main.go
echo "Executable written to ${output}"
