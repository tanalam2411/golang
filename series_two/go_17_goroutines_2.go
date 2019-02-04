package main

import (
	"fmt"
	"time"
)

// Check for race condition at compile time

func main() {
	msg := "Hello Go!"

	go func() {
		fmt.Println(msg)
	}()
	msg = "Bye Go!"
	time.Sleep(100 * time.Millisecond)
}

/*

$ go run -race go_17_goroutines_2.go
==================
WARNING: DATA RACE
Read at 0x00c00008a1c0 by goroutine 6:
  main.main.func1()
      /home/tan/tanveer/golang/series_two/go_17_goroutines_2.go:14 +0x3c

Previous write at 0x00c00008a1c0 by main goroutine:
  main.main()
      /home/tan/tanveer/golang/series_two/go_17_goroutines_2.go:16 +0x9c

Goroutine 6 (running) created at:
  main.main()
      /home/tan/tanveer/golang/series_two/go_17_goroutines_2.go:13 +0x8e
==================
Bye Go!
Found 1 data race(s)
exit status 66

*/

// Comment line 16 -> msg = "Bye Go!" // and  try race flag
/*
$ go run -race go_17_goroutines_2.go
Hello Go!
*/
