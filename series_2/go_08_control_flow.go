package main

import (
	"fmt"
)

/* AGENDA

- if statements
  - Operators
  - if - else and if - else if statements

- Switch statements
  - Simple cases
  - Cases with multiple tests
  - Falling through
  - Type switches

*/

func main() {

	if true {
		fmt.Println("true worked.")
	}

	// Initializer syntax

	deviceMap := map[string]int{
		"vm":     50,
		"router": 5,
		"switch": 3,
	}

	// Here -> count, ok := deviceMap["vm"]; -> part is the initializer
	if count, ok := deviceMap["vm"]; ok && count > 10 {
		fmt.Println("Device of type 'vm' is present with count: ", count)
	}

	// fmt.Println(count, ok) // will thorugh error because scopr of count and ok is within if block

	number := 50
	guess := 60

	if guess < 1 {
		fmt.Println("The guess must be greater than 1!")
	} else if guess < 1 || guess > 100 {
		fmt.Println("The guess must greater than 0 and less than 100")
	} else {
		if guess < number {
			fmt.Println("Too low")
		} else if guess > number {
			fmt.Println("Too high")
		} else {
			fmt.Println("matched")
		}

	}

	fmt.Println("****************** Switch Case ************************")

	// switch 1 {  // Simple swtich with case 1
	// switch with initializer
	// Scope of i is within switch case
	// default: is optional
	// break keyword is implicit // not required

	switch i := 1 + 3; i {
	case 1:
		// Can have multiple statements
		// do somthing
		fmt.Println("Case 1 selected")
	case 2:
		fmt.Println("Case 2 selected")
	case 3, 4, 5:
		fmt.Println("Case 3 or 4 or 5 selected")
	default:
		fmt.Println("Non of the case matched os excuting the default option")
	}

	// tag less switch case

	i := 10

	switch {
	// In this case both case are valid i <= 10 and also i <= 20, but only first one will be picked(i <= 10)
	case i <= 10:
		fmt.Println("------less than or equal to 10-------")
		fallthrough
		// fallthrough is logic, will just allow the next case to execute even if the case is not valid
	case i >= 20:
		fmt.Println("------less than or equal to 20-------")
	case i <= 30:
		fmt.Println("------less than or equal to 30-------")

	default:
		fmt.Println("------greater than 20-------")
	}

	fmt.Println("****************** Type Switch Case ************************")

	// Any type can be assigned to interface - int, string, float, func, struct ...
	// var t interface{} = 1
	var t interface{} = [3]int{}

	// fmt.Println(" t :- ", t, t.(type)) // use of .(type) outside type switch // so .(type) is only avilable within switch

	switch t.(type) {

	case int:
		fmt.Println("t is an int")
	case float64:
		fmt.Println("t is an float64")
	case string:
		fmt.Println("t is an string")
	case [3]int:
		// note - [3]int is not of same type of [4]int
		fmt.Println("t is an array of int of size 3")

		// break keywork is available
		break
		fmt.Println("this statement will not get executed")
	default:
		fmt.Println("different type")
	}
}

/* SUMMARY

- if statements
  - Initializer
  - Comparision operators (< > <= >= == !=)
  - Logical operators (&& || !)
  - short circuiting - (true || flase || false) -> true, (false && true && true) -> false //same as in python
  - if-else statements
  - if-else if statements
  - Equality and floats

- switch statements
  - Switching on a tag
  - Cases with multiple tests
  - Initializers (will generate tags not boolean)
  - Switches with no tag
  - Fallthrough - is logic less, will just execute the next case
  - Type switches - t.(type)
  - Breakingout early - using break keyword
*/
