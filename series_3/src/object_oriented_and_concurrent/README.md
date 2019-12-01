## Go: Object Oriented and Concurrent 

Go is:

1. Statically types - type on object never changes
2. Part of the C family
3. Garbage collected
4. Statically compiled
5. Object oriented
6. Concurrency friendly


---
#### Object orientation in Go

##### Methods:

```go
type Door struct{
    opened bool
}

func (d *Door) Open() {
    d.opened = true
}

func (d *Door) Close() {
    d.opened = false
}
``` 

Note :- We pass pointer(*) of Door struct to modify the referenced object other wise it will just modify the local copy of that struct.


Creating methods on other type rather then `struct`:

```go
type Door bool

func (d *Door) Open() {
    *d = true
}

func (d *Door) Close() {
    *d = flase
}
```

But only on the type within the same package.

No monkey patching is allowed

So this is not allowed: ```go func (s string`)Length() { return 2*len(s) }```

---
#### Polymorphism

- Objects of multiple forms can be accessed through the same `interface`.
- Each type can provide its own, independent implementation of this `interface`.

```go
type OpenCloser interface {
    Open()
    Close()
}
```

- Interfaces with one of two methods can be named as <method_name>+`er`.
- Go interfaces are satisfied `implicitly`. 
- If a type defines all the methods of an interface, the type satisfies that interface.
- So if two interfaces have different name but have same methods, then they're same interface.

---
#### Inheritance vs Composition

- There's no inheritance in Go.

##### Composition

```go
type Runner struct { name string }

func (r *Runner) Name() string { return r.name }

func (r *Runner) Run(t Task) {
    t.Run()
}

func (r *Runner) RunAll(ts []Task) {
    for _, t := range ts {
        r.Run(t)
    }
}


type RunCounter struct {
    runner Runner
    count int 
}

func NewRunCounter(name string) *RunCounter {
    return &RunCounter{runner: Runner{name}}
}

func (r *RunCounter) Run(t task) {
    r.count ++
    r.runner.Ruun(t)
}

func (r *RunCounter) RunAll(ts []Task) {
    r.count += len(ts)
    r.runner.RunAll(ts)
}

func (r *RunCounter) Count() int { return r.count }

func (r *RunCounter) Name() string { return r.runner.Name()}
```            

Here we're repeating `Name()` twice, which can be solved using `struct embedding`

---
#### Struct embedding

- Expressed in Go as unnamed fields in a struct.
- The fields and methods of embedded type are defined on the embedding type.
- Similar to inheritance, but the embedded type doesn't know it's embedded.

Example of `struct embedding`:

```go
package main

import "fmt"

type Person struct { Name string }

func (p Person) Introduce() string { return p.Name }

type Employee struct {
	Person
	EmployeeID int
}

func (e Employee) GetID() int { return e.EmployeeID }

// All fields and methods from Person are available on Employee.


func main() {
	var e Employee
	e.Name = "Person One"
	e.EmployeeID = 123

	fmt.Println( "Employee name is: ", e.Introduce())
	fmt.Println( "Employee ID is: ", e.GetID())

}

```

Updating our runner code:

```go
type RunCounter2 struct {
    Runner
    count int
}

func NewRunCounter2(name string) *RunCounter2 {
    return &RunCounter2 {Runner{name}, 0}
}

func (r *RunCounter2) Run(t Task) {
    r.count++
    r.Runner.Run(t)
}

func (r *RunCounter2) RunAll(ts []Task) {
    r.count += len(ts)
    r.Runner.RunAll(ts)
}

func (r *RunCounter2) Count() int {
    return r.count
}
```

Diamond Problem - If a struct has two other structs embedded having one ore more same method, in that case call that method with its parent struct name.
  - Example if `RunCounter2` has two structs `Runner` and `Runner2` in that case you can call `Runner.Run` and `Runner2.Run` which wil explicit call with no ambiguous. 
  
So, `struct embedding` is a composition.
  - You can't reach into another type and change the way it works.
  - Method dispatch is explicit.
  
---

##### Struct embedding of interfaces

```go
// WriteCounter tracks the total number of bytes written
type WriteCounter struct {
    io.ReadWritter
    count int
}

func (w *WriteCounter) Write(b []byte) (int, error) {
    w.count += len(b)
    return w.ReadWriter.Write(b)
}

// WriteCounter can be used with any io.ReadWriter

func main() {
    buf := &bytes.Buffer{}
    w := &WriteCounter{ReadWriter: buf}
    
    fmt.Fprintf(w, "Hello, gophers!\n")
    fmt.Printf("Printed %v bytes", w.count)
}
```

---

#### Easy mocking

Faking `net.Conn`

```go

type Conn interface {
    Read(b []byte) (n int, err error)
    Write(b []byte) (n int, err error)
    Close() error
    LocalAddr() Addr
    RemoteAddr() Addr
    SetDeadline(t time.Time) error
    SetReadDeadline(t time.Time) error
    SetWriteDeadline(t time.Time) error
}
```