package main

import "fmt"

/*
- An array in Go lang is a built in datatype with the following features:

	- Stores server values of the same type
	- Provides compile-time bounds checking on integer literal index
	- Provides runtime-time bounds checking on integer variable index
	- Provides non-negative integer index
	- Uses 0-base indexing
*/

const size int8 = 10


type Date struct {
	year int
	month int
	day int
}

type Person struct {
	dob Date
	name string
	age int
}


func main() {
	printTestResults()
	readPerson()
}

func printTestResults() {
	// arrays are 0 based index
	// non-negative index are not supported
	var testResult [size]int // [0 0 0 0 0 0 0 0 0 0]

	// Provides compile-time bounds checking on integer literal index
	//testResult[10] = 100

	// Provides runtime-time bounds checking on integer variable index
	//i := 12
	//testResult[i] = 200

	fmt.Printf("testResult: %#v \n\n", testResult)
	testResult = [10]int{1,2,3,4,5,6,7,8,9,10}

	fmt.Printf("testResult: %#v \n\n", testResult)

	var compNumbers = [size]complex64{10, 20, 30}
	fmt.Printf("compNumbers: %#v\n", compNumbers)

	var cityNames = [size]string{}
	cityNames[0] = "Pune"
	fmt.Printf("city names: %#v\n", cityNames)
}


func readPerson() {
	var people [10]Person

	fmt.Printf("Person: %#v \n\n", people[0])
}