package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Go concurrency pattern - fanin
*/

func main() {
	rand.Seed(time.Now().Unix())

	c0 := pusher("text1")
	c1 := pusher("text2")
	c2 := pusher("text3")
	c3 := pusher("text4")

	resChan := poper(c0, c1, c2, c3)

	go func() {
		for {
			fmt.Println(<-resChan)
		}
	}()

	time.Sleep(5 * time.Second)
}


func pusher(s string) <-chan string{
	out := make(chan string)

	go func() {

		for {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Nanosecond)
			out <- s
		}
	}()

	return out
}


func poper(ch0, ch1, ch2, ch3 <-chan string) <-chan string {

	out := make(chan string)

	reader := func(in <-chan string) {
		for v := range in {
			out <- v
		}
	}

	go reader(ch0)
	go reader(ch1)
	go reader(ch2)
	go reader(ch3)

	return out
}