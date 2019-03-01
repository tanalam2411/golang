package main

import "fmt"

/*
A channel provides a mechanism for concurrently executing functions to communicate by sending
and receiving values of a specified element type.
The value of an uninitialized channel is nil.
*/


func main() {
	var c chan int
	fmt.Printf("c: %#v \n", c)
	fmt.Println("len of c: ", len(c))

	// initialize channel
	c = make(chan int, 2)  // if size 2 is not given it would be unbuffered channel, now its a buffered channel

	fmt.Printf("c: %#v \n", c)
	fmt.Println("len of c: ", len(c))

	// putting value in channel

	c <- 10
	c <- 20

	fmt.Printf("c: %#v \n", c)
	fmt.Println("len of c: ", len(c))

	// reading from channel
	var out = <- c
	fmt.Printf("out: %#v \n", out)


	fmt.Printf("c: %#v \n", c)
	fmt.Println("len of c: ", len(c))
}