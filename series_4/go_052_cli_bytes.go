package main

import (
	"bytes"
	"fmt"
)

func main() {

	var b bytes.Buffer

	(&b).Write([]byte("Hello, world!!!\n"))
	fmt.Fprintf(&b, "Holy %s \n", "smokes batman!")
	fmt.Println(b.String())
}