package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	q := "query"

	now := time.Now()
	timeout := time.After(1000 * time.Millisecond)

	r0 := search(q)
	r1 := search(q)
	r2 := search(q)
	r3 := search(q)

	select {
	case t := <-r0:
		fmt.Printf("Search0 took %v\n", t.Sub(now))
	case t := <-r1:
		fmt.Printf("Search1 took %v\n", t.Sub(now))
	case t := <-r2:
		fmt.Printf("Search2 took %v\n", t.Sub(now))
	case t := <-r3:
		fmt.Printf("Search3 took %v\n", t.Sub(now))
	case <- timeout:
		fmt.Printf("Too slow, %v", <-timeout)
	}

	fmt.Println("------------------- selectDefault -------------")
	selectDefault()
}


func search(q string) <-chan time.Time {
	out := time.After(time.Duration(rand.Intn(500)) * time.Millisecond)

	return out
}


func selectDefault () {

	c := make(chan int, 2)

	for {
		select {

		case v := <-c:
			fmt.Printf("v: %v", v)
			os.Exit(v)
		default:
			c <- 1
		}
	}


}