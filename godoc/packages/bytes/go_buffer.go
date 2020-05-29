package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	exampleOne()
}

func exampleOne() {
	var b bytes.Buffer  // A Buffer needs no initialization
	b.Write([]byte("hello "))
	fmt.Fprintf(&b, "world!")
	b.WriteTo(os.Stdout)
}
