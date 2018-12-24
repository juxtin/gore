#!/bin/bash

set -euo pipefail

# this is bad and wrong, but it works for now
go test test/smoke/smoke_test.go
