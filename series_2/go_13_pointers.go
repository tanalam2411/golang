package main

import (
	"fmt"
)

/* AGENDA

- Creating pointers
- Dereferencing pointers
- The new function
- Working with nil
- Types with internal pointers

*/

func main() {

	a := 42
	b := a // creates a seperate copy for b variable

	fmt.Printf("a: %v, type: %T, pointer: %p  \n", a, a, &a)
	fmt.Printf("b: %v, type: %T, pointer: %p  \n", b, b, &b)
	/*
		a: 42, type: int, pointer: 0xc0000140c8
		b: 42, type: int, pointer: 0xc0000140e0
	*/

	var c int = 42
	var d *int = &c                                     // var d int = &c  -> cannot use &c (type *int) as type int in assignment
	fmt.Printf("c: %v: %v, d: %v: %v \n", c, &c, d, *d) // *d -> is called dereferencing the pointer

	c = 43
	fmt.Printf("c: %v, d: %v \n", c, *d)
	*d = 44
	fmt.Printf("c: %v, d: %v \n", c, *d)

	/*
		c: 42: 0xc000014120, d: 0xc000014120: 42

		0xc000014120 -> is the numerical representation of the address holding the value of c
		So d is holding address of c's data i.e 42

		c: 43, d: 43
		c: 44, d: 44
	*/

	e := [...]int{1, 2, 3, 4, 5}
	f := &e[0]
	g := &e[1]
	h := &e[2]
	i := &e[3]
	// Pointer arethmatic operation are not allowed over pointers // incase required checkout - https://golang.org/pkg/unsafe/
	//k := &e[3] + 8 // invalid operation: &e[3] + 8 (mismatched types *int and int)

	fmt.Printf("e: %v %T, f: %p, g: %p, h: %p, i: %p \n", e, e, f, g, h, i)
	// e: [1 2 3 4 5], f: 0xc000018390, g: 0xc000018398, h: 0xc0000183a0, i: 0xc0000183a8

	var ms *myStruct
	ms = &myStruct{foo: 42}
	fmt.Printf("ms: %v, %T \n", ms, ms)
	// ms: &{42 }, *main.myStruct

	var ms1 *myStruct
	// Uninitialized pointer's default value is <nil>
	fmt.Println("uninitialized ms1: ", ms1)
	ms1 = new(myStruct)
	fmt.Println("initialized ms1: ", ms1)
	/*
		uninitialized ms1:  <nil>
		initialized ms1:  &{0 }
	*/
	(*ms1).foo = 100        // () are used around *ms1 because . has more presedence than *
	ms1.bar = "some string" // in this case ms1 is not the struct its pointing to myStruct,
	// but in go it has be syntactically sugared so (*ms1).foo can also be written as ms1.foo
	fmt.Println("ms1: ", ms1)
}

type myStruct struct {
	foo int
	bar string
}

/* SUMMARY

- Creating pointers
  - Pointer types use an asterisk(*) as a prefix to type pointed to
	- *int - a pointer to an integer
  - Use theaddressof operator (&) to get address of variable

- Dereferecing pointers
  - Dereference a pointer by preceding with an asterisk(*)
  - Complex types (e.g. structs) are automatically dereferenced

- Create pointers to objects
  - Can use the addressof operator (&) if value type already exists
	- ms := myStruct{foo: 42}
	  p := &ms
	- Use addressof operator before initializer
	  - &myStruct{foo: 100}
	- Use the new keyword
	  - Can't initialize fields at the same time

- Types with internal pointers
     - All assignment operations in Go are copy operations
IMP  - Slices and maps container poiinters, so copies point to same underlying data
*/
