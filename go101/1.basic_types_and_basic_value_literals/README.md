
##### Basic Types and Basic Value Literals

- Types cab be viewed as value templates, and values can be viewed as type instances.


- Build-in Basic Types in GO:

  - bool
  - int8, uint8, int16, uint16, int32, uint32, int64, uint64, int, uint, uintptr
  - float32, float64
  - complex64, complex124
  - string
  - byte - is a built-in alias of `uint8`
  - rune - is a built-in alias of `int32`
  
  - Note: `uintptr`'s size value must be large enough to store the uninterpreted bits of any memory address.
  - string: sequence of bytes having UTF-8 encoding.
  
---

- Zero Values
  - `false` : boolean
  - `0` : integers
  - `0.0` :  for floats
  - `""` : strings
  - `nil` :
	- pointers
	- functions
	- interfaces
	- slices
	- channels
	- maps