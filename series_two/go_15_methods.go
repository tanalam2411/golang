package main

import (
	"fmt"
)

/* AGENDA
- Methods
*/

type greeter struct {
	greeting string
	name     string
}

// value receiver method, as greeter g would be passed as value - value receiver
func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
	g.name = "this will not reflect"
}

// reference receiver method, as greeter g would be passed as a pointer receiver
func (g *greeter) greetPointer() {
	fmt.Println(g.greeting, g.name)
	g.name = "Changed name"
}

// func (g greeter) or func (g * greeter) - are called here g is receiver for the method
func main() {
	g := greeter{
		greeting: "Hello",
		name:     "Go",
	}

	g.greet()
	fmt.Println("g.name : ", g.name)

	g.greetPointer()
	fmt.Println("g.name : ", g.name)
}

/* SUMMARY

- Methods
  - Function that executes in context of a type
  - Format
    - func (g greeter) greet() {
		.....
	}
  - Receiver can be value or pointer
	- Value receiver gets copy of type
	- pointer receiver gets pointer to type

*/
