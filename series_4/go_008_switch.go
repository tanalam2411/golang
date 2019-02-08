package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	log.Println("switch statement -----\n")

	now := time.Now()
	rand.Seed(now.Unix())

	switchMatch()
	switchCondition()
}


func switchMatch() {

	day := strconv.Itoa(rand.Intn(7))

	switch day {

	case "0":
		fmt.Println("Sunday")
	case "1":
		fmt.Println("Monday")
	case "2":
		fmt.Println("Tuesday")
	case "3":
		fmt.Println("Wednesday")
	case "4":
		fmt.Println("Thusday")
	case "5":
		fmt.Println("Friday")
	default:
		fmt.Println("Saturday")
	}
}


func switchCondition() {
	age := rand.Intn(150)

	switch {

	case age < 18:
		fmt.Println("less than 18")
	case age < 65:
		fmt.Println("less than 65")
	case age < 65:
		fmt.Println("less than 100")
	default:
		fmt.Println("greater than 100")
	}
}