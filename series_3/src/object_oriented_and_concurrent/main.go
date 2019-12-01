package main

import (
	"bytes"
	"fmt"
	"io"
)

type Person struct { Name string }

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

func (w *WriteCounter) Write(b []byte) (int ,error) {
	w.count += len(b)
	return w.ReadWriter.Write(b)
}


func main() {
	var e Employee
	e.Name = "Person One"
	e.EmployeeID = 123

	fmt.Println( "Employee name is: ", e.Introduce())
	fmt.Println( "Employee ID is: ", e.GetID())

	buf := &bytes.Buffer{}
	w := &WriteCounter{ReadWriter: buf}

	fmt.Fprintf(w, "Hello, gophers!\n")
	fmt.Printf("Printed %v bytes", w.count)
}

