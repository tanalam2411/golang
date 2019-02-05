package main

import (
	"fmt"
)

func main() {
	// var wc WriterCloser = myWriterCloser{} => wc is holding value of myWriterCloser,
	// in this case all the methods must have value receiver not pointer receiver
	// so func (mwc myWriterCloser) Write => func (mwc *myWriterCloser) Write // will not work
	// so in case of methods as a value receiver interface must be implemented as a value reciver
	var wc WriterCloser = myWriterCloser{}
	fmt.Println(wc)

	// if any of the method needs to be pointer receiver then interface needs to be implemented as pointer receiver
	// var wc WriterCloser = &myWriterCloser{} // will work with any of the following combination

	// func (mwc *myWriterCloser)
	// func (myc myWriterCloser)

	// func (mwc myWriterCloser)
	// func (myc *myWriterCloser)

	// func (mwc *myWriterCloser)
	// func (myc *myWriterCloser)

	// func (mwc myWriterCloser)
	// func (myc myWriterCloser)
}

type Writer interface {
	Write([]byte) (int, error)
}

type Closer interface {
	Close() error
}

type WriterCloser interface {
	Writer
	Closer
}

type myWriterCloser struct{}

func (mwc myWriterCloser) Write(data []byte) (int, error) {
	return 0, nil
}

func (myc myWriterCloser) Close() error {
	return nil
}

/* BEST PRACTICES

- Use many, small interfaces then large monolith interface
  - Single method interfaces are some of the most powerful and flexible
	- io.Writer, io.Reader, interface{}

- Don't export interfaces for types that will be consumed
- Do export interfaces for types that will be used by package
- Design functions and methods to receive interfaces whenever possible

*/


/* SUMMARY

- Basics 
- Composing interfaces
- Type conversion
- The empty interface and type switches
- Implementing with values vs. pointers
  - Method set of value is all the methods with value receivers
  - Method set of pointer is all methods, regardless of receiver type
  