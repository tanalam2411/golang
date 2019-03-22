package main

import "fmt"

type ReadCounter struct {
	count uint64
}


func main() {
	var rc ReadCounter

	var s0 []byte
	var s1 = make([]byte, 20)
	var s2 = make([]byte, 34)

	rc.Read(s0)
	rc.Read(s1)
	rc.Read(s2)

	fmt.Printf("%v \n", rc)
}


func (r *ReadCounter) Read(b []byte) (n int, err error) {
	r.count += uint64(len(b))
	return len(b), nil
}


func (r ReadCounter) String() string {
	return fmt.Sprintf("%v", r.count)
}