package main

import (
	"fmt"
)

/* AGENDA
- Basic syntax
- Parameters
- Return values
- Anonymous functions
- Functions as types
*/

/*
Entry point is always from main function within main package.
main function doesn't take any parameter.
*/

/*
Uppercase functions are visible outside the package and lower case functions are visible within the package
*/

func main() {
	greeting, name := "Hello", "Go"
	// Here we're calling sayGreeting func with pass by value
	// Thats why changing the value within sayGreeting function doesn't change greeing variable
	fmt.Println(greeting, name)

	// Trying passing address instead of copy value -> &name
	// Passing pointers is the only way using which variables outside of function scope can be maniplated
	// Passing a pointer is much effecient spicialy in case of large data structrue
	// Be sure that by passing pointer you're allowing the calling function to manipulate or change the original data
	// We don't need to pass pointers for maps and slices because it already shares pointers not the underlying data

	sayGreeting(greeting, &name)
	fmt.Println(greeting, name)

	// very attic parameters
	sum("The sum is: ", 1, 2, 3, 4, 5)

	sum1Result := sum1(1, 2, 3, 4, 5, 6)
	fmt.Println("sum1 result: ", sum1Result)

	sum2Result := sum2(1, 2, 3, 4, 5, 6, 7)
	fmt.Println("sum2 result: ", sum2Result, *sum2Result)

	sum3Result := sum3(1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Println("sum3 result: ", sum3Result)

	divRes, err := divide(5.0, 0.0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("divRes: ", divRes)

	anonymousFunction()

	functionAsVariable()

}

func sayGreeting(greeting string, name *string) {
	/*
		sayGreeting(greeting, name string) -> Here the params both greeting and name will given type string
	*/

	fmt.Println("***********Start- Inside sayGreeting function ************")
	fmt.Println(greeting, *name)
	greeting = "Hi"
	*name = "Python"
	fmt.Println(greeting, *name)
	fmt.Println("***********End- Inside sayGreeting function ************")
}

func sum(msg string, values ...int) {
	// very attic parameters (slice) - is equivalent to python's *args
	// it can only be declared as last argument to function
	// take all the last arguments and wrap them into a slice with the name values in our sum func case
	// you can only have one very attic parameter in a function
	// sum(v ...int, values ...int) -> can only use ... with final parameter in list

	fmt.Println("***********Start- Very attic function ************")
	fmt.Printf("values: %v, type: %T \n", values, values)

	result := 0
	for _, num := range values {
		result += num
	}

	fmt.Println(msg, result)
	fmt.Println("***********End- Very attic function ************")
}

func sum1(values ...int) int {
	// here we're declaring the return type in our function syntax
	result := 0
	for _, num := range values {
		result += num
	}

	return result
}

func sum2(values ...int) *int {
	// here the return type is a int pointer, so as result is in the scope of sum2 fuction but when it is returned to the main function
	// it will be added to the heap memory that means promoted from sum2 function scope to main function scope.
	result := 0
	for _, num := range values {
		result += num
	}

	return &result
}

func sum3(values ...int) (result int) {
	// named return values
	// as we have declared result variable name in our function declaration (result int), simple return at the will return result variable

	for _, num := range values {
		result += num
	}

	return // will reutrn -> result
}

func divide(a, b float64) (float64, error) {
	//multiple value return example
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zerro.")
	}
	return a / b, nil
}

func anonymousFunction() {
	fmt.Println("***********Start- anonymousFunction ************")
	// In this case scope i is within the inner function also.
	for i := 0; i < 5; i++ {
		func() {
			fmt.Println(i)
		}()
	}

	// But accessing outer vairable directly inside of inner function is not safe.
	// It causes in case of asynchronous programming (go routines)
	// So the safe way is to pass it as argument of inner function

	fmt.Println("---")
	for j := 0; j < 6; j++ {
		func(x int) {
			fmt.Println(x)
		}(j)
	}

	fmt.Println("***********End- anonymousFunction ************")
}

func functionAsVariable() {
	// var f func() = func () {}
	f := func() {
		fmt.Println("Inside function f -----")
	}
	f()

	// divide := func(a, b float64) (float64, error) { ... } // will work similar as below
	var divide func(float64, float64) (float64, error)
	divide = func(a, b float64) (float64, error) {
		if b == 0.0 {
			return 0.0, fmt.Errorf("Cannot divide by zero.")
		} else {
			return a / b, nil
		}
	}

	d, err := divide(5.0, 3.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)
}

/*
tan@tan-ThinkPad-E450:~/tanveer/golang/series_two$ go run go_14_functions.go
Hello Go
***********Start- Inside sayGreeting function ************
Hello Go
Hi Python
***********End- Inside sayGreeting function ************
Hello Python
***********Start- Very attic function ************
values: [1 2 3 4 5], type: []int
The sum is:  15
***********End- Very attic function ************
sum1 result:  21
sum2 result:  0xc000014128 28
sum3 result:  36
Cannot divide by zerro.
divRes:  0
***********Start- anonymousFunction ************
0
1
2
3
4
---
0
1
2
3
4
5
***********End- anonymousFunction ************
Inside function f -----
1.6666666666666667
*/

/*
Cannot divide by zerro.
divRes:  0
*/

/* SUMMARY

- Basic syntax
  - func foo() { // () { has to be on same line}
	  ....
	}

- Parameters
  - Commad delimited list of variables and types
	- func foo(bar string, baz int)
  - Parameters of same type list type once
	- func foo(bar, baz int)
  - When pointers are passed in, the function can change the value in the caller
	- This is always true for data of slices and maps

  - Use variadic parameters to send list of same types in
	- Must be last parameter
	- Received as a slice
	- func foo(bar string, baz ...int)

- Return values
  - Single return values just list type
	- func foo() int
  - Multiple return value list types surrouunded by parentheses
	- func foo() (int, error)
	- The (result type, error) paraadigm is a very common idiom
  - Can use named return values
	- Initializes returned variable
	- Return using return keyword on its own
  - Can return addresses of local variables
	- Automatically promoted from local memory(stack) to shared memory(heap)

- Anonymous function
  - Functions don't have names if they are:
	- Immediately invoked
	  - func() {
		  ...
	  } ()
	  - Assigned to a variable or passed as an argument to a function
	    - a := func() {
			...
		}
		a()

- Functions as types
  - Can assign functions to variables or use arguments and return values in functions
  - Type signauture is like function signature, with no parameter names
	- var f func(string, string, int) (int, error)

*/
