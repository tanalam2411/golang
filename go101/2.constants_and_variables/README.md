##### Constants and Variables 


- Constants in Go may be typed or untyped. [ref](https://riptutorial.com/go/example/12431/typed-vs--untyped-constants)
  - Untyped:
    ```go
    const foo = "bar"
    ```
    ```go
    const Pi = 3.1416
    
    // Declare multiple constants in a group
    const (
    	No         = !Yes
    	Yes        = true
    	MaxDegrees = 360
    	Unit       = "radian"
    )
    
    // Declare multiple constants in one line
    const One, Two = 1, 2
    ```  
  - Typed:
    ```go
    const typedFoo string = "bar"
    
    const X = float32(3.14)
    
    const(
        A, B int64   = -3, 5
        Y    float32 = 2.718
    )
    ```

  
  - The constants declared in function bodies are called local constants.
  - The variables declared out of any function body are called package-level constants or global constants.    
  - Constants or Variables defined outside function and if starts with Capital letter than it's scope is across packages.
  - Constants are immutable.
---

- Iota - represents successive untyped integer constants.
  ```go
  const (
      x = iota   // by default starts with 0, you can set other number 
                 // like x = iota + 100, then y would be 101 and z would be 102
      y          // then y value will set to 1
      z          // then y value will set to 2
  )  
  ```
  
---

##### Variables, Variables Declarations and Value Assignments

- All variables are typed values.
  ```go
    var lang string = "Go"
    var x int  // will set default null value of int i.e., 0
    var num, status = 100, false // num as int, status as bool
    
    var (
      a = 1
      b = "abc"
      c float32 = 1.4
      d string
    )
  
    var aa, bb int // is invalid
  ```
  
---

#####  Short variable declaration forms

- Short variable declarations can only be used to declare local variables.
  ```go
  func main() {
      lang := "Go"
  }
  ```   

- local variables declared but used is not allowd.
- Package level variables have no such limit.

---

##### Value Addressability

- All variables are addressable.

  