package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)


func main() {
	log.Println("If statement")

	now := time.Now()
	rand.Seed(now.Unix()) // Seed should each time to generate new random number, that's why used time

	age := rand.Intn(150)

	if age < 18 {
		fmt.Println("Too young: ", age)
	} else if age < 64 {
		fmt.Println("Adult: ", age)
	} else if age < 100 {
		fmt.Println("Old: ", age)
	} else {
		fmt.Println("Ninjas: ", age)
	}

}
