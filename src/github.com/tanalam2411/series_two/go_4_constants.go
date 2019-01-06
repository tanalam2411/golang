package main

import (
	"fmt"
)

/*
Agenda
	- Naming Convention
	- Typed constants
	- Untyped constants
	- Enumerated constants
	- Enumeration expressions

*/

const shadowConst int16 = 23

const a = iota

func returnsInt() int {
	return 1
}

func main() {
	fmt.Println("***************** CONSTANTS ********************")

	// If const is decalred all upper case as we do in other laguages, then this const would be exported.
	// So make sure to use camel case if you don't want it to be exported.

	// Typed constant
	const myConst int = 33
	fmt.Printf("%v, %T \n", myConst, myConst)

	// myConst = 44 // will thorugh error (cannot assign to myConst) as it is of type constant

	// Imp - value of constant should be assignable at compile time
	// You cann't assign any value that has to be calculated at run time
	//const newConst int = returnsInt() // will through error -> const newConst int = returnsInt()
	// var newVar int = returnsInt()
	// const newConst int = newVar // will through error -> const initializer newVar is not a constant

	// Constant can be of type - int, string, float32, bool etc.
	// Collections(array, map ...) type can't be of constant type

	fmt.Println("***************** Shadow Constant **********************")

	// Constant can be shadowed.
	// Inner constant would be given preference
	// Inner constant can be of any type, doesn't need to be of outter const type

	fmt.Printf("%v, %T \n", shadowConst, shadowConst)
	const shadowConst int = 99
	fmt.Printf("%v, %T \n", shadowConst, shadowConst)

	// Different operation can be performed between const and variable if they are of same type
	var var1 int = 100
	fmt.Printf("%v, %T \n", shadowConst+var1, shadowConst+var1)

	fmt.Println("***************** Infer type of constant **********************")

	const inferedConst = 40
	var num16 int16 = 30
	fmt.Printf("%v, %T \n", inferedConst, inferedConst)

	// So it allows to add inferedConst with int16 var, as what compiler does is it replaces inferedConst with actual literal value of it.
	// so inferedConst + num16 becomes 42 + num16 -> literal + int16.
	// So the type of inferedConst is implicitly converted to int16

	// If type of constant if given then this will not be possbile. i.e const inferedConst int = 40
	// As it will try to add int + int16
	fmt.Printf("%v, %T \n", inferedConst+num16, inferedConst+num16)

	fmt.Println("***************** Enumerated constants **********************")
}
