package main

import "fmt"

type ServerResponse struct {
	q string
	a string
}


type CliResponse struct {
	q string
	a string
}


type StudentId int


type MyFunc func(i int, s string)


func main() {

	var sr = ServerResponse{"Q1", "A1"}
	var cr = CliResponse{"Q2", "A2"}
	fmt.Println(sr)
	fmt.Println(cr)

	var x StudentId = 5
	printStudentId(x)

	var mFun MyFunc = func (i int, s string) {
		//
	}

	go mFun(1, "I")

	goo(mFun)
}

func printStudentId(sId StudentId) {
	fmt.Println("Student id: ", sId)
}


func goo(f MyFunc) {
	f(100, "Hundred")
}