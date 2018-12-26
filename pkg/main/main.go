package main

import (
	"os"

	"github.com/juxtin/gore/pkg/debug"
	"github.com/juxtin/gore/pkg/gore"
)

func main() {
	rootDir := os.Args[1]
	debug.Print("Root dir:", rootDir)
	gopath := os.Getenv("GOPATH")
	gore.Smoke(gopath, rootDir)
}
