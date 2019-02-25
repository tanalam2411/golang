package main

import "fmt"

/*
key - only those type which supports comparision, like int, string, float but not function or slice

Construct - m := map[key_type]value_type{}
Insert - m[k] = v
Lookup - v = m[k]
Delete - delete(m, k)
Iterate - for k, v := range m
Length - len m
*/


func main() {
	var m map[int]int

	fmt.Printf("m: %v, len: %v \n", m, len(m))

	v := m[5]
	fmt.Println("m[5]: ", v)
	// m[4] = 100 // will trow as error - panic: assignment to entry in nil map

	m = make(map[int]int)
	m[0] = 101
	m[1] = 102
	m[2] = 103
	fmt.Println(m)
	delete(m, 1)
	fmt.Println(m)

}