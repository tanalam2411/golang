package main

import (
	"fmt"
	"sync"
)

/* AGENDA

- Channel basics
- Restricting data flow
- Buffered channels
- Closing channels
- For...range loops with channels
- Select statements

https://blog.golang.org/share-memory-by-communicating
*/

var wg = sync.WaitGroup{}

func main() {
	fmt.Println("*********** simpleSendRecOverChan ****************")
	simpleSendRecOverChan()

	fmt.Println("*********** multiSenderRecverOverSingleChn ****************")
	multiSenderRecverOverSingleChn()

	fmt.Println("*********** multiSenderSingleRecverOverSingleChn ****************")
	// multiSenderSingleRecverOverSingleChn()

	fmt.Println("*********** actingBothSenderRecv ****************")
	actingBothSenderRecv()

	fmt.Println("*********** sendOnlyRecOnlyCH ****************")
	sendOnlyRecOnlyCH()

	fmt.Println("*********** bufferedChannel ****************")
	bufferedChannel()

	fmt.Println("*********** bufferedChannelForOk ****************")
	bufferedChannelForOk()

}

func simpleSendRecOverChan() {
	ch := make(chan int) // channel is strongly typed

	wg.Add(2)

	go func() {
		i := <-ch
		fmt.Println("receive: ", i)
		wg.Done()
	}()

	go func() {
		i := 42
		ch <- i // sends copy of the data
		i = 23
		wg.Done()
	}()

	wg.Wait()
}

func multiSenderRecverOverSingleChn() {
	ch := make(chan int) // unbuffered channel, only one msg can stay within this channel

	for j := 0; j < 5; j++ {
		wg.Add(2)

		go func() {
			i := <-ch
			fmt.Println("receive: ", i)
			wg.Done()
		}()

		go func() {
			ch <- 43
			wg.Done()
		}()

	}
	wg.Wait()
}

func multiSenderSingleRecverOverSingleChn() {
	ch := make(chan int)

	wg.Add(1)
	go func() {
		i := <-ch
		fmt.Println("receiver: ", i)
		wg.Done()
	}()

	for j := 0; j < 5; j++ {

		wg.Add(1)
		go func() {
			ch <- 45
			wg.Done()
		}()

	}
	wg.Wait()

	/* this will throw runtime error, as the channel is unbuffered channel, so only one msg can exist within channel.
	   But here we're sending 5 msg to channel and only one recevier, that's why deadlock occured.
	*/

	/*
	   receiver:  45
	   fatal error: all goroutines are asleep - deadlock!

	   goroutine 1 [semacquire]:
	   sync.runtime_Semacquire(0x5690b0)
	           /usr/local/go/src/runtime/sema.go:56 +0x39
	   sync.(*WaitGroup).Wait(0x5690a8)
	           /usr/local/go/src/sync/waitgroup.go:130 +0x64
	   main.multiSenderSingleRecverOverSingleChn()
	           /home/tan/tanveer/golang/series_two/go_18_channels.go:94 +0xe1
	   main.main()
	           /home/tan/tanveer/golang/series_two/go_18_channels.go:29 +0xf0

	   goroutine 19 [chan send]:
	   main.multiSenderSingleRecverOverSingleChn.func2(0xc00001a180)
	           /home/tan/tanveer/golang/series_two/go_18_channels.go:89 +0x37
	   created by main.multiSenderSingleRecverOverSingleChn
	           /home/tan/tanveer/golang/series_two/go_18_channels.go:88 +0xc3

	   goroutine 20 [chan send]:
	   main.multiSenderSingleRecverOverSingleChn.func2(0xc00001a180)
	           /home/tan/tanveer/golang/series_two/go_18_channels.go:89 +0x37
	   created by main.multiSenderSingleRecverOverSingleChn
	           /home/tan/tanveer/golang/series_two/go_18_channels.go:88 +0xc3

	   goroutine 21 [chan send]:
	   main.multiSenderSingleRecverOverSingleChn.func2(0xc00001a180)
	           /home/tan/tanveer/golang/series_two/go_18_channels.go:89 +0x37
	   created by main.multiSenderSingleRecverOverSingleChn
	           /home/tan/tanveer/golang/series_two/go_18_channels.go:88 +0xc3

	   goroutine 22 [chan send]:
	   main.multiSenderSingleRecverOverSingleChn.func2(0xc00001a180)
	           /home/tan/tanveer/golang/series_two/go_18_channels.go:89 +0x37
	   created by main.multiSenderSingleRecverOverSingleChn
	           /home/tan/tanveer/golang/series_two/go_18_channels.go:88 +0xc3
	   exit status 2
	*/
}

func actingBothSenderRecv() {
	ch := make(chan int)

	wg.Add(2)
	go func() {
		i := <-ch
		fmt.Println(i)
		ch <- 27
		wg.Done()
	}()

	go func() {
		ch <- 45
		fmt.Println(<-ch)
		wg.Done()
	}()

	wg.Wait()
}

func sendOnlyRecOnlyCH() {
	ch := make(chan int)
	wg.Add(2)

	go func(ch <-chan int) {
		// Receive only channel
		i := <-ch
		fmt.Println(i)
		// ch <- 27 // will throw an error as its receive only channel
		wg.Done()
	}(ch)

	go func(ch chan<- int) {
		ch <- 21
		wg.Done()
	}(ch)

	wg.Wait()
}

func bufferedChannel() {
	ch := make(chan int, 50)
	wg.Add(2)

	go func(ch <-chan int) {
		for i := range ch {
			fmt.Println(i)
		}
		wg.Done()
	}(ch)

	go func(ch chan<- int) {
		ch <- 10
		ch <- 20
		close(ch) // its nessacary to close channel, or else deadlock will occur as the receiver's for loop will not know when to stop.
		// ch <- 30 // will panic as you can't send msg on closed channel
		wg.Done()
	}(ch)

	wg.Wait()
}

func bufferedChannelForOk() {
	ch := make(chan int)
	wg.Add(2)

	go func(ch <-chan int) {

		for {
			if i, ok := <-ch; ok {
				fmt.Println(i)
			} else {
				break
			}
		}

		wg.Done()
	}(ch)

	go func(ch chan<- int) {
		ch <- 100
		ch <- 200
		close(ch)

		wg.Done()
	}(ch)

	wg.Wait()
}

/* SUMMARY

- Channel basics
  - create a channel with make command
	- make(chan int)
  - Send message into channel
	- ch <- val
  - Receive message from channel
	- val := <- ch
  - Can have multiple senders and receivers

- Restricting data flow
  - Channel can be cast into send-only or receive only versions
	- Send-only: chan <- int
	- Receiver-only: <-chan int

- Buffered channels
  - Channels block sender side till receiver is available
  - Block receiver side till message available
  - Can decouple sender and receiver with buffered channels
	- make(chan int, 50)
  - Use buffered channels when send and receive have assymmetric loading

- For...range loops with channels
  - Use to monitor channel and process messages as they arrive
  - Loops exits when channel is closed

*/
