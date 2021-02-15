package main

import "fmt"

const untypedConst = "untyped Hello" // untyped constant
const typedConst string = "typed hello" // typed constant

func main() {
	fmt.Printf("untypedConst-> value: %v, type: %T\n", untypedConst, untypedConst)
	fmt.Printf("typedConst-> value: %v, type: %T\n", typedConst, typedConst)

	type MyString string
	var ms string

	// Unlike the typed constants, the untyped constants have no type.
	// Assigning them to a variable of any type compatible with strings works without error
	ms = untypedConst
	fmt.Printf("ms: value: %v, type: %T\n", ms, ms)

	// cannot use typedHello (type string) as type MyString in assignment
	ms = typedHello
	fmt.Printf("ms: value: %v, type: %T\n", ms, ms)

}