package main

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
)

type Student struct {
	name string
	age uint8
	ssn string
}


func raiseError(n int, err error) {
	if err != nil {
		log.Fatalf("n: %v, err: %v \n", n, err)
	}
}

func main() {

	stud1 := Student{"stud_1", 19, "100-01-1100"}
	stud2 := Student{"stud_2", 19, "100-01-1101"}
	stud3 := Student{"stud_3", 18, "100-01-1102"}

	var buf = make([]byte, 50)
	offset, n := 0,0
	var err error

	n, err = stud1.Read(buf[offset:])
	offset += n
	raiseError(n, err)

	n, err = stud2.Read(buf[offset:])
	offset += n
	raiseError(n, err)

	n, err = stud3.Read(buf[offset:])
	offset += n
	raiseError(n, err)

	fmt.Println(buf)
}

/*
format: [
			<name-len>: byte
			<name>: []byte
			<age>: byte
			<ssn>: []byte
*/


func (s Student) Read(b []byte) (n int, err error) {

	var l = len(s.name)
	var buf = make([]byte, l+1)

	// encode s.name

	for i := 0; i < l; i++{
		buf[i+1] = s.name[i]
	}

	// encode r.age
	buf = append(buf, byte(s.age))


	// encode s.ssn
	l = len(s.ssn)
	for i := 0; i< l; i++{
		buf = append(buf, s.ssn[i])
	}

	if len(b) < len(buf) {
		e := fmt.Sprintf("Not enough buffer to write. needed: %v, have: %v", len(buf), len(b))
		return 0, errors.New(e)
	}

	n = copy(b, buf)

	return n, nil
}


func (s Student) String() string{
	return fmt.Sprintf("Student -> name: %v, age: %v, ssn: %v", s.name, s.age, s.ssn)
}