package main

import (
	"fmt"
	"math"
)

/*
A channel provides a mechanism for concurrently executing functions to communicate by sending
and receiving values of a specified element type.
The value of an uninitialized channel is nil.

Associativity of <-

Channel of Channels                  Meaning
chan<- chan int                    chan<- (chan int)
chan<- <-chan int                  chan<- (<-chan int)
<-chan <-chan int                  <-chan (<-chan int)
chan (<-chan int)
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


	fmt.Println("------------------ sendOnlyChannel ---------------")
	sendOnlyChannel()

	fmt.Println("------------------- funcReturningChannel ------------------")
	funcReturningChannel()
}


func sendOnlyChannel() {

	var producerChannel chan int
	producerChannel = make(chan int, 5)

	canOnlyWrite(producerChannel)
	fmt.Println("len of producerChannel: ", len(producerChannel))

	//for each := range producerChannel {
	//	fmt.Println("each: ", each)
	//}
	canOnlyRead(producerChannel)


	fmt.Println("len of producerChannel: ", len(producerChannel))
}


func canOnlyWrite(c chan<- int) {
	/*
	chan<-  // means you can only to this channel
	so if we try to read from it, it will throw an error
	v := <-c
	or if you pass this channel to another function and that trys to read from it, even then it will throw an error

	channel needs to be closed other wise while iterating it when when there is no element in the channel,
	the for loop will still be in waiting state as channel is not closed.
	 */
	c <- 10
	c <- 20
	c <- 30
	close(c)
}


func canOnlyRead(c <-chan int) {

	/*
	c <- 40 // is not allowed as the channel is receive only <-
	 */

	for v := range c{
		fmt.Printf("\n Reading value v: %#v \n", v)
	}


}


func funcReturningChannel() {


	c := sendReadOnlyChannel()
	out := squRootNums(c)

	canOnlyRead(out)
}

func sendReadOnlyChannel() <-chan int {

	out := make(chan int, 5)
	out <- 100
	out <- 200
	out <- 300
	close(out)
	return out
}

func squRootNums(c <-chan int) <-chan int {

	var result chan int

	result = make(chan int, len(c))

	for num := range c {
		x := float64(num)
		result <- int(math.Sqrt(x))
	}

	close(result)
	return result
}