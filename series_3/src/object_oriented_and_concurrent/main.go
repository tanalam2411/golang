package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"sort"
)

type Person struct{ Name string }

func (p Person) Introduce() string { return p.Name }

type Employee struct {
	Person
	EmployeeID int
}

// All fields and methods from Person are available on Employee.
func (e Employee) GetID() int { return e.EmployeeID }

// WriteCounter tracks the total number of bytes written.
type WriteCounter struct {
	io.ReadWriter
	count int
}

func (w *WriteCounter) Write(b []byte) (int, error) {
	w.count += len(b)
	return w.ReadWriter.Write(b)
}

var nextId = make(chan int)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>You got %v</h1>", <-nextId)
}

var battle = make(chan string)

func fightHandler(w http.ResponseWriter, r *http.Request) {
	select {
	case battle <- r.FormValue("usr"):
		fmt.Fprintf(w, "You won")
	case won := <-battle:
		fmt.Fprintf(w, "You lost, %v is better than you.", won)
	}
}

func main1() {
	var e Employee
	e.Name = "Person One"
	e.EmployeeID = 123

	fmt.Println("Employee name is: ", e.Introduce())
	fmt.Println("Employee ID is: ", e.GetID())

	buf := &bytes.Buffer{}
	w := &WriteCounter{ReadWriter: buf}

	fmt.Fprintf(w, "Hello, gophers!\n")
	fmt.Printf("Printed %v bytes", w.count)

	// func literals and lambdas
	people := []string{"Alice", "Bob", "Dave"}
	less := func(i, j int) bool {
		return len(people[i]) < len(people[j])
	}
	sort.Slice(people, less)
	sort.Ints([]int{1, 2, 3})
	fmt.Println(people)

	people = nil
	fmt.Printf("%v:%T:%v", people, people, people)

	// ---

	http.HandleFunc("/next", handler)
	http.HandleFunc("/fight", fightHandler)
	go func() {
		for i := 0; ; i++ {
			nextId <- i
		}
	}()

	fmt.Println("serving at: http://localhost:8000/next")
	http.ListenAndServe("localhost:8000", nil)

}
