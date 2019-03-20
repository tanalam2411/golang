package main

import "fmt"

/*

An "interface" is a type that defines the set of operations/methods available
- A 'type' can implement or provide operations for several interfaces.

Example of interfaces -
Interface		Method

Duck			Quack()
				Waddle()

Stringer		String(): string

Reader			Read(...): len, err

Writer			Write(...): len, err
*/


/* Struct  vs Interface

Struct - Defines the memory layout of a type
         How should a value of this type be stored in memory, the memory footprint.

Interface - Defines which operations are available
            What actions can be taken in manipulation a value which implements this interface.


*/



type Kid struct {
	name string
	age int
}


type IndianRunner struct {
	name string
	weight int
}


type Watervale struct {
	name string
	weight int
}


type Duck interface {
	Quack()
	Waddle()
}


func main() {
	var kid1 = Kid{"Max", 30}
	var kid2 = Kid{"Trut", 29}

	printKid(kid1)
	fmt.Println(kid2)

	var duck0 = Watervale{"Swimmer", 10}
	//duck0.Waddle()
	callDuck(duck0)

	var duck1 = IndianRunner{"Speedy", 100}
	duck1.Waddle()
	duck1.Quack()

}


func printKid(k Kid) {
	fmt.Printf("Kid's name: %v, age %v \n", k.name, k.age)
}


func (k Kid) String() string {
	return fmt.Sprintf("Kid's name: %v, age %v \n", k.name, k.age)
}


func callDuck(d Duck) {
	d.Quack()
	d.Waddle()
}


func (d Watervale) Waddle() {
	fmt.Printf("Watervale Waddle - %v\n", d.name)
}

func (d Watervale) Quack() {
	fmt.Printf("Watervale Quack - %v\n", d.name)
}


func (d IndianRunner) Waddle() {
	fmt.Printf("IndianRunner Waddle - %v\n", d.name)
}

func (d IndianRunner) Quack() {
	fmt.Printf("IndianRunner Quack - %v\n", d.name)
}