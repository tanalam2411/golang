package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main(){
	rand.Seed(time.Now().Unix())

	var wg sync.WaitGroup

	chatter("station 1 ", &wg)
	chatter("station 2 ", &wg)
	chatter("station 3 ", &wg)
	//chatter("station 4 ", &wg)

	wg.Wait()

	fmt.Println("Done")
}

func chatter(s string, wg *sync.WaitGroup) {
	//talkTime := rand.Intn(3000)
	talkTime := 300000
	timeUp := time.After(time.Duration(talkTime) * time.Nanosecond)

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			select {
			case <-timeUp:
				fmt.Println(s, " is qutting after: ", time.Duration(talkTime))
				return
			default:
				fmt.Println(s)
			}
		}
	}()
}