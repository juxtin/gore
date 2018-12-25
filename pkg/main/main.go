package main

import (
	"fmt"

	rel "github.com/juxtin/gore/pkg/relationships"
)

import "os"

func main() {
	filenames := os.Args[1:]
	results, err := rel.ImportsInFile(filenames)
	if err != nil {
		fmt.Println("ERROR!")
		fmt.Println("   ", err)
		return
	}
	fmt.Println("Imports:", results)
	fmt.Println("All done")
}
