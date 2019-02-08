package main


import ()


/* Objectives

- Playing with values
	- Bytes
	- Integers
	- Strings
	- Booleans
	- Arrays
	- Maps
- Data type and Storage
*/

/* Values
- literal - 1, true, "text"
- computed - 1 + 4
*/

/* Everything is a Number
- Everything typed into a computer is a character
- Every character is represented by a number
- A number is build up by bit by bit, quite literally
*/

/* What is a bit ?
- A bit, is like a switch, it is either on or off
- Each bit represents a binary value of 0 or 1

1 bit - 2 states - 0 or 1
2 bits - 4 states - 00, 01, 10, 11

*/

/* Bit Grouping Illustrated

MSB  0 1 2 3 4 5 6 7  LSB

value - 1208
Big Endian (left-to-right): 1-2-0-8
Little Endian (right-to-left): 8-0-2-1
*/

 /* Grouping Bits

 - Grouping bits to represents numbers

 - 4 bits = nibble
 - 8 bits = 1 byte
 - 16 bits
 - 32 bits
 - 64 bits
 - 128 bits

 // 16, 32, 64, 128 - depends on programming language and sometimes on processor

 */

 /* Binary Number system

 - The `binary numbering system` uses bits tio represent numbers.
 - Since a certain number of bits allows us to represent a dsitinct number of states, we can choose to represent
   either positve or negative numbers.

 - To represent a negative number, we designate the MSB as the `signed bit`
 - The `sign bit` tells us if the number is positive or negative.

 - Signed and Unsigned Bits:

 Unsigned Bits	Unsigned Value Signed Bits Signed Value
 000			0				0-00		0
 001			1				0-01		1
 010			2				0-10		2
 011			3				0-11		3
 100			4				1-00		0
 101			5				1-01		-3
 110			6				1-10		-2
 111			7				1-11		-1
  */

  /* Bit Grouping and Numeric Range
  # of bits		# of unique values 		Unsigned Number Range 				Signed Number Range
  8				2^8 = 256				0 to 255							-127 to 127
  16			2^16 = 65,535			0 to 65,534							-32,767 to 32,767
  32			2^32 = 4,294,967,295	0 to 4,294,967,294					-2,147,483,647
  64			2^16 = 18,446,744,073	0 to 18,446,744,073,709,			-9.22337203685478E18 to
  						709,551,615			 551,614						9.22337203685478E18
  */

  /* ASCII Table

  https://www.rapidtables.com/code/text/ascii-table.html
   */