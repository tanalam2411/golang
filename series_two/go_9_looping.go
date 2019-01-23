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

	fmt.Println("--- nested loop")
	// exit from inner loop using label
	// here Loop is the label name where it get back once called using break statement.
	// Loop is not a keyword so you can assign different name to it.

Loop:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println(i * j)
			if i*j >= 3 {
				break Loop
			}
		}
	}

	fmt.Println("********************* Iterate over collection *********************")

	s := []int{1, 2, 3}

	for i, v := range s {
		fmt.Println(i, v)
	}

	m := make(map[string]int)
	m["vm"] = 100
	m["router"] = 10
	m["switch"] = 5

	for k, v := range m {
		fmt.Println(k, v)
	}

	// Just iterate over keys

	for k := range m {
		fmt.Println("Key: ", k)
	}

	// Just iterate over values

	for _, v := range m {
		fmt.Println("Value: ", v)
	}

	// Iterate string

	st := "hello world !!!"

	for i, v := range st {
		fmt.Println(i, string(v))
	}
}

/* SUMMARY

- For statements
  - simple loops
    - for initializer; test; incrementer {}
    - for test {}
    - for {} //while True
  - exiting early
    - break
    - continue
    - labels
  - looping over collections
    - arrays, slices, maps, strings, channels
    - for k, v := range collection {}

*/
