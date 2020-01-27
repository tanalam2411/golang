package main

import (
	"fmt"
	"reflect"
)

type Vector struct {
	x int
}

func main() {

	v1 := new(Vector)
	v2 := &Vector{}

	fmt.Println("Type of vector v1: ", reflect.TypeOf(v1))
	fmt.Println("Type of vector v2: ", reflect.TypeOf(v2))

	createNewOfSlice()
}


func createNewOfSlice() {

	i := new([]int)
	j := make([]int, 0)

	fmt.Println("Type of i: ", i, reflect.TypeOf(i), *i == nil)
	fmt.Println("Type of j: ", j, reflect.TypeOf(j), *(&j) == nil)

	fmt.Println("Is i a pointer? ", reflect.TypeOf(i).Kind() == reflect.Ptr)
	fmt.Println("Is j a pointer? ", reflect.TypeOf(j).Kind() == reflect.Ptr)
	fmt.Println("Is j a slice? ", reflect.TypeOf(j).Kind() == reflect.Slice)
}