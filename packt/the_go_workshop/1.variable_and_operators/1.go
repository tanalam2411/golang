package main

import (
	"errors" // go doc builtin.error
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)

var helloList = []string{
	"Aaa",
	"Bbb",
	"Ccc",
}

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
}

func hello(index int) (string, error) {
	if index < 0 || index >= len(helloList) {
		return "", errors.New("Index out of range: " + strconv.Itoa(index))
	}
	return helloList[index], nil
}
