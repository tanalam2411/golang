package main

import (
	"errors"
	"fmt"
)

type person struct {
	name string
	age uint8 // 0 - 127
}


type AgeAverager struct {
	sum uint64
	count uint64
}


func main() {

	var bb byte
	bb = 255 // > 255 // constant 256 overflows byte
	fmt.Println("bb: ", bb)

	a := person{"a", 23}
	b := person{"b", 56}
	c := person{"c", 19}

	var aa AgeAverager

	aa.Write(a.AgeToByte())
	aa.Write(b.AgeToByte())
	aa.Write(c.AgeToByte())
	var p []byte
	n, err := aa.Write(p)
	fmt.Println(n, err)

	n, err = aa.Write([]byte{200})
	fmt.Println(n, err)

	fmt.Printf("Average Age: %v\n", aa)
}


func (p person) AgeToByte() []byte {
	var buf = []byte{byte(p.age)}
	return buf
}


func (c *AgeAverager) Write(b []byte) (n int, err error) {
	l := len(b)

	switch {

	case b == nil:
		return 0, errors.New("nil is not valid age")

	case b[0] > 127:
		return int(b[0]), errors.New("too old")
	}

	c.count++
	c.sum += uint64(b[0])

	return l, nil
}


func (a AgeAverager) String() string {
	return fmt.Sprintf("%v", a.sum/a.count)
}