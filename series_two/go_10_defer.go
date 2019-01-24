package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/* AGENDA
- Defer
- Panic
- Recover
*/

func justDefer() {
	// defer executes the statement just `before the function return
	fmt.Println("start")
	defer fmt.Println("middle")
	fmt.Println("end")
}

func deferLifoOrder() {
	// defer will execute in last in first out order ->LIFO
	defer fmt.Println("----- 1")
	defer fmt.Println("----- 2")
	defer fmt.Println("----- 3")
}

func getRobots() {
	res, err := http.Get("http://www.ggogle.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("len of robots: ", len(robots))
}

func deferVariable() {
	a := "start"
	defer fmt.Println(a)
	a = "end"
}

func main() {

	fmt.Println("--------------- defer example Start -----------------")
	justDefer()
	fmt.Println("--------------- defer example Start -----------------")

	fmt.Println("--------------- LIFO example Start ------------------")
	deferLifoOrder()
	fmt.Println("--------------- LIFO example end --------------------")

	getRobots()

	fmt.Println("--------------- defer variable example Start ------------------")
	deferVariable()
	fmt.Println("--------------- defer variable example end --------------------")
}

/* SUMMARY

- Defer
  - Used to delay execution of a stetement until function exits
  - Useful to group "open" and "close" functions together
	- Be careful in loops (If your looping lots of resources and defering them, then all the resources will be
						   open until the for loop finishes. One way could calling another function within the forloop
						and handling defer inside that function)
  - Run in LIFO(last-in, first-out)order
  - Arguments evaluated at time defer is executed, not at time of called function exectution

- Panic
  - Occur when program cannot continue at all
	- Don't use when file can't be opened, unless it is critical
	- Use for unrecoverable events - cannot obtain TCP port for web server
  - Function will stop executing
	- Deferred functions will still fire
  - If nothing handles panic, program will exit

- Recover
  - Used to recover from panics
  - Only useful in deferred functions
  - Current function will not attempt to continue, but higher functions in call stack will
*/
