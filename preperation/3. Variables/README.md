##### Variables

- A variable is a storage location for holding a value.
- The set of permissible values is determined by the variable's `type`.
- A type may be denoted by a `type` name, or specified using `type literal`, which compose a type from existing types.
- ```go
  var s1 string // from type literal string
  var s2 s1 // by type name `s1`  
  ```
- Predeclared types:
  - Types: bool, byte, complex64, complex128, error, float32, float64, 
           int, int8, int 16, int32, int64, rune, string, 
           uint, uint8, uint16, uint32, uint64, uintptr
  
  - Constants: true, false, iota           
  
  - Zero value: nil
  
  - Functions: append, cap, close, complex, copy, delete, imag, len,
               make, new, panic, print, println, real, recover
               
- Composite types - array, struct, function, interface, slice, map, and channel types-may be constructed using type literals.                

- Each type T hs an underlying type: If T is one of the predeclared type, then the corresponding underlying type is T itself.
  - Otherwise, T's underlying type is the underlying type of the type which T refers in its type declaration.
  -
    ```go
    type (
        A1 = string
        A2 = A1
    )
    
    type (
        B1 string
        B2 B1
        B3 []B1
        B4 B3
    )
    ```
    
    - The underlying type of A1, A2, B1, and B2 is string.
    - The underlying type of []B1, B3, and B4 is []B1.
    
- Variables of interface type also have a distinct dynamic type, which is the concrete type of the value assigned to the variable at run time(unless the value is the predeclared identifier `nil`, which has no type).
  - The dynamic type may vary during execution but values stored in interface variables are always assignable to the static type of the variable.
  - ```go
    package main
    
    import "fmt"
    
    func main() {
        var x interface{}
        var v *int
        fmt.Printf("x: %v, %T\n", x, x)
        x = 42
        fmt.Printf("x: %v, %T\n", x, x)
        x = "string"
        fmt.Printf("x: %v, %T\n", x, x)
        x = v
        fmt.Printf("x: %v, %T\n", x, x)
        x = nil
        fmt.Printf("x: %v, %T\n", x, x)
    }
    ```
    ```bash
    x: <nil>, <nil>
    x: 42, int
    x: string, string
    x: <nil>, *int
    x: <nil>, <nil>
    ```
    ```go
    var x interface{}  // x is nil and has static type interface{}
    var v *T           // v has value nil, static type *T
    x = 42             // x has value 42 and dynamic type int
    x = v              // x has value (*T)(nil) and dynamic type *T
    ```
    - A variable's value is retrieved by referring to the vairable in an expression, it is the most resent value assigned ot th variable.
    - If a variable has not been assigned a value, its value is the zero value of its type.
    - ```go
        package main
        
        import (
            "fmt"
            "reflect"
        )
        
        type S struct {
            Name string
        }
        
        func main() {
            var b bool
            var i int
            var f float32
            var s string
            var arr [10]int
            var sli []int
            var mp map[string]string
            var ifc interface{}
            var s1 S
        
            fmt.Printf("boolean: %v, %T, %v\n", b, b, reflect.ValueOf(b).IsZero())
            fmt.Printf("integer: %v, %T, %v\n", i, i, reflect.ValueOf(i).IsZero())
            fmt.Printf("float: %v, %T, %v\n", f, f, reflect.ValueOf(f).IsZero())
            fmt.Printf("string: %v, %T, %v\n", s, s, reflect.ValueOf(s).IsZero())
            fmt.Printf("array: %v, %T, %v\n", arr, arr, reflect.ValueOf(arr).IsZero())
            fmt.Printf("slice: %v, %T, %v\n", sli, sli, sli == nil)
            fmt.Printf("map: %v, %T, %v\n", mp, mp, mp == nil)
            fmt.Printf("interface: %v, %T, %v\n", ifc, ifc, ifc == nil)
            fmt.Printf("struct var: %v, %T, %v\n", s1, s1, reflect.ValueOf(s1).IsZero())
            snew := new(S)
            fmt.Printf("new struct: %v, %T, %v\n", snew, snew, reflect.ValueOf(snew).IsZero())
        }
      ```
      ```bash
        boolean: false, bool, true
        integer: 0, int, true
        float: 0, float32, true
        string: , string, true
        array: [0 0 0 0 0 0 0 0 0 0], [10]int, true
        slice: [], []int, true
        map: map[], map[string]string, true
        interface: <nil>, <nil>, true
        struct var: {}, main.S, true
        new struct: &{}, *main.S, false
      ```
      
      
##### Type Declarations:
- Type declaration binds an identifier, the type name, to a type.
- Type declarations cme in two forms:
  - Alias declarations:
    ```go
    type (
        nodeList = []*Node // nodeList and []*Node are identical types
        Polar    = polar   // Polar and polar denote identical types
    )
    ```
  - Type definitions: creates new, distinct type with same underlying type.
    - The new type is called defined type.
    - It is different from any other type, including the type it is created from
        ```go
        package main
        
        import "fmt"
        
        type Node struct {
            child string
        }
        
        func (n *Node) GetChild() string {
            return n.child
        }
        
        type (
            N  = Node
            N1 Node
        )
        
        func main() {
            var n N
            var n1 N1
        
            fmt.Printf("n: %v, %T\n", n, n)
            fmt.Printf("n1: %v, %T\n", n1, n1)
        
            n = Node{child: "n child"}
            //n1 = Node{child: "n1 child"} // cannot use Node literal (type Node) as type N1 in assignment
            n1 = N1{child: "n1 child"}
            fmt.Printf("n: %v, %T\n", n, n)
            fmt.Printf("n1: %v, %T\n", n1, n1)
        
            n.GetChild()
            //n1.GetChild() // n1.GetChild undefined (type N1 has no field or method GetChild)
        
        }
        ```    
        ```bash
        n: {}, main.Node
        n1: {}, main.N1
        n: {n child}, main.Node
        n1: {n1 child}, main.N1
        
        ```
    - A defined type may have methods associated with it. It does not inherit any methods bound to the given type.
    - But method set of an interface type or of elements a composite type remain unchanged.
      ```go
        package main
        
        import "fmt"
        
        type I interface {
            PrintName()
            PrintAge()
        }
        
        type I1 I
        
        type S1 struct {
            Name string
        }
        
        func (s1 *S1) PrintName() {
            fmt.Printf("Name: %v\n", s1.Name)
        }
        
        type S2 struct {
            S1
            Age int
        }
        
        func (s2 *S2) PrintAge() {
            fmt.Printf("Age: %d\n", s2.Age)
        }
        
        func main() {
            s2 := &S2{
                S1:  S1{Name: "Max"},
                Age: 10,
            }
        
            s2.PrintName()
            s2.PrintAge()
        
            fmt.Printf("s2 details, Name: %v, Age: %v\n", s2.Name, s2.Age)
            fmt.Printf("s2 details, Name: %v, Age: %v\n", (s2).Name, (*s2).Age)
        
            var i interface{} = s2
            val, ok := i.(I)
            fmt.Printf("Does S2 implements interface I: %v, %v\n", val, ok)
        
            val1, ok1 := i.(I1)
            fmt.Printf("Does S2 implements interface I1: %v, %v\n", val1, ok1)
        }
      ``` 
      ```bash
        Name: Max
        Age: 10
        s2 details, Name: Max, Age: 10
        s2 details, Name: Max, Age: 10
        Does S2 implements interface I: &{{Max} 10}, true
        Does S2 implements interface I1: &{{Max} 10}, true
      ```
      From doc
      ```go
        // A Mutex is a data type with two methods, Lock and Unlock.
        type Mutex struct         { /* Mutex fields */ }
        func (m *Mutex) Lock()    { /* Lock implementation */ }
        func (m *Mutex) Unlock()  { /* Unlock implementation */ }
        
        // NewMutex has the same composition as Mutex but its method set is empty.
        type NewMutex Mutex
        
        // The method set of PtrMutex's underlying type *Mutex remains unchanged,
        // but the method set of PtrMutex is empty.
        type PtrMutex *Mutex
        
        // The method set of *PrintableMutex contains the methods
        // Lock and Unlock bound to its embedded field Mutex.
        type PrintableMutex struct {
            Mutex
        }
        
        type Block interface {
            BlockSize() int
            Encrypt(src, dst []byte)
            Decrypt(src, dst []byte)
        }
        // MyBlock is an interface type that has the same method set as Block.
        type MyBlock Block
      ```