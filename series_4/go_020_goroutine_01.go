package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	go producer(c)
	go consumer(c)

	time.Sleep(50000 * time.Nanosecond)
}


func producer(out chan<- int) {

	for i:=0; i< 10; i++ {
		out <- i*i
	}

	//close(out)
}

func consumer(in <-chan int) {
	for {
		select{
		case v := <-in:
			fmt.Println("Consuming: ", v)
		//default:
		//	fmt.Println("Waiting for producer...")
		}
	}
}