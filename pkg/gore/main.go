package main

import (
	"fmt"

	"github.com/juxtin/gore/pkg/gore"
)

import "os"

func main() {
	filenames := os.Args[1:]
	results, err := gore.Run(filenames)
	if err != nil {
		fmt.Println("ERROR!")
		fmt.Println("   ", err)
		return
	}
	fmt.Println("Imports:", results)
	fmt.Println("All done")
}
