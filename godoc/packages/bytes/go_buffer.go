package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"os"
)

func main() {
	exampleOne()
	readerExample()
}

func exampleOne() {
	var b bytes.Buffer  // A Buffer needs no initialization
	b.Write([]byte("hello "))
	fmt.Fprintf(&b, "world! \n\n")
	b.WriteTo(os.Stdout)
}


func readerExample() {
	// A Buffer can turn a string or a []byte into an io.Reader.
	buf := bytes.NewBufferString("AbCdEfGh==")  // returns &Buffer{buf: []byte(s)}
	dec := base64.NewDecoder(base64.StdEncoding, buf)
	fmt.Printf("%", dec)
	io.Copy(os.Stdout, buf)
}
