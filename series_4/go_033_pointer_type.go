package main

import "fmt"

/* Objective

- Understanding Method Sets
 - Method Set of an interface
 - Method Set of a type
   - T vs *T
*/

/*
- A type may have a method set associated with it. The method set of an interface type is its interface.
  The method set of any other type T consists of all methods declared with receiver type T.
  The method set of the corresponding pointer type *T is the set of all methods declared with
  receiver *T or T(that is, it also contains the method set of T).
  Further rules apply to structs containing anonymous fields.
- Any other type has an empty method set.
  In a method set, each method must have a unique non-blank method name.
- The method set of a type determines the interfaces that the type implements and
  the methods that can be called using a receiver of that type.
*/

type Hero struct {
	name string
	age int
}

type EditHero interface {
	ChangeName(v string)
	ChangeAge(v int)
	SetName(v string)
}


func main() {
	var h1 = Hero{"IronMan", 45}
	var h2 = Hero{"CaptainC", 34}

	fmt.Println(h1)
	fmt.Println(h2.String())

	(&h2).ChangeName("CaptainA") // &h2 and h2 both works with (h *hero) as golang implicitly converts h2 to pointer type
	h2.ChangeAge(23)
	fmt.Println(h2)


	// antMan := new(Hero) or := &Hero{}
	var antMan *Hero
	antMan = new(Hero) // func new(Type) *Type
	antMan.SetName("AntMan")
	fmt.Println(antMan) // fmt.Println((*antMan).String())

	map0 := map[string]*Hero{"AntMan": antMan}
	fmt.Println(map0)
	map0["AntMan"].SetName("New AntMan")
	fmt.Println(map0)

	map1 := map[string]Hero{"CaptainA": h2}
	fmt.Println(map1)
	//map1["CaptainA"].SetName("New CaptainA") // cannot call pointer method on map1["CaptainA"] // cannot take the address of map1["CaptainA"]
	//&map1["CaptainA"].SetName("New CaptainA") // map1["CaptainA"].SetName("New CaptainA") used as value
	tmpHero := map1["CaptainA"]
	tmpHero.SetName("New CaptainA")
	fmt.Println(map1)
	fmt.Println(tmpHero)
}


func (h *Hero) ChangeName(v string) {
	h.name = v
}

func (h *Hero) SetName(v string) {
	h.name = v
}


// if h is of type Hero not *Hero then change to h won't reflect outside as it gets copy of Hero not reference
func (h *Hero) ChangeAge(v int) {
	h.age = v
}

func (h Hero) String() string {
	return fmt.Sprintf("SuperHero: %v, %v", h.name, h.age)
}