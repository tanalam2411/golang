package main

import "fmt"

/*

An "interface" is a type that defines the set of operations/methods available
- A 'type' can implement or provide operations for several interfaces.

Example of interfaces -
Interface		Method

Duck			Quack()
				Waddle()

Stringer		String(): string

Reader			Read(...): len, err

Writer			Write(...): len, err
*/


/* Struct  vs Interface

Struct - Defines the memory layout of a type
         How should a value of this type be stored in memory, the memory footprint.

Interface - Defines which operations are available
            What actions can be taken in manipulation a value which implements this interface.


*/



type Kid struct {
	name string
	age int
}


type IndianRunner struct {
	name string
	weight int
}



func main() {
	var kid1 = Kid{"Max", 30}
	var kid2 = Kid{"Trut", 29}

	printPerson(kid1)
	printPerson(kid2)
}


func printPerson(k Kid) {
	fmt.Printf("Person name: %v, age %v \n", k.name, k.age)
}