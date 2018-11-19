package main

import (
	"fmt"
)

/*

& - holds memory address
* - returns value at that address

*/


func main() {
	x := 15
	a := &x // holds memory address of x

	fmt.Println(a)

	fmt.Println(*a) // prints value at address a(i.e address of x's value)

	*a = 10 // will update value at address a i.e will update value of var x

	fmt.Println(*a)

	*a = *a * *a // will do multiplication between values at address a, i.e 10*10

	fmt.Println(*a)
	fmt.Println(x)
}
