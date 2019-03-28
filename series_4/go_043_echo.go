package main

import (
	"fmt"
	"os"
)

var lastChar = "\n"


func main() {
	offset := 1

	if len(os.Args) >= 2 && os.Args[offset] == "\n" {
		lastChar = "-"
		offset++
	}

	for _, a := range os.Args[offset:] {
		fmt.Print(a, " ")
	}

	fmt.Print(lastChar)
}