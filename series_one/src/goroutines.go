package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Defer - works in first in last out.

defer fmt.Println("First")
defer fmt.Println("Second")

it will result in:
Second
First
*/

//wg - wait group
var wg sync.WaitGroup

func cleanup() {

	defer wg.Done()

	if r := recover(); r != nil {
		fmt.Println("recovered in cleanup: ", r)
	}
}
func say(s string) {

	defer cleanup()

	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)

		// Panic if numer is 2
		if i == 2 {
			panic("i is 2")
		}
	}

}

func main() {

	wg.Add(1)
	go say("go")
	wg.Add(1)
	go say("python")
	wg.Wait()

	// Below will only start once above wait group will finish.
	// Because of wg.Wait. Comment above wg.Wait to make it all concurrent.
	wg.Add(1)
	go say("Docker")
	wg.Wait()
}
