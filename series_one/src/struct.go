package main

import "fmt"

type person struct {
	first_name string
	last_name  string
	age        int16
	salary     uint16 //unsigned int - min 0 max 65535
	savings    float64
	address    string
}

func main() {

	person_1 := person{first_name: "Tanveer",
		last_name: "Alam",
		age:       28,
		salary:    0000,
		savings:   0.0001,
		address:   "India"}

	// person_1 := person{"Tanveer", "Alam", 28, 100, 0.1, "India"}

	fmt.Println(person_1.first_name, "details: ")
	fmt.Println(person_1)
}
