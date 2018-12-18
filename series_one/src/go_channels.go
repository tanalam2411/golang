package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func foo(c chan int, someValue int) {

	defer wg.Done()
	c <- someValue * 5
}

func main() {
	fooVal := make(chan int, 10)

	// go foo(fooVal, 5)
	// go foo(fooVal, 3)

	// v1 := <-fooVal
	// v2 := <-fooVal

	// v1, v2 := <-fooVal, <-fooVal

	// fmt.Println(v1, v2)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go foo(fooVal, i)
	}

	wg.Wait()
	close(fooVal)

	for item := range fooVal {
		fmt.Println(item)
	}

}
