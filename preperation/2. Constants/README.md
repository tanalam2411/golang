##### Constants

- Unchanging value

- Different types of Constants:
  - Boolean constants
  - Rune Constants
  - Integer Constants
  - Floating-point Constants
  - Complex Constants
  - String Constants
  
  

- String constants: 
  - ```go
    const hello = "Hello" // or `Hello`
    ```
  - What type does this string constant have?
    - Its an `untyped string constant`
    - Untyped constants are those constants whose type is not explicitly declared.
    - An Untyped constant has a default type which is the type to which the constant is implicitly converted in contexts where a typed value is required. 
    - The default type of an untyped constant is `bool, rune, int, float64, complex128 or string`
  
  - Types constant:
    - ```go
      const hello string = "hello"  
      ```  
      
- Typed vs Untyped constant
  - ```go
    package main
    
    import "fmt"
    
    const untypedConst = "untyped Hello" // untyped constant
    const typedConst string = "typed hello" // typed constant
    
    func main() {
        fmt.Printf("untypedConst-> value: %v, type: %T\n", untypedConst, untypedConst)
        fmt.Printf("typedConst-> value: %v, type: %T\n", typedConst, typedConst)
    
        type MyString string
        var ms MyString
    
        // Unlike the typed constants, the untyped constants have no type.
        // Assigning them to a variable of any type compatible with strings works without error
        ms = untypedConst
        fmt.Printf("ms: value: %v, type: %T\n", ms, ms)
    
        // cannot use typedHello (type string) as type MyString in assignment
        ms = typedConst
        fmt.Printf("ms: value: %v, type: %T\n", ms, ms)
    
    }
    ```