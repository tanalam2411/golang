package main

import (
	"fmt"
	"reflect"
	"strconv"
)

/*

Agenda:
* Variable declaration
* Redeclaration and shadowing
* Visibility
* Type conversions

*/

//Variable declared at package level can't use := shorthand, and have to use var varName type format
// Scope of var:
// 1) if lower case at module level then the scope is throughout the package
// 2) if upercase then its accessible everywhere (globally)
// 3) Block scope - var declared inside function is scoped within that block
var a float32 = 10

//Var block

// two seperate var blockes have been created just for clearity

var (
	length int     = 10
	width  float32 = 32
)

var (
	class   string = "A"
	ipRange string = "0-255"
)

func main() {

	// Three ways to declare variable
	// First one

	var i int
	i = 100

	var j = 200 //type is inferred based on the type of value provided on the  right side

	k := "Hello"

	fmt.Println(i, j, k)

	x := 100 // will set type of x as int

	fmt.Printf("%v, %T \n", x, x)

	x = 100.0 // will set type of x as int
	fmt.Printf("%v, %T \n", x, x)

	y := 100.0 // will set type of y as float 64

	fmt.Printf("%v, %T \n", y, y)

	// So if we want to create var of type float32 then := wont work, and will have to explicitly give type to var

	var z float32 = 32.4

	fmt.Printf("%v, %T \n", z, z)

	var a float32 = 64 // this is called shadowing as variable a is redeclared inside func main

	fmt.Printf("%v, %T \n", a, a)

	// Naming convention
	// use camel case
	// keep acronym uppercase
	// e.g theHTTP, ipURL

	var fetchURL string = "www.google.com"

	fmt.Printf("%v, %T \n", fetchURL, fetchURL)

	//Type casting
	var ii int = 33
	fmt.Printf("%v, %T \n", ii, ii)

	var jj float32
	jj = float32(ii) //flotat32 to int is possible but will lose the decimal part
	fmt.Printf("%v, %T \n", jj, jj)

	var kk string
	// kk = string(ii) // will return ascii value of the int number, eg: 42 will give *, 33 - !
	// use strconv to do convert (int, float, ...) type to string
	kk = strconv.Itoa(ii)
	fmt.Printf("%v, %T \n", kk, kk)

	fmt.Println(kk, reflect.TypeOf(kk))
}

// Failing Forward
