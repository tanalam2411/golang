package main

import "fmt"

type Employee struct {
	// firstName, lastName	string		// is also a valid syntax
	firstName	string
	lastName	string
	salary		int
	fullTime	bool
}


func (e *Employee)Print() {
	fmt.Printf("\n----------------- Employee: %s %s ----------------\n", e.firstName, e.lastName)

	fmt.Println("FirstName: ", e.firstName)
	fmt.Println("LastName: ", e.lastName)
	fmt.Println("Salary: ", e.salary)
	fmt.Println("Fulltime: ", e.fullTime)
}

func (e *Employee)newEmployee(firstName, lastName string, salary int, fullTime bool) {
	e.firstName = firstName
	e.lastName = lastName
	e.salary = salary
	e.fullTime = fullTime
}

func main() {
	var ross Employee   // Just defines variable ross but doesn't initializes it.
	fmt.Println(ross)
	ross.Print()

	ross.newEmployee("Ross", "Tailor", 5000, true)
	ross.Print()

	james := Employee{
		firstName: "James",
		salary:    10000,
		lastName:  "Bing",
	}
	// james := Employee{"James", "Bing", 10000, false} // Values should be provided in order when field names are not provided
	james.Print()
}

