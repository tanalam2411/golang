package main

import (
	"bytes"
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

/*
Naming convension for interfaces having just one method :
  - Just add "er" to method name to have get interface name
  - method name "Write" - interface name - "Writer"
*/
type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

// Any type can have method assosited with it

// ---------------------Interface using type int --------------

type Incrementer interface {
	Increment() int
}

type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}

// ---------------------- Interface with more than one methods ----------------------

// type Writer interface {
// 	Write([]byte) (int, error)
// }

type Closer interface {
	Close() error
}

type WriterCloser interface {
	Writer
	Closer
}

type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)
	if err != nil {
		return 0, err
	}

	v := make([]byte, 8)
	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}

		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}

	return n, nil
}

func (bwc *BufferedWriterCloser) Close() error {

	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}

func NewBufferedWriterCloser() *BufferedWriterCloser {
	return &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}
func main() {
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello GO!!!"))

	// ---- Interface using type int

	myInt := IntCounter(0)
	var inc Incrementer = &myInt
	for i := 0; i < 5; i++ {
		fmt.Println(inc.Increment())
	}

	// ---- Interface with more than one methods
	var wc WriterCloser = NewBufferedWriterCloser()
	wc.Write([]byte("Hello Go !!!, this will be read in eight byte format."))
	wc.Close()

	// --- Type conversion
	r, ok := wc.(*BufferedWriterCloser)
	if ok {
		fmt.Println(r)
	} else {
		fmt.Println("Conversion failed.")
	}

	// Empty Interface
	// empty interface can be type casted to any other type
	var newObj interface{} = NewBufferedWriterCloser()
	if wc, ok := newObj.(WriterCloser); ok {
		wc.Write([]byte("Some new text ...!!!***"))
		wc.Close()
	}

	// type switching
	var i interface{} = 0
	switch i.(type) {
	case int:
		fmt.Println("i is an integer")
	case string:
		fmt.Println("i is an string")
	default:
		fmt.Println("i is aunknown")
	}

}
