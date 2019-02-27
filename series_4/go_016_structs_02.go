package main

import "fmt"

type Person struct {
	id   string
	name string
	age  int
	snn  string
	Address
}

type Address struct {
	id     string
	stree1 string
	stree2 string
	city   string
	state  string
	Zip
}

type Zip struct {
	id string
	code int
	pobox int
}


func main() {
	var people []Person

	fmt.Printf("People: %#v \n", people)

	var p Person
	p = Person{"P1", "Max", 23, "000-111-222",
		Address{"A1", "Road-Block1", "23/4", "Pune", "MH",
		Zip{"Z1", 1212, 110022}},
	}

	people = append(people, p)

	//p = Person{"Yoyo", 33, "000-111-333", "Road-Block2",
	//	"23/7", "Pune", "MH", "412334"}
	//people = append(people, p)

	for _, p := range people {
		fmt.Printf("%v %v %v %v \n", p.id, p.name, p.age, p.snn)
		fmt.Printf("%v %v %v %v %v \n", p.Address.id, p.stree1, p.stree2, p.city, p.state)
		fmt.Printf("%v %v %v \n", p.Zip.id, p.code, p.pobox)
	}
}
