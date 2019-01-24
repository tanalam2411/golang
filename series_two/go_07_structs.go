package main

import (
	"fmt"
	"reflect"
)

/* AGENDA

- Structs
  - What are they?
  - Creating
  - Naming conventions
  - Embedding
  - Tags
*/

/*
Structs are value type
*/

// Doctor - D is capital so it is exposed outside package main
// but all inside attributes as starting with lower case which will not be exposed
type Doctor struct {
	number     int
	actorName  string
	companions []string
}

// Struct embedding

type Animal struct {
	Name   string
	Origin string
}

// Bird : has no relation with Animal struct its just embeds the feilds of it
type Bird struct {
	Animal
	SpeedKPH float32
	CanFly   bool
}

// Struct embedding
// Human - feild Name has tag which is just addiotnal info as metadta.
type Human struct {
	Name   string `required max:"100"`
	Origin string
}

func main() {

	aDoctor := Doctor{
		number:    3,
		actorName: "Max",
		companions: []string{
			"Lili",
			"Jhon",
			"Gopher",
		},
	}

	fmt.Printf("%v, %T \n", aDoctor, aDoctor)
	fmt.Printf("Doctor's name: %v \n", aDoctor.actorName)
	fmt.Printf("Doctor's one of the companions: %v \n", aDoctor.companions[1])

	// We can also create struct using positional syntax, but this is not recommended for many reasons:
	// Any change in position or update or deletition of struct attr will impact every place it is being used

	bDoctor := Doctor{
		5,
		"Yelo",
		[]string{"Melo", "Cilo"},
	}

	fmt.Printf("bDoctor: %v, %T \n", bDoctor, bDoctor)

	fmt.Println("****************** Anonymus struct *******************")

	cDoctor := struct {
		name   string
		number int
	}{name: "Test Doctor", number: 33}

	fmt.Printf("%v, %T \n", cDoctor, cDoctor)

	// Structs are value type, meaning if assigned to another variable a seperate copy is given not reference

	dDoctor := cDoctor
	dDoctor.name = "DTest Doctor"

	fmt.Printf("cDoctor: %v, %T \n", cDoctor, cDoctor)
	fmt.Printf("dDoctor: %v, %T \n", dDoctor, dDoctor)

	// We can point to the same underlying struct using & operator

	eDoctor := &cDoctor // pointer to the struct cDoctor
	eDoctor.name = "ETest Doctor"

	fmt.Printf("cDoctor: %v, %T \n", cDoctor, cDoctor)
	fmt.Printf("eDoctor: %v, %T \n", eDoctor, eDoctor)

	fmt.Println("************************** Struct embedding **************************")

	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 48.4
	b.CanFly = false

	fmt.Printf("Bird d: %v, %T \n", b, b)

	c := Bird{
		Animal:   Animal{Name: "Emu", Origin: "Australia"},
		SpeedKPH: 48.4,
		CanFly:   false,
	}
	fmt.Printf("Bird c: %v, %T \n", c, c)

	fmt.Println("************************** Struct tagging **************************")
	h := reflect.TypeOf(Human{})
	field, _ := h.FieldByName("Name")
	fmt.Println(field.Tag)
}

/* SUMMARY

- Struct
  - Collections of disparate data types that describe a single concept
  - keyed by named fields
  - Normally created as types, but anonymous structs are allowed
  - Struct are value types
  - No inheritance, but can use composition via embedding
  - Tags can be added to struct fields to describe field

*/
