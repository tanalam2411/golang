package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Say hello...")
	stdOut()
}


func stdOut() {
	myStr := ""
	arguments := os.Args
	if len(arguments) == 1 {
		myStr = "Please give me one argument.!"
	} else {
		myStr = arguments[1]
	}

	io.WriteString(os.Stdout, myStr)
	io.WriteString(os.Stdout, "\n")
}