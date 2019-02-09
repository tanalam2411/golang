package main

import "fmt"

/* Function argument list

1. No parameter - func foo()
2. A single parameter - func foo(a int)
3. One or more parameter of same type - func foo(a, b, c float32)
4. One or more different type parameters - func foo( a int, b string, c float32)
5. Optional variadic parameter - func foo(a int, b string, ...z float32)
 */

 /* Function Return List
 1. No return - func foo(...)
 2. A single unnamed return - func foo(...)int
 3. A single named return - func foo(...)a int
 4. One or more parameter of the same type
 func foo(...) (a, b, c float32)
 5. One or more different parameters
 func foo(...) (a int, b string, c float32)
  */

  /* Variable inside of Functions
  - Functions can have variables definition
  - Variables defined in a function, is visible inside the function
  - Unline C/C++. function can return the address of their variables

  */

func main() {
	piAddr := addressOfLocalVar()

	fmt.Println(piAddr)
}

func addressOfLocalVar() *float64 {
	pi := 3.14
	return &pi
}


