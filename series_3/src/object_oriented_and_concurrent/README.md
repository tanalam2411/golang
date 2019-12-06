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

We want to test handleCon:
```go
func handleConn(conn net.Conn)
```

- We could create a fakeConn and define all the methods of Conn on it.

- we can test handleCon with loopBack type:
```go
type loopBack struct {
    net.Conn
    buf bytes.Buffer    
}
```

Any calls to the methods of `net.Conn` will fail, since the field is `nil`.

We redefine the operations we support:

```go
func (c *loopBack) Read(b []byte) (int, error) {
    return c.buf.Read(b)
}
```

```go
func (c *loopBack) Write(b []byte) (int, error) {
    return c.buf.Write(b)
}
```

---
- [Book on Concurrency](http://www.r-5.org/files/books/computers/languages/sql/nosql/Paul_Butcher-Seven_Concurrency_Models_in_Seven_Weeks-EN.pdf)
- [Callback](http://callbackhell.com/)
---

#### Go Concurrency
- Is is part of the language, not a library.
- Based on two concepts:
  - goroutines: lightweight threads
  - channels: typed pipes used to communicate and synchronize between goroutines. 

Communicating through channels:
`sleepAndTalk` sends the string into the channel instead of printing it.
```go
func sleepAndTalk(secs time.Duration, msg string, c chan string) {
    time.Sleep(sec * time.Second)
    c <- msg
}
```

We create channel and pass it to `sleepAndTalk`, then wait for the values to be sent.
- channels are pass by reference.
```go
func main() {
    c := make(chan string)
    
    go sleepAndTalk(0, "Hello", c)
    go sleepAndTalk(1, "Gophers!", c)
    go sleepAndTalk(2, "What's", c)
    go sleepAndTalk(3, "up?", c)

    for i:=0; i < 4; i++ {
        fmt.Printf("%v", <-c)
    }
}
```

`fmt.Printf("%v", <-c)`:
  - When the `channel` is empty and if we try read, its will be blocking `<-c`.
  - Same way if something is send over channel and there is no receiver on the other end, it will block. look buffered channel where you can define size of channel `c := make(chan string, 4)`.
  - So by default channels are synchronous.
  

---

Let's count on the web

We receive the next id from a channel:

```go
var nextID = make(chan int)

func handler(w http.ResponseWritter, r *http.Request) {
    fmt.Fprintf(w, "<h1>You got: %v</h1>", <-nextID)
}

// We need a goroutine sending ids into the channel

func main() {
    http.HandleFunc("/next", handler)
    
    go func() {
        for i := 0; ; i++ {
            nextID <- i
        }   
    }()
    
    http.ListenAndServer("localhost:8000", nil)
}
```

http://localhost:8000/next


Here `nextID <- i` will be blocked until next http call is triggered.

---

Let's fight

`select` allows us to chose among multiple channel operations:

```go
var battle = make(chan string)

func fightHandler(w http.ResponseWriter, r *http.Request) {
	select {
		case battle <- r.FormValue("usr"):
			fmt.Fprintf(w, "You won")
		case won := <-battle:
			fmt.Fprintf(w, "You lost, %v is better than you.", won)
	}
}

```
Go - http://localhost:8000/fight?usr=go

Java - http://localhost:8000/fight?usr=java

--- 
##### Chain of gophers

```go

package main

import (
	"fmt"
	"time"
)

func f(left, right chan int) {
	left <- 1 + <-right
}

func main() {
	start := time.Now()
	const n = 100000
	leftmost := make(chan int)

	right := leftmost
	left := leftmost

	for i :=0; i < n; i++ {
		right = make(chan int)
		go f(left, right)
		left = right
	}

	go func(c chan int) { c <- 0 }(right)

	fmt.Println(<-leftmost, time.Since(start))
}

```


---
ref - https://www.youtube.com/watch?v=Ng8m5VXsn8Q