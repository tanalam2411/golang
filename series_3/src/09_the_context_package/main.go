package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	// this is call the cancel(and the sleepAndTalk will stop) if user enters anything
	go func() {
		s := bufio.NewScanner(os.Stdin)
		s.Scan()
		cancel()
	}()

	// or Just cancel the sleepAndTalk func after a second
	time.AfterFunc(time.Second, cancel)

	sleepAndTalk(ctx, 5*time.Second, "Hello")

}


func sleepAndTalk(ctx context.Context, d time.Duration, msg string) {
	select {
	case <-time.After(d):
		fmt.Println(msg)
	case <-ctx.Done():
		log.Print(ctx.Err())
	}
}