package main

import "fmt"

type BigB struct {
	name string
	age int
}


type OtherStringer interface {
	String() string
}


type Stringer interface {
	String() string
}


func main() {

	var b1 = BigB{"BigB-1", 33}
	var b2 = BigB{"BigB-2", 45}

	fmt.Println(b1)
	fmt.Println(b2)

	useMainStringer(b1)
	useMainOtherStringer(b2)

}


func useMainOtherStringer(s OtherStringer) {
	fmt.Printf("Implements main.OtherStringer -> %v\n", s)
}

func useMainStringer(s Stringer) {
	fmt.Printf("Implements main.Stringer -> %v\n", s)
}


func (b BigB) String() string{
	return fmt.Sprintf("BigB: %v, %v \n", b.name, b.age)
}