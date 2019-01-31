package main

import (
	"fmt"
)

/* AGENDA

- Basics
- Composing interfaces
- Type conversion
  - The empty interface
  - Type switches
- Implementing with values vs. pointers
- Best practices

*/

type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

func main() {
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello GO!!!"))
}
