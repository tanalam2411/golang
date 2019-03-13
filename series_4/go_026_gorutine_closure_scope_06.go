package main

import (
	"fmt"
	"sync"
)

func main() {

	fmt.Println("-------------------- caseOne --------------------------")
	caseOne()

	fmt.Println("-------------------- caseTwo --------------------------")
	caseTwo()

	fmt.Println("-------------------- caseThree --------------------------")
	caseThree()
}


func caseOne() {

	var wg sync.WaitGroup

	for i := 0 ; i< 5; i++ {

		wg.Add(1)
		go func() {
			fmt.Println("Case one - parent function scope: ", i)
			wg.Done()
		}()
	}

	wg.Wait()
}


func caseTwo() {

	var wg sync.WaitGroup

	for i := 0 ; i< 5; i++ {
		ii := i
		wg.Add(1)
		go func() {
			fmt.Println("Case two - parent function scope: ", ii)
			wg.Done()
		}()
	}

	wg.Wait()
}


func caseThree() {

	var wg sync.WaitGroup

	for i := 0 ; i< 5; i++ {
		wg.Add(1)
		go func(j int) {
			fmt.Println("Case three - parent function scope: ", j)
			wg.Done()
		}(i)
	}

	wg.Wait()
}


/* ------------ Output --------------------
-------------------- caseOne --------------------------
Case one - parent function scope:  5
Case one - parent function scope:  5
Case one - parent function scope:  5
Case one - parent function scope:  5
Case one - parent function scope:  5
-------------------- caseTwo --------------------------
Case two - parent function scope:  4
Case two - parent function scope:  0
Case two - parent function scope:  1
Case two - parent function scope:  2
Case two - parent function scope:  3
-------------------- caseThree --------------------------
Case three - parent function scope:  4
Case three - parent function scope:  0
Case three - parent function scope:  1
Case three - parent function scope:  2
Case three - parent function scope:  3
 */