package main


import (
	"fmt"
	"time"
)


/*
Go concurrency pattern - Generators
*/

func main() {

	c := msgProducer()
	go msgConsumer(c)
	time.Sleep(1 * time.Millisecond)
}


func msgProducer() chan int {
	out := make(chan int)

	go func() {
		fmt.Println("Producer")
		for i := 0; i < 20; i++{
			out <- i*i
		}
		close(out)
	}()


	return out
}


func msgConsumer(inputChn <-chan int) {
	fmt.Println("Consumer")

	for {
		select {
		case v := <-inputChn:
			fmt.Println(v)
		default:
			fmt.Println("Waiting for msg")
		}
	}
}