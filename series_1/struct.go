package main

import "fmt"

type person struct {
	first_name string
	last_name  string
	age        int16  // signed int   = -32k - +32k
	salary     uint16 // unsigned int = min 0 max 65535
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

	fmt.Println(person_1.first_name, "details: ")
	fmt.Println(person_1)
}
