package main

import (
	"fmt"
	"os"

	"github.com/juxtin/gore/pkg/debug"
	"github.com/juxtin/gore/pkg/gore"
)

func main() {
	filename := os.Args[1]
	debug.Print("Filename:", filename)
	gopath := os.Getenv("GOPATH")
	result := gore.BuildGraph(gopath, filename)

	fmt.Println(result)
}
