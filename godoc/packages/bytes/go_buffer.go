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
	newBufferExample()
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
	//fmt.Printf("%s", dec)
	io.Copy(os.Stdout, dec)
}


func newBufferExample() {
	_b := make([]byte, 0, 1)
	buf := bytes.NewBuffer(_b)
	buf.Write([]byte("abc"))
	buf.Write([]byte{97, 98, 99})
	buf.WriteString("\nnewBufferExample: Hello, world\n")
	buf.WriteTo(os.Stdout)
}