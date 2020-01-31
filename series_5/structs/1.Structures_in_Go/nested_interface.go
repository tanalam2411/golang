package main

import "fmt"


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
	//salary				Salaried  // interface type
	Salaried

}


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
