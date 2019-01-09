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

	// grades := [5]int{90, 93, 97, 85, 66}
	// ... -> Creates array of size provided num of literals. ({90, 93, 97, 85, 66} -> 5)
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

}
