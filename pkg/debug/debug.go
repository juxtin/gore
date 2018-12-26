package debug

import (
	"fmt"
	"os"
)

func Print(a ...interface{}) {
	enabled := os.Getenv("DEBUG")
	if enabled != "" {
		fmt.Println(a...)
	}
}
