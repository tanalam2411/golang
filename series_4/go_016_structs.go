package main

import "fmt"

/* What is struct?

A struct is a sequence of named elements, called fields, each of which has a name and a type.
Field names may be specified explicitly(IdentifierList) or implicitly(AnonymousField).
Within a struct, non-blank field names must be unique.
Composite of multiple types
*/


func main() {
	fmt.Println("Introductions to Structs")

	var p struct{
		name string
		age int
		ssn string
	}

	fmt.Printf("p: %#v\n", p)

	p.name = "Hero"
	p.age = 23
	p.ssn = "1-2-3-4-5"

	fmt.Printf("%v\n", p)

	var p1 = struct{
		name string
		age int
		ssn string
	}{"Max", 21, "000-123-234"}
	/* or
	struct{
		name string
		age int
		ssn string
	}{
	age: 27,
	name: "Max", "000-123-234",
	ssn: "000-123-234"
	}

	 */

	fmt.Printf("p1: %#v\n", p1)


	var friends = []struct{
		name string
		age int
		ssn string
	}{{
		age: 12,
		name: "user_1",
		ssn: "222-33-4444",
	},
	   {
	   age: 13,
		   name: "user_2",
		   ssn: "33-22-5555",
	   },
	}

	fmt.Printf("friends: %#v\n", friends)

	friends = append(friends, struct {
		name string
		age  int
		ssn  string
	}{name: "Tan", age: 28, ssn: "111-222-3333"})
	fmt.Printf("friends: %v\n", friends)
}