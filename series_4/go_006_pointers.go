package main

import "fmt"

/* Using Pointer

 - There are two parts to using a `pointer`
 - Reference and dereference
 	- reference: refers to the act of getting an address, the pointer
 	- We use the `address of` operator, &, to get the address of a variable
 		>> Eg: countAddr = &count
 - Dereference: means take this address/pointer, and return the value stored there.
 	- We use the `dereference` operator, *, to get the value of the indirection.
 		>> Eg.: copyOfCount = *countAddr

 */

 /* Declaring Pointer Variables

 - How do you declare a `variable` which holds a `pointer`?
 - Syntax:
 	var name/identifier [type][=pointer value]
 	- Where `type` is:
 		- <# indirections><type of value>
 	- Examples:
 		- var addressOfCount = &count
 		- var buffer *byte
 		- var pMessage = &message
 		- var ptrLoggingEnabled = &isLoggingEnabled

 Note: A pointer variable that is not given a value, gets the value `nil`, which is like 0.

  */

 func main() {

 	var i int8 = 30

 	iAddr := &i

 	iiAddr := &iAddr

 	fmt.Println(i, iAddr, iiAddr, **iiAddr)
 	fmt.Printf("%T, %T, %T \n\n",i, iAddr, iiAddr )


 	var buffer *byte

 	fmt.Println(buffer)

 	//const con = 3.14
	 //
 	//conPntr := &con  // we cannot take address of constants and nil
	 //
 	//fmt.Println(conPntr)

 }

/*
30 0xc00001e0bf 0xc00000c028 30
int8, *int8, **int8

<nil>
*/

/* Try dereferencing pointer itself - fmt.Println(*buffer)
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x486a02]

goroutine 1 [running]:
main.main()
 */