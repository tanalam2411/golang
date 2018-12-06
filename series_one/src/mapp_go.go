package main

import (
	"fmt"
)

func main() {

	// map is a refernace, so to create map apply make on it
	TollFee := make(map[string]float32)

	TollFee["pune"] = 30
	TollFee["mumbai"] = 50
	TollFee["banglore"] = 60

	for city, fare := range TollFee {

		fmt.Println(city, ":", fare)
	}

}
