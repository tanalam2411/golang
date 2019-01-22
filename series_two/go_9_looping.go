package main

import (
	"fmt"
)

/* AGENDA

- For statements
  - simple loops
  - exiting early
  - looping through collections
*/

func main() {

	// i++ is not an expression, its an statement
	// scope of i is within for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for i, j := 0, 0; i < 5; i, j = i+1, j+1 { // here i, j = i++, j++ will throw an error -> syntax error: unexpected ++, expecting { after for clause
		fmt.Println(i, j)
	}

	i := 0
	// here scope of i is within this whole module
	// for i<9; i++ -> syntax error
	// here we have outer incremenet i++ and also inner increment i = i + 3,
	// so in this case both increment  will be executed
	for ; i < 9; i++ {
		fmt.Println(i)
		i = i + 3
	}

	j := 0
	// here scope of i is within this whole module
	for j < 9 { // this systax can be used as while loop
		fmt.Println("j: ", j)
		j += 3
	}

	// Infinite loop -> while True // in Python
	flag := false
	for {
		fmt.Println("forverver")
		if flag {
			break
		}
		flag = true
	}

	fmt.Println("************************************ continue statement *******************************************")
	// continue statement

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}
