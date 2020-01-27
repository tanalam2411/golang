

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

```go
	monica := struct{
		firstName, lastName string
		salary 				int
		fullTime			bool
	}{
		firstName: "Monica",
		lastName:  "Geller",
		salary:		1200,
	}

        fmt.Printf("%T\n", monica)
```

-----

##### Pointer to struct

- Creating `pointer` that points to the value of a struct.
```go
	ross := &Employee{
		firstName: "Ross",
		lastName:  "Bing",
		salary:    7000,
		fullTime:  true,
	}

	fmt.Println("firstName: ", (*ross).firstName)
	fmt.Println("firstName: ", ross.firstName)
```
- `*ross` - is a `dereferencing syntax` to get the actual value of the struct it is pointing to.
- `(*ross)` - parenthesis around the pointer is used so that compiler doesn't get confused between `(*ross).firstName` vs `*(ross.firstName)` -> when using `*ross.firstName` syntax.
- `Go` provides dereferncing by default, so we don't need to explicitly using `*` to derefer a pointer. Just `ross.firstName` would work.

-----
##### Anonymous fields

- One can define a struct type without declaring any field names. 
```go
type Data struct {
	string
	int
	bool
}

type Employee struct {
    firstName, lastName string
    salary              int
    bool                // anonymous field
}
```
- Go will use the data type declarations(keywords) as the field names.

-----

##### Nested struct

- A struct field can jhave a data type that is a `struct` type, called nested struct.
```go
type Department struct {
	name	string
	code	string
}

type Employee struct {
	// firstName, lastName	string		// is also a valid syntax
	firstName	string
	lastName	string
	salary		int
	fullTime	bool
	department	Department
}
```        

-----

##### Promoted fields

- Nesting struct without providing field name. `InsuranceDetails` 

```go
type InsuranceDetails struct {
	policyName 		string
	totalMembers  	int
}

type Employee struct {
	// firstName, lastName	string		// is also a valid syntax
	firstName	string
	lastName	string
	salary		int
	fullTime	bool
	department	Department
	InsuranceDetails
}
```   

- All fields of Anonymous nested struct `InsuranceDetails` (not Department) gets promoted to parent struct `Employee` and fields can be accessed as it they were defined on the `Employee` struct.
- will work
    ```go
    e.totalMembers
    ```
- will not work
    ```go
    e.code // e.department.code // will work
    ```  
  
- Notes: If a nested anonymous struct has a same field(field name) that conflicts with field name defined in the parent struct, then that field won't get promoted. Only the `non-conflicting` field will get promoted.  


  




























-----
[Ref](https://medium.com/rungo/structures-in-go-76377cc1)