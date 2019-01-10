package main

import (
	"fmt"
)

/*
Collection Types - Arrays and slices

AGENDA

- Arrays
  - Creation
  - Built-in functions
  - Working with arrays

- Slices
  - Creation
  - Built-in functions
  - Working with slices
*/

func main() {

	// Arrays - are of fixed size and has to be known at compile time.
	// grades := [5]int{90, 93, 97, 85, 66}
	// ... -> Creates array of size provided num of literals. ({90, 93, 97, 85, 66} -> 5)

	fmt.Println("************** Arrays *******************")
	grades := [...]int{90, 93, 97, 85, 66}

	fmt.Printf("Grades: %v \n", grades)

	fmt.Println("************** Array intialized with no values *******************")

	var students [3]string
	fmt.Printf("Students: %v \n", students)

	// Add value to array students
	students[0] = "Max"
	students[1] = "Ahmed"
	students[2] = "Arnold"
	fmt.Printf("Students: %v, len: %v\n", students, len(students)) // So if var students [5]string but we only have 3 values in it, will return len 5 not 3

	//Get values by index
	fmt.Printf("Student #1: %v \n", students[1])

	// 2D array
	var identityMatrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}} // var identityMatrix = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}} // is also correct as it will assign type based on the data on right side
	fmt.Printf("identityMatrix: %v \n", identityMatrix)

	// identityMatrix can also be declared as:

	var identityMatrixNew [3][3]int
	identityMatrixNew[0] = [3]int{1, 0, 0}
	identityMatrixNew[1] = [3]int{0, 1, 0}
	identityMatrixNew[2] = [3]int{0, 0, 1}
	fmt.Printf("identityMatrixNew: %v \n", identityMatrixNew)

	fmt.Println("************************ Arrays in Go are values *******************************")

	a1 := [...]int{1, 2, 3, 4, 5}
	a2 := a1 // Will create a new array a2 with seperate copy of data. It doesn't creates reference
	a2[4] = 7
	fmt.Println(a1)
	fmt.Println(a2)

	// Create reference(Pointer address) instead of creating a seperate copy. So this will update both a1 and a3
	a3 := &a1
	a3[1] = 999
	fmt.Println(a1)
	fmt.Println(a3)

	// Slices
	fmt.Println("************** Slices *******************")

}
