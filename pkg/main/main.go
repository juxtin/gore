package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/juxtin/gore/pkg/debug"
	"github.com/juxtin/gore/pkg/gore"
)

func die(msg string) {
	fmt.Fprintf(os.Stderr, msg+"\n")
	os.Exit(1)
}

func usage() {
	die("Usage:\n\tgore path/to/project/root")
}

func main() {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		die("GOPATH must be set!")
	}
	if len(os.Args) != 2 {
		usage()
	}
	arg := os.Args[1]
	rootDir, err := filepath.Abs(arg)
	if err != nil {
		die("Unable to find absolute path to '" + arg + "': " + err.Error())
	}
	debug.Print("Root dir:", rootDir)
	gore.Smoke(gopath, rootDir)
}
