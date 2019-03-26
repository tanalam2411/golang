package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fn := "notes"

	fi, err := os.Stat(fn)

	if err != nil {
		log.Fatalf("%v\n", err)
	}

	switch mode := fi.Mode(); {
	case mode.IsDir():
		fmt.Println("is Dir")
	case mode.IsRegular():
		fmt.Println("is file")
	default:
		fmt.Println("!!!")

	}
}