package main

import (
	"fmt"
)

/*
https://golang.org/ref/spec#Constants
https://stackoverflow.com/questions/22624124/what-does-it-mean-for-a-constant-to-be-untyped-in-golang
https://stackoverflow.com/questions/14426366/what-is-an-idiomatic-way-of-representing-enums-in-go
Agenda
	- Naming Convention
	- Typed constants
	- Untyped constants
	- Enumerated constants
	- Enumeration expressions

*/

const shadowConst int16 = 23

// iota is a counter used when we create enumerated constants
type Base float32

const (
	// constantsA float32 will not change constantsB and constantsC's type to float32
	constantsA Base = iota // resets to 0
	constantsB      = iota // increaments to 1
	constantsC      = iota // increaments to 2
)

/*
const (
	_ = iota			  // _ is the only write only variable which is being ignored by the compiler
	constantsD 			  // increaments to 1
	constantsE            // increaments to 2
	constantsF            // increaments to 3
)
*/
const (
	constantsD = iota + 1 // resets to 0 as iota is scoped to a constant block
	constantsE            // increaments to 2
	constantsF            // increaments to 3
)

const (
	_ = iota
	// KB - left shift -> 1 << 10*1 -> 2^10
	KB = 1 << (10 * iota)
	// MB - left shift -> 1 << 10*2 -> 2^20
	MB
	// GB - Untyped constant are implicitly converted to the type of type variable if possbile
	// Suppose const x = 10 and var y float32 = 20.5
	// then x + y will not through any error because of different type instead will result 30.5
	// const x = int(10) // will thorugh error on x + y -> invalid operation: x + y (mismatched types int and float32)
	GB
	TB
	PB
	EB
	ZB
	YB
)

const (
	// We did left shift so that each const will have one 1 bit in the binary format for each 00000001, 00000010, 00000100, 00001000, 00010000, 00100000, 01000000, 10000000
	// So that we can easly do bit wise operation (& | ^ &^)
	isAdmin          = 1 << iota // 1 -> 00000001
	isHeadquarters               // 2 -> 00000010
	canSeeFinancials             // 4 -> 00000100

	canSeeAfrica       // 8 -> 00001000
	canSeeAsia         // 16 -> 00010000
	canSeeEurope       // 32 -> 00100000
	canSeeNorthAmerica // 64 -> 01000000
	canSeeSouthAmerica // 128 -> 10000000
)

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
	fmt.Println("---------------- iota -------------------")

	fmt.Printf("%v, %T \n", constantsA, constantsA)
	fmt.Printf("%v, %T \n", constantsB, constantsB)
	fmt.Printf("%v, %T \n", constantsC, constantsC)

	fmt.Printf("%v, %T \n", constantsD, constantsD)
	fmt.Printf("%v, %T \n", constantsE, constantsE)
	fmt.Printf("%v, %T \n", constantsF, constantsF)

	fmt.Printf("%v, %T \n", KB, KB)
	fmt.Printf("%v, %T \n", MB, MB)
	fmt.Printf("%v, %T \n", GB, GB)

	filesize := 4000000000.
	res := filesize / GB
	fmt.Printf("Filesize : %v, %T \n", filesize, filesize)
	fmt.Printf("GB : %v, %T \n", GB, GB)
	fmt.Printf("%.2fGB, %T \n", res, res)

	fmt.Println("******************* Optimized Flag using single byte *********************")

	// Suppose roles is the role allocated to a particular user
	var roles byte = isAdmin | canSeeFinancials | canSeeEurope // 00000001 | 00000100 | 00100000 -> 00100101
	fmt.Printf("%b \n", roles)
	fmt.Printf("%b \n", canSeeSouthAmerica)

	// Want to check if the user is admin we can use & operator
	fmt.Printf("Is admin? %v \n", isAdmin&roles == isAdmin)
	// Check if it is Headquarters
	fmt.Printf("Is HQ? %v \n", isHeadquarters&roles == isHeadquarters)

}

/*
Summary

- Immutable(can not be changed), but can be shadowed (shadowed constant can change both its value and its type)

- Replaced by the compiler at compile time
  - Value must be calculable at compile time
  - So we can't use functions or command line argument to assign constant a value
  - But we can use expressions like bitwise or arithmetic etc (Enumerated expressions)

- Named like variables
  - PascalCase for exprted constants // will be accessible outside the package
  - camelCase for internal constants // will be scoped within the package

- Typed constants work like immutable variables
  - Can interoperate only with same type // const c = int(10)

- Untyped constants work like literals
  - Can interoperate with similar types // const c = 10, var v float = 20.5 -> will work as c will be replaced with actual literal value 10

- Enumerated Constants
  - Special symbol iota allows related constants to be created easily
  - iota starts at 0 in each const block and increments by one
  - Watch out of constant values that match zero values of variables

- Enumerated expressions
  - Operations that can be determined at compile time are allowed
	- Arithmetic
	- Bitwise operations
	- Bitshifting

Failing Forward
*/
