package main

import "fmt"

// will not throw unused error
// declare & assign = initialization
var x = 43
var y int = 43 // explicitly providing the type

// Declaring Identifier z - it will assigned the default value of int. i.e., 0
// https://golang.org/ref/spec#The_zero_value
// When storage is allocated for a variable, either through a declaration or a call of new,
// or when a new value is created, either through a composite literal or a call of make,
// and no explicit initialization is provided, the variable or value is given a default value.
// Each element of such a variable or value is set to the zero value for its type: false for booleans,
// 0 for numeric types, "" for strings, and nil for pointers, functions, interfaces, slices, channels, and maps.
// This initialization is done recursively, so for instance each element of an array of structs will have its
// fields zeroed if no value is specified.
var z int

func main() {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
