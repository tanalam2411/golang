package main

import "fmt"



/* Objective

- Empty interface
- Wrapping one type within another
 - How to discover the hidden type of an expression?
- Type Switching
*/


type Developer struct {
	name string
	age int
}


type Printer interface {
	Print()
}


type IRunner struct {
	name string
}


func main() {

	var d1 = Developer{"Dev1", 32}
	var d2 = Developer{"Dev1", 27}
	m0 := map[string]Developer{"1st": d1}

	myPrinter(5)
	myPrinter("Some random string")
	myPrinter(d1)
	myPrinter(d2)
	myPrinter(m0)
	myPrinter(IRunner{"Quick"})

}


func myPrinter(i interface{}) {

	d, ok := i.(Developer)

	fmt.Printf("Developer: %v, ok:%v \n", d, ok)

	switch i.(type) {
	case int:
		fmt.Printf("MyPrinter (int): - %T -> %v \n", i, i)
	case string:
		fmt.Printf("MyPrinter (string): - %T -> %v \n", i, i)
	case map[string]Developer:
		fmt.Printf("MyPrinter (map[string]Developer): - %T -> %v \n", i, i)
	default:
		fmt.Println("Ignoring others")
	}
}


func(i IRunner) String() string {
	return fmt.Sprintf("Duck - %v", i.name)
}


func (d Developer) String() string {
	return fmt.Sprintf("Developer: %v, %v", d.name, d.age)
}