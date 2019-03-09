package main

import "fmt"

/*
Channel selection - select statement
*/

func main() {

	quit := make(chan bool, 2)

	quit <- true
	quit <- false

	c := selectChannel(quit)

	for v := range c {
		fmt.Println(v)
	}
}


func selectChannel(quit <-chan bool) <-chan string {
	out := make(chan string, 10)

	flag := false

	for {

		fmt.Println("---")

		select {
			case out <- "First message in":
			case out <- "Second message in":
			case out <- "Third message in":
			case v := <-quit:
				if v {
					fmt.Printf("Exiting ...... with %v messages\n", len(out))
					flag = v
					break
				}
		}

		if flag{
			break
		}
	}

	close(out)
	return out
}

