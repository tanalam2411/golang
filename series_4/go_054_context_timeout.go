package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	Find()
	time.Sleep(10 * time.Second)
}

type Responsee struct {
	data   interface{}
	status bool
}

func Find() {
	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan Responsee, 1)

	go func() {
		time.Sleep(0 * time.Second)

		select {
		case <-ctx.Done():
			fmt.Println("Canceled by timeout")
			return
		}

		fmt.Println("test")
		fmt.Println("test1")

		ch <- Responsee{data: "data", status: true}
	}()

	select {
	case <-ch:
		fmt.Println("Read from ch")
	case <-time.After(50 * time.Millisecond):
		fmt.Println("Timed out")
		cancel()
	}

}
