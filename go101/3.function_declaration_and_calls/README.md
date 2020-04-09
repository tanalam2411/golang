##### Function Declaration and Calls


- Go doesn't support default parameter values.
- The initial value of each result is the zero value of its type.
- Function can be declared before or after any of its calls.
    ```go
    func main() {
        v := VersionString()
    }
    
    func VersionString() string {
        return "v1.0"
    }
    ```
- As function types are incomparable types, they can't be used as the key types of map types.  
---

##### Variadic parameters and variadic function types

- The last parameter of a function can be variadic parameter.
- Each function can have at most one variadic parameter.
- The type of a variadic parameter is always a slice type.
```go
func (values ...int64) (sum int64)
func (sep string, tokens ...string) string
```
Builtin variadic functions:
```go
func Print(a ...interface{}) (n int, err error)
func Printf(format string, a ...interface{}) (n int, err error)
func Println(a ...interface{}) (n int, err error)
```
