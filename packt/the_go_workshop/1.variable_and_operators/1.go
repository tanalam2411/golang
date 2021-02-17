package main

import (
	"errors" // go doc builtin.error
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
)


var (
	helloList = []string{
		"Aaa",
		"Bbb",
		"Ccc",
	}
	var1 string = "abc"
	var2 int = 123
	var3 = 3.14
	var4 bool
)



func main() {
	fmt.Printf("Seed input value: %v\n", time.Now().UnixNano())
	rand.Seed(time.Now().UnixNano())

	index := rand.Intn(len(helloList))
	fmt.Printf("Random Index value: %v\n", index)

	if msg, err := hello(index); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(msg)
	}

	stringOperations()
	variables()
	logicalOperators()
}

func hello(index int) (string, error) {
	if index < 0 || index >= len(helloList) {
		return "", errors.New("Index out of range: " + strconv.Itoa(index))
	}
	return helloList[index], nil
}


func stringOperations() {
	if res := strings.Compare("abc", "abc"); res==0 {
		fmt.Printf("String Compare true: %v, %T\n", res, res)
	} else{
		fmt.Printf("String Compare true: %v, %T\n", res, res)
	}
}

func variables() {
	var s string
	fmt.Printf("s: %v, %T\n", s, s)

	s = "abc"
	fmt.Printf("s: %v, %T\n", s, s)

	var m = make(map[string]string)

	m["a"]  = "b"

	fmt.Printf("m: %v, %T\n", m, m)
}


func logicalOperators() {
	var v1 int
	var v2 int

	if v1==0 || v2==0 {
		fmt.Println("|| operators")
	} else{
		fmt.Printf("v1: %v1, v2: %v2", v1, v2)
	}

	if v1==0 && v2==0 {
		fmt.Println("&& operators")
	} else{
		fmt.Printf("v1: %v1, v2: %v2", v1, v2)
	}

	if !(v1==0 && v2==0) {
		fmt.Println("&& operators")
	} else{
		fmt.Printf("v1: %v, v2: %v\n", v1, v2)
	}
}