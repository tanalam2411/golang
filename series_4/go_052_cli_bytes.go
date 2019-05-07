package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {

	var b bytes.Buffer

	(&b).Write([]byte("Hello, world!!!\n"))
	fmt.Fprintf(&b, "Holy %s \n", "smokes batman!")
	fmt.Println(b.String())

	f, err := os.Open("main.go")

	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()
}