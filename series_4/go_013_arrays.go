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
	year  int
	month int
	day   int
}

type Person struct {
	dob  Date
	name string
	age  int
}

func main() {
	printTestResults()
	readPerson()

	arr1, arr2 := arraySender()
	arr3 := arr1 // will create a separate copy of arr1 and assign it to arr3, so arr1 will not change by assignment(next line)
	arr3[0] = 999
	arr4 := arr2 // arr4 will point to the same addr of arr2, so arr2 will change by assignment(next line)
	arr4[0] = 1000
	arrayReceiver(arr3, arr4)

	fmt.Printf("main, arr1: %v, %T, %v \n", arr1, arr1, &arr1)
	fmt.Printf("main, arr2: %v, %T, %v, %v \n", arr2, arr2, &arr2, &arr4)

	interateArray()
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
	testResult = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Printf("testResult: %#v, len: %v \n\n", testResult, len(testResult))
	fmt.Printf("testResult 0-4: %#v \n\n", testResult[0:4])

	var slice1 = testResult[:5]
	fmt.Printf("slice1: %#v, Type: %T \n\n", slice1, slice1)
	slice1[0] = 11111
	fmt.Println("----------- slice1[0] = 11111 // so slice is a reference of the array-------------")
	fmt.Printf("slice1: %#v \n\n", slice1)
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

func arraySender() ([10]int, *[10]int) {

	arr1 := [10]int{10, 20, 30, 40, 50}
	arr2 := [10]int{60, 70, 80, 90, 100}

	fmt.Printf("arraySender, arr1: %v, %T, %v \n", arr1, arr1, &arr1)
	fmt.Printf("arraySender, arr2: %v, %T, %v \n", arr2, arr2, &arr2)
	return arr1, &arr2
}

func arrayReceiver(arr1 [10]int, arr2 *[10]int) {

	fmt.Printf("arrayReceiver, arr1: %v, %T, %v \n", arr1, arr1, &arr1)
	fmt.Printf("arrayReceiver, arr2: %v, %T, %v \n", arr2, arr2, &arr2)
}

func interateArray() {

	arr1 := [5]int{11, 22, 33, 44, 55}

	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}

	for _, v := range arr1 {
		fmt.Println(v)
	}
}
