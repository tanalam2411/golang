package main

import (
	"fmt"
	"errors"
)

const SSN_LEN = 11


type Employee struct {
	name string
	age uint8  //0 - 127
	ssn string // (11 byte)
}


func main() {
	/*
	Byte Format: [
	             <name-len>:byte,
				 <bytes-of-name>:[]byte;
				 <age>:byte,
				 <bytes-of-ssn>:[11]byte
	 			]
	 */

	 var data = []byte{5, 66, 111, 98, 97, 121, 35, 57, 57, 49, 45, 57, 48, 45, 49, 48, 51, 50}

	 fmt.Println(data)

	 var e Employee
	 fmt.Println(e)

	 total_bytes, err := e.Write(data)
	 fmt.Printf("err: %v \n, Employee: %v \n", err, e)
	 fmt.Println("Total bytes: ", total_bytes, len(data))
}

func (e *Employee)Write(b []byte) (n int, err error) {
	if e == nil {
		return 0, errors.New("nil is not a valid Employee type.")
	}

	if len(b) == 0 {
		return 0, errors.New("No data found.")
	}

	nameLen := int(b[0])

	if (nameLen + 1) > len(b) {
		return 0, errors.New("Insufficient data for name.")
	}

	e.name = string(b[1: nameLen+1])

	if (nameLen + 2) > len(b) {
		return 0, errors.New("Insufficient data for age.")
	}
	e.age = uint8(b[nameLen+1])

	buf := b[nameLen+2:]
	if len(buf) < SSN_LEN {
		return 0, errors.New("Insufficient data for ssn.")
	}

	e.ssn = string(buf[:SSN_LEN])

	return (1 + nameLen + 1 + SSN_LEN), nil
}

func (e Employee)String() string{
	return fmt.Sprintf("Employee: \n\t Name: %v \n\t Age: %v \n\t SSN: %v \n", e.name, e.age, e.ssn)
}
