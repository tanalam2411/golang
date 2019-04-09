package main

import (
	"fmt"
	"os"
)

func main() {

	name := ReadInput("Name: ")
	age := ReadInput("Age: ")
	fmt.Printf("Name: %v, age: %v \n", name, age)

}


func ReadInput(p string) (d string) {

	buf := make([]byte, 100)
	fmt.Println("%v: ", p)

	n, err := os.Stdin.Read(buf)

	if err != nil {
		return ""
	}

	d = string(buf[:n-1])
	return
}