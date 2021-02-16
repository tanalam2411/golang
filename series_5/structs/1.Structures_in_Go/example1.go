package main

import "fmt"

type InsuranceDetails struct {
	policyName   string
	totalMembers int
}

type Department struct {
	name string
	code string
}

type Employee struct {
	// firstName, lastName	string		// is also a valid syntax
	firstName  string
	lastName   string
	salary     int
	fullTime   bool
	department Department
	InsuranceDetails
}

func (e *Employee) Print() {
	fmt.Printf("\n----------------- Employee: %s %s ----------------\n", e.firstName, e.lastName)

	fmt.Println("FirstName: ", e.firstName)
	fmt.Println("LastName: ", e.lastName)
	fmt.Println("Salary: ", e.salary)
	fmt.Println("Fulltime: ", e.fullTime)
	fmt.Println("Department name: ", e.department.name)
	fmt.Println("Department code: ", e.department.code) // e.code won't work
	fmt.Println("PolicyName: ", e.InsuranceDetails.policyName)
	fmt.Println("TotalMembers: ", e.totalMembers)
}

func (d *Department) Print() {

}

func (e *Employee) newEmployee(firstName, lastName string, salary int, fullTime bool) {
	e.firstName = firstName
	e.lastName = lastName
	e.salary = salary
	e.fullTime = fullTime
}

func main() {
	var ross Employee // Just defines variable ross but doesn't initializes it.
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

	anonymousStruct()

	pointerToStruct()

	anonymousFields()

	nestedStruct()

	promotedFields()
}

func anonymousStruct() {
	fmt.Println("----------- anonymous struct ----------------")
	monica := struct {
		firstName, lastName string
		salary              int
		fullTime            bool
	}{
		firstName: "Monica",
		lastName:  "Geller",
		salary:    1200,
	}

	fmt.Println(monica)
	fmt.Printf("%T\n", monica)
}

func pointerToStruct() {
	fmt.Println("----------- pointer to struct ----------------")
	ross := &Employee{
		firstName: "Ross",
		lastName:  "Bing",
		salary:    7000,
		fullTime:  true,
	}

	fmt.Println("firstName: ", (*ross).firstName)
	fmt.Println("firstName: ", ross.firstName)

	fmt.Printf("Type of ross %T\n", ross)
	fmt.Printf("Type of ross %T\n", *ross)
}

type Data struct {
	string
	int
	bool
}

func anonymousFields() {
	fmt.Println("----------- anonymous fields in struct----------------")
	d := Data{"ABC", 1000, true}
	fmt.Println(d.string, d.int, d.bool)
}

func nestedStruct() {
	fmt.Println("----------- nested struct----------------")
	e := Employee{
		firstName:  "Ross",
		lastName:   "Gailer",
		salary:     11000,
		fullTime:   false,
		department: Department{"IT", "IT-001"},
	}

	e.Print()

}

func promotedFields() {
	fmt.Println("----------- promoted fields nested struct----------------")
	e := Employee{
		firstName:        "Ross",
		lastName:         "Geller",
		salary:           1000,
		fullTime:         false,
		department:       Department{"Sales", "Sales-002"},
		InsuranceDetails: InsuranceDetails{"Health", 3},
	}

	e.Print()
	fmt.Println(" Ross: ", e)
}

/*
{  0 false { } { 0}}

----------------- Employee:   ----------------
FirstName:
LastName:
Salary:  0
Fulltime:  false
Department name:
Department code:
PolicyName:
TotalMembers:  0

----------------- Employee: Ross Tailor ----------------
FirstName:  Ross
LastName:  Tailor
Salary:  5000
Fulltime:  true
Department name:
Department code:
PolicyName:
TotalMembers:  0

----------------- Employee: James Bing ----------------
FirstName:  James
LastName:  Bing
Salary:  10000
Fulltime:  false
Department name:
Department code:
PolicyName:
TotalMembers:  0
----------- anonymous struct ----------------
{Monica Geller 1200 false}
struct { firstName string; lastName string; salary int; fullTime bool }
----------- pointer to struct ----------------
firstName:  Ross
firstName:  Ross
Type of ross *main.Employee
Type of ross main.Employee
----------- anonymous fields in struct----------------
ABC 1000 true
----------- nested struct----------------

----------------- Employee: Ross Gailer ----------------
FirstName:  Ross
LastName:  Gailer
Salary:  11000
Fulltime:  false
Department name:  IT
Department code:  IT-001
PolicyName:
TotalMembers:  0
----------- promoted fields nested struct----------------

----------------- Employee: Ross Geller ----------------
FirstName:  Ross
LastName:  Geller
Salary:  1000
Fulltime:  false
Department name:  Sales
Department code:  Sales-002
PolicyName:  Health
TotalMembers:  3
 Ross:  {Ross Geller 1000 false {Sales Sales-002} {Health 3}}

*/
