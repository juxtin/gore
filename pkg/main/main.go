package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/juxtin/gore/pkg/debug"
	"github.com/juxtin/gore/pkg/gore"
)

func main() {
	arg := os.Args[1]
	rootDir, err := filepath.Abs(arg)
	if err != nil {
		fmt.Println("Oops:", err)
	}
	debug.Print("Root dir:", rootDir)
	gopath := os.Getenv("GOPATH")
	gore.Smoke(gopath, rootDir)
}
