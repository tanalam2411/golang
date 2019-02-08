package main

import ()


 /*
 - Understanding Data Types

 - bool
 - string
 - Numeric types
 	- uint either 32 or 64 bits // depends on arch
 	- int  either 32 or 64 bits // depends on arch
 	- uintptr	an unsigned integer large enough to store the uninterpreted bits of a pointer value
 	- uint8, uint16, uint32, uint64
 	- int8, int16, int32, int64
 	- float32 	the set of all IEEE-754 32-bit floating-point number
 	- float62 	the set of all IEEE-754 64-bit floating-point number
 	- complex64		the set of all complex numbers with float32 real and imaginary parts
 	- complex128	the set of all complex numbers with float64 real and imaginary parts
	- byte 		alias for uint8
 	- rune 		alias for int32 (represents a Unicode code point)

 */

  /* Constant
 - Format Syntax - const name/identifier [type] = value
  - Examples:
  	- const pi float32 = 3.14
  	- const hoursInDay uint = 24
  	- const presenterName string = "Max"

  - Type Inferred Syntax: (Go always assigns the bigger data type - (int64, float 64)
  - Format - const name/identifier = value
  - Examples:
  	- const pi = 3.14
  	- const hoursInDay = 24
  	- const presenterName = "Max"
  */

/* Variable
- Format Syntax - var name/identifier [type] = value
 - Examples:
 	- var pi float32 = 3.14
 	- var hoursInDay uint = 24
 	- var presenterName string = "Max"

 - Type Inferred Syntax: (Go always assigns the bigger data type - (int64, float 64)
 - Format - var name/identifier = value
 - Examples:
 	- var pi = 3.14
 	- var hoursInDay = 24
 	- var presenterName = "Max"
*/