package main

import (
	"fmt"
)

/*
Agenda - Primitives type
- Boolean type
- Numeric types
  - Integers
  - Floating point
  - Complex numbers
- Text types

For more checkout - https://golang.org/pkg/builtin/
*/

func main() {

	// Boolean types - true/false
	var b bool = true

	fmt.Printf("%v, %T \n", b, b)

	n := 10 == 10
	m := 11 == 12

	fmt.Printf("%v, %T \n", n, n)
	fmt.Printf("%v, %T \n", m, m)

	//Default value for bool is false
	var flag bool
	fmt.Printf("%v, %T \n", flag, flag)

	// numeric type
	// Integer type (Signed numbers are both positive nad negative, Unsigned numbers are always positive numbers)
	// int range:
	// int8		-	-128	    					+127
	// int16	- 	-32768							+32767
	// int32	- 	-2 147 483 648					+2 147 483 647          	// Approx neg to pos 2 Billion
	// int64	- 	-9 223 372 036 854 775 808		+9 223 372 036 854 775 807  // Approx neg to pos 9 Quintillion

	// Also have unsigned int - uint, uint8, uint16, uint32, uint64, uintptr
	// uint8	-	0 - 255
	// uint16	-	0 - 65 535
	// uint32	-	0 - 4 294 967 295

	// byte - is a unsigned 8bit int type (alias for uint8)
	// rune - is a alias for int32(represents a unicode point)

	// Operation on same type of different range will through erro - int + int8 not possible,
	// so in this senario explicit type cast is required int + int(int8)

	var n1 int = 10
	var n2 int8 = 11

	println(n1 + int(n2) // as n1 + n2 will through error (invalid operation: n1 + n2 (mismatched types int and int8))
}
