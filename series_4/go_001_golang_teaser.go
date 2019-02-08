package main

import (
	"io"
	"log"
	"os"
)

func main() {

	var f, err = os.Open("data.txt")

	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()

	io.Copy(os.Stdout, f)
}
