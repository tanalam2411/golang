package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().Unix())

	quit := make(chan bool)
	c0 := dumpMsg("Station 1", quit)
	c1 := dumpMsg("Station 2", quit)
	c2 := dumpMsg("Station 3", quit)
	c3 := dumpMsg("Station 4", quit)

	go func(){
		for {
			select {
				case v := <-c0:
					fmt.Println(v)
				case v := <-c1:
					fmt.Println(v)
				case v := <-c2:
					fmt.Println(v)
				case v := <-c3:
					fmt.Println(v)
			}
		}
	}()

	time.Sleep(2 * time.Second)
	quit <- true
	quit <- true
	quit <- true
	quit <- true
	time.Sleep(500 * time.Millisecond)
}

func dumpMsg(msg string, quit <-chan bool) <-chan string {
	out := make(chan string)

	go func() {
		for {
			select {
				case <- quit:
					fmt.Println(msg, " Exiting...")
					return
				case out <- msg:
					time.Sleep(time.Duration(rand.Intn(100)) * time.Nanosecond)

			}
		}
	}()

	return out
}


