package main

import (
	"fmt"
	"reflect"
)

/*

*/


func zeroValues(){
	var boolean bool
	var integer int
	var float float32
	var stringv string
	var pointer = &struct{}{}
	var function func()
	var interfacev interface{}
	var slice []int
	var channel = make(chan int)
	var mapv map[string]string

	boolean = true
	fmt.Printf("boolean-> value: %v, type: %T\n", boolean, boolean)
	fmt.Printf("integer-> value: %v, type: %T\n", integer, integer)
	fmt.Printf("float-> value: %v, type: %T\n", float, float)
	fmt.Printf("stringv-> value: %v, type: %T\n", stringv, stringv)
	fmt.Printf("pointer-> value: %v, type: %T\n", pointer, pointer)
	fmt.Printf("function-> value: %v, type: %T\n", function, function)
	fmt.Printf("interfacev-> value: %v, type: %T\n", interfacev, interfacev)
	fmt.Printf("slice-> value: %v, type: %T\n", slice, slice)
	fmt.Printf("channel-> value: %v, type: %T\n", channel, channel)
	fmt.Printf("mapv-> value: %v, type: %T\n", mapv, mapv)

	// check nil type
	mapv = make(map[string]string)
	mapv["k"] = "v"
	fmt.Printf("=========== %v, %v, %v\n", mapv, reflect.Zero(reflect.TypeOf(mapv)), )

	if reflect.ValueOf(mapv).IsNil() {
		fmt.Println("**************************")
	}

}

func main() {
	// Zero Values
	fmt.Println("================== Zero Values ==========================")
	zeroValues()
}

