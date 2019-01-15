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
	// type of array will return -> [size]type eg: [3]int, [5]string

	fmt.Println("************** Arrays *******************")
	grades := [...]int{90, 93, 97, 85, 66}

	fmt.Printf("Grades: %v, %T \n", grades, grades)

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
	// We don't have to be specify size.
	// Slice is a projection of the underlying array.
	// thats why slice has another ket feature Capacity.
	// Unlike array slices are bydefault reference type. so s2 := s1 will create reference not a sperate copy of data
	// type of slice will return -> []type eg: []int, []string
	// Slices can be updated, elements can be removed, appended

	s1 := []int{1, 2, 3}
	s2 := s1
	s2[1] = 100
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Printf("Length of slice s1: %v %T\n", len(s1), s1)
	fmt.Printf("Capacity of slice s1: %v \n", cap(s1))

	s3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// following slice operation works with arrays too. Try makings3 as array -> s3 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// [start:end] -> start index is inclusive and end index is exclusive
	s4 := s3[:]   // slice of all elements // start index to end index// just like Python
	s5 := s3[3:]  // from index 3 till end [4 5 6 7 8 9 10]
	s6 := s3[:6]  // from index 0 to 6-1 i.e 5 [1, 2, 3, 4, 5, 6]
	s7 := s3[3:6] // from index 3 to 6-1 i.e 5 [4, 5, 6]

	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(s5)
	fmt.Println(s6)
	fmt.Println(s7)

	fmt.Println("************** make function to create slice **********************")

	// make arguments -> type of slice, size, capacity
	// initializes emppty slice with default value 0
	ms1 := make([]int, 3) // will in create [0 0 0], type: []int
	fmt.Println(ms1)
	fmt.Printf("Lenght: %v, %T \n", len(ms1), ms1)
	fmt.Printf("Capacity: %v \n", cap(ms1))

	// Limiting the capacity
	// Capacity should always be larger then size or len of slice or else it will through error - len larger than cap in make([]int)
	// The use Capacity is to avoid updating of the underlying array each time when we update the slice
	ms2 := make([]int, 3, 100)
	fmt.Println(ms2)
	fmt.Printf("Lenght: %v, %T \n", len(ms2), ms2)
	fmt.Printf("Capacity: %v \n", cap(ms2))

	//
	s8 := []int{}
	fmt.Println(s8)
	fmt.Printf("Length: %v\n", len(s8))
	fmt.Printf("Capacity: %v\n", cap(s8))

	//append func can take two or more parameters, this is called variadic-function - https://gobyexample.com/variadic-functions

	s8 = append(s8, 0)
	s8 = append(s8, 1, 2, 3, 4, 5, 6)
	fmt.Println(s8)
	fmt.Printf("Length: %v\n", len(s8))
	fmt.Printf("Capacity: %v\n", cap(s8))

	fmt.Println("************** Appending one slice to another **********************")
	// Appending one slice to another
	ms3 := make([]int, 0, 100)
	ms3 = append(ms3, 1, 2, 3)
	fmt.Println(ms3)
	fmt.Printf("Length: %v\n", len(ms3))
	fmt.Printf("Capacity: %v\n", cap(ms3))

	//ms3 = append(ms3, []int{4, 5, 6}) //cannot use []int literal (type []int) as type int in append
	ms3 = append(ms3, []int{4, 5, 6}...) // unpacking the slice -> append(ms3, 4, 5, 6)
	fmt.Println(ms3)
	fmt.Printf("Length: %v\n", len(ms3))
	fmt.Printf("Capacity: %v\n", cap(ms3))

	fmt.Println("************** Updating slice **********************")
	ms4 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(ms4)
	ms4 = append(ms4[1:4], ms4[5:8]...)
	fmt.Println(ms4)

	ms5 := append(ms4[1:4], ms4[5:8]...)
	fmt.Println(ms5)
	fmt.Println(ms4)
	/*
		ms5 := append(ms4[1:4], ms4[5:8]...)
		ms5 - [3 4 6 8 8 9]
		ms4 - [2 3 4 6 8 8]

		So while updating a slice if it is referenced to other slice unexpected behaviour occurs.
		So use iteration logic to get elements from one slice to other based in index
	*/

}

/* SUMMARY

- Arrays
  - Collection of items of same type
  - Fixed size, can't be changed
  - Declaration styles
	- a := [3]int{1, 2, 3}
	- a := [...]int{1, 2, 3}
	- var a [3]int // will create [0 0 0]
  - Access via zero-based index
	- a := [3]int{1, 2, 3, 4, 5} // a[1] wil return 2
  - len function returns size of array
  - Copies refer to different underlying data // a2 := a1

- Slices
  - Backed by array
  - Creation styles
	- Slice existing array or slice
	- Literal style
	- Via make function
	  - a := make([]int, 10) // create slice with capacity and length == 10
	  - a := make([]int, 10, 100) // slice with length == 10 and capacity == 100
	- len function returns length of slice
	- cap function returns capacity of slice
	- append function to add elements to slice
	  - May cause expensive copy operation if inderlying array is too small
	- Copies refer to same underlying array
*/
