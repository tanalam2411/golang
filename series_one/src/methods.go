package main

import "fmt"

/*
Two types of method:
i) Value receiver
ii) Pointer receiver

https://golang.org/doc/faq#methods_on_values_or_pointers
*/

const max_salary float64 = 1000000
const min_salary float32 = 10000

type person struct {
	first_name string
	last_name  string
	age        uint16
	salary     float64 //unsigned int - min 0 max 65535
	savings    float64
	address    string
}

func (p person) calTax() float64 {

	// Value receiver method
	// You can't update person's attr
	// p.salary = new_value // will not be updated

	if float64(p.salary) > max_salary {
		return .30*(max_salary) + .40*(float64(p.salary)-max_salary)
	} else {
		return .30 * float64(p.salary)
	}
}

func (p *person) add_bonus(bonus_amount float64) {

	// Pointer receiver method
	// it can updated person object or ...

	p.salary = bonus_amount + p.salary
}

func increment(p person, percentage float64) person {
	//function to act as both Value and Pointer receiver

	p.salary = float64(p.salary) + (percentage/100)*float64(p.salary)

	return p
}

func main() {

	person_1 := person{first_name: "Tanveer",
		last_name: "Alam",
		age:       28,
		salary:    30000,
		savings:   0.0001,
		address:   "India"}

	// person_1 := person{"Tanveer", "Alam", 28, 100, 0.1, "India"}

	fmt.Println(person_1.first_name, "details: ")
	fmt.Println(person_1)
	fmt.Println("Tax on salary without bonus amount:", person_1.calTax())

	person_1.add_bonus(5000)
	fmt.Println("Tax on salary with bonus amount:", person_1.calTax())
	fmt.Println("Salary before increment is: ", person_1.salary)

	person_1 = increment(person_1, 30)

	fmt.Println("Incremented salary is: ", person_1.salary)

}
