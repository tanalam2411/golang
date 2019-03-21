package main

import "fmt"

type ByteCounter int


func main() {
	var bc ByteCounter

	var d0 = []byte{0}
	var d1 = []byte{1, 2, 3}
	var d2 = []byte("abcdef ,.? #@ 0")

	fmt.Println("d2: ", d2)

	bc.Write(d0)
	bc.Write(d0)
	bc.Write(d1)
	bc.Write(d2)
	fmt.Println("Byte Counter: ", bc)

}


func (bc  *ByteCounter) Write(b []byte) (n int, err error) {
	l := len(b)
	*bc = ByteCounter(int(*bc) + l)
	return l, nil
}