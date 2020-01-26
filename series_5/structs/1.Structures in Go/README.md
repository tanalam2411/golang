

##### Structures in Go (structs)

- Struct or Structure is mainly used to define a `schema` with different fields(properties), like a `class`.
- When is say it is a struct means we're calling an `instance` of type struct.
- When we say `type struct` meaning we're representing schema of a type.

Syntax to create a struct type:

```go
type StructNAme struct {
    field1  fieldType1
    field2  fieldType2 
}
```

-----
[Ref example](example1.go)

##### Let's create a struct type `Employee`

```go

type Employee struct {
	// firstName, lastName	string		// is also a valid syntax
	firstName	string
	lastName	string
	salary		int
	fullTime	bool
}
```

##### Creating a struct(instance)

```go
func main() {
	var ross Employee   // Just defines variable ross but doesn't initializes it. 
	fmt.Println(ross)
}
```

output:
```bash
$ go run example1.go 
{  0 false}
```
It prints default values of the struct(Employee) fields, i.e `'' '' 0 false`

Here `Employee` is a struct type(class) and `ross` is a struct(instance or an object).

-----
##### Getting and setting struct fields

```go
	james := Employee{
		firstName: "James",
		salary:    10000,
		lastName:  "Bing",
	}
	// james := Employee{"James", "Bing", 10000, false} // Values should be provided in order when field names are not provided

```
```bash
----------------- Employee: James Bing ----------------
FirstName:  James
LastName:  Bing
Salary:  10000
Fulltime:  false
```

- The order of the appearance of struct's fields does not matter.
- `,` Comma is necessary after value value assignment of the last field. 
  - This way `go` won't add a `semicolon (;)` just after the last field while compiling the code  
- Initialization with not all field`s value will assign it's default value.


-----

##### Anonymous struct or Inline struct
- Anonymous struct is a struct with no explicitly defined struct type alias.
































-----
[Ref](https://medium.com/rungo/structures-in-go-76377cc1)