package main

import (
	"io"
	"log"
	"os"
)

func main() {

	f, err := os.Open("main.go")

	if nil != err {
		log.Fatalln(err)
	}

	defer f.Close()

	io.Copy(os.Stdout, f)
}