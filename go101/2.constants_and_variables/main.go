package main

import "fmt"

const Pi = 3.1416

// Declare multiple constants in a group
const (
	No         = !Yes
	Yes        = true
	MaxDegrees = 360
	Unit       = "radian"
)

// Declare multiple constants in one line
const One, Two = 1, 2

var lang = "Go"

var x int  // will set default null value of int i.e., 0

func main() {

	fmt.Printf("%T\n", lang)
	fmt.Println(x)
}