

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
##### Nested interface

- Like a struct, an interface can also be nested in a struct means a field can have a `data type of a type interface`.
- Any data type that implements an interface can also be represented as a type of the interface(polymorphism)
- Using this `polimorphism` principle, we can have a struct field of an interface type.
  - Value of this field can be anything that implements that interface.
  
```go
type Salaried interface {
	getSalary() int
}


type Salary struct {
	basic		int
	insurance 	int
	allowance 	int
}


func(s Salary)getSalary() int {
	return s.basic + s.insurance + s.allowance
}

type Employee1 struct {
	firstName, lastName string
	salary				Salaried  // interface type
}

```  
 
- Since `Salary` struct implements `getSalary()` method, means it implements `Salaried` interface.
- Hence, we can store an instance of `Salary` struct type in a field of `Salaried` type. -> `salary				Salaried`
- **Note:** If no value is assigned to the `salary` field while creating an `Employee` struct, the Go will panic with a runtime error as we are trying to call a method on a `nil` value which is the default dynamic value of an interface.

```go
	ross := Employee1{
		firstName: "Ross",
		lastName:  "Geller",
	}

	fmt.Println("Ross's salary is: ", ross.salary.getSalary())
``` 

```bash
$ go run nested_interface.go 
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x48cf18]

goroutine 1 [running]:
main.main()
        /home/tan/tanveer/golang/src/golang/series_5/structs/1.Structures in Go/nested_interface.go:48 +0x78
exit status 2

```


- **`Anonymously nested interface`**: Similar to the field `promotions`, `methods` are also `promoted` when a struct field is an `anonymous interface`.  `ross.getSalary()`
```go
type Employee1 struct {
	firstName, lastName string
	Salaried
}

func main() {
	ross := Employee1{
		firstName: "Ross",
		lastName:  "Geller",
		Salaried:    Salary{1000, 50, 50},
		
	}

	fmt.Println("Ross's salary is: ", ross.getSalary())
}
```

- Here, we have removed the `salary` field from the `Employee` struct type and now `Salaried` is both the field name and field data type. 
- This will promote all the methods from `Salaried` interface on the parent struct `Employee` as if the `Employee` struct type implements those methods.
- **Note:** Similar to the field promotions of an anonymously nested struct, only the non-conflicting methods will get promoted.
  - Hence, if the `Employee` struct type also implements the `getSalary` method, then that will be used instead.

```go
func (e Employee1)getSalary() int {
	return 0
}

func main() {
	ross := Employee1{
		firstName: "Ross",
		lastName:  "Geller",
		//salary:    Salary{1000, 50, 50},
		Salaried:    Salary{1000, 50, 50},

	}

	//fmt.Println("Ross's salary is: ", ross.salary.getSalary()) // salary:    Salary{1000, 50, 50},
	fmt.Println("Ross's salary is: ", ross.getSalary())
	fmt.Println("Ross's salary is: ", ross.Salaried.getSalary())
}
```  

```bash
Ross's salary is:  0
Ross's salary is:  1100
```
  




























-----
[Ref](https://medium.com/rungo/structures-in-go-76377cc106a2)
