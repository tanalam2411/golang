package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	log.Println("for loop")

	now := time.Now()
	rand.Seed(now.Unix())

	var age int
	count := 1

	for ; ; count++ {
		age = rand.Intn(150)

		if age < 25 {
			fmt.Println("age less than 25, ", count)
			continue
		}

		if age < 100 {
			fmt.Println("Not 100 yet, ", count)
		} else {
			fmt.Println("Over 100 years great!!! ", count)
			break
		}

	}
}