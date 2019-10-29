package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World !!!")
	evenNums(0, 10)
}


func evenNums(from, to int) {

	s := make([]int, 1)
	a := [10]int{1, 2}
	//a := make([]int, 10)

	fmt.Printf("%T, %v \n", a, a)
	fmt.Printf("%T, %v \n", s, s)

	s[0] = 100
	s = append(s, 1)
	s = append(s, 2, 3 ,4)
	a[0] = 1

	fmt.Printf("%T, %v \n", a, a)
	fmt.Printf("%T, %v \n", s, s)
	for i, v :=range s{
		fmt.Println(i, v)
	}
}