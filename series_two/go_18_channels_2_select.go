package main

import (
	"fmt"
	"time"
)

/* AGENDA

- Select statements

*/

const (
	logInfo    = "INFO"
	logWarning = "WARNING"
	logError   = "ERROR"
)

type logEntry struct {
	time     time.Time
	severity string
	message  string
}

var logCh = make(chan logEntry, 50)

// This is a signal only channel
// It can't send any msg, just sends or receive signal
var doneCh = make(chan struct{}) // struct with no feilds requires 0 memory

func main() {
	go logger()

	logCh <- logEntry{time.Now(), logInfo, "App is starting"}

	logCh <- logEntry{time.Now(), logWarning, "App is shutting down"}

	time.Sleep(100 * time.Millisecond)

	doneCh <- struct{}{} // struct{} is the type and {} is assigning empty values to it

}

func logger() {
	for {
		select {
		case entry := <-logCh:
			fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
		case <-doneCh:
			fmt.Println("Breaking go routine")
			break
			// default:
			// fmt.Println("------")
		}
	}
}

/* SUMMARY

- Select statements
  - Allows goroutine to monitor serverl channels at once
	- Blocks if all channels block
	- if multiple channels receive value simultaneously, behavior is undefined
*/
