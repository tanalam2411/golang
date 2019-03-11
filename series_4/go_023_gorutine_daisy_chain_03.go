package main

import (
	"fmt"
)

/*
Go concurrency pattern - Daisy Chain/Pipeline
*/

const N = 200000


func main() {
	c := make(chan int)
	in := c
	var out chan int

	for i := 0; i < N; i++ {
		out = worker(in)
		in = out
	}


	c <- 0
	//close(c)
	fmt.Println(<-out)

	//time.Sleep(5 * time.Second)
}


func worker(in chan int) chan int {
	out := make(chan int)

	go func() {
		v := <- in
		out <- 1 + v
	}()

	return out
}