package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Person1 struct {
	name string
	age int
	height float64
}


func main() {

	f, err := os.Open("data.txt")

	if err != nil {
		log.Fatalf("Unable to open file for reading: %v\n", err)
	}
	defer f.Close()

	header, err := getHeader(f)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%-15s %4s %8s\n",
				strings.Title(header[0]), strings.Title(header[1]),
				strings.Title(header[2]))

	var p *Person1

	for err == nil {
		p, err = getRow(f)

		if io.EOF == err {
			return
		}

		if nil != err {
			log.Fatalf("Error reading row: %v\n", err)
		}

		fmt.Printf("%-15v %4v %8.1f\n", p.name, p.age, p.height)
	}


}



func getRow(r io.Reader) (p *Person1, err error) {
	var name string
	var age int
	var height float64

	_, err = fmt.Fscanln(r, &name, &age, &height)

	if nil != err {
		return
	}

	p = &Person1{name: name, age: age, height: height}
	return
}


func getHeader(r io.Reader) (header []string, err error) {
	var input string
	_, err = fmt.Fscanln(r, &input)
	if err != nil {
		return
	}
}