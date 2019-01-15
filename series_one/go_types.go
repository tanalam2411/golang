package main

import (
	"fmt"
	"reflect"
)

const pi = 3.14

func add_float(x float32, y float32) float32 {
	return x + y
}

// if arguments is of same type, we mention the type only once at the rather repeating.

func add_int(x, y int) int {

	return x + y
}

func print_string(a, b string) (string, string) {
	return a, b
}

func main() {

	var num1 float32 = 5.1
	var num2 float32 = 3.4

	var num3, num4 int = 5, 10

	fmt.Println("Float Addition is:", add_float(num1, num2))

	/*
		if we run add_int with float input values[add_int(1.1, 2.2)] it exists with status 2.
		fmt.Println("Int Addition is:", add_int(num1, num2))

		./go_types.go:24:41: cannot use num1 (type float64) as type int in argument to add_int
		./go_types.go:24:41: cannot use num2 (type float64) as type int in argument to add_int
		``go run go_types.go`` exited: exit status 2

	*/

	/*
		if we call with unknown var, here that is var5
		fmt.Println("Ind Addition is:", add_int(num3, num5))

		./go_types.go:34:41: cannot use num1 (type float64) as type int in argument to add_int
		./go_types.go:34:41: cannot use num2 (type float64) as type int in argument to add_int
		./go_types.go:35:48: undefined: num5
		``go run go_types.go`` exited: exit status 2

	*/
	fmt.Println("Ind Addition is:", add_int(num3, num4))

	// we can also let the complier define the data type of variables
	// but it will by default define float64, so if you need float32
	// you will have to explicitly define float32 type
	num5, num6 := 4.5, 9

	fmt.Println("Data type of var num5 is: ", reflect.TypeOf(num5))
	fmt.Println("Data type of var num6 is: ", reflect.TypeOf(num6))

	// constant
	fmt.Println("Constant value for var pi is: ", pi)

	//string
	wd1, wd2 := "Hello", "world"
	// var wd1, wd2 string = "Hello", "world"

	fmt.Println(print_string(wd1, wd2))

	// Type casting

	fmt.Println("Addition of float64 to 32 and int to float32: ", add_float(float32(num5), float32(num6)))

	// value passing

	x := wd1

	fmt.Println(" x as wd1: ", x)
}
