package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("---------------------------")

	j := 0

	for j < 5 {
		fmt.Println(j)
		j += 1
	}

	fmt.Println("---------------------------")

	x := 5

	for {
		fmt.Println("While x is less than 25, now x is:", x)

		x += 5
		if x > 25 {
			fmt.Println("Done")
			break
		}
	}

	// Simlimar to
	for x := 5; x <= 25; x += 5 {
		fmt.Println("While x is less than 25, now x is:", x)
	}

	fmt.Println("------------Infinite loop---------------")

	// for {
	// 	fmt.Println("Will keep running forever. press ctrl+c to exit")
	// }
}
