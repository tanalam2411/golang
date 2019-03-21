package main

import "fmt"

type Counter int


func main() {
	var c Counter

	var d0 []byte
	var d1 []byte
	var d2 []byte

	c.Write(d0)
	c.Write(d1)
	c.Write(d2)

	fmt.Println("Writer Counter: ", c)
}


func (c *Counter) Write(b []byte)(n int, err error) {
	*c = Counter(int(*c) + 1)
	return len(b), nil
}