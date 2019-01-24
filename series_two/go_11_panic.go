package main

import (
	"fmt"
	"net/http"
)

/* AGENDA
- Defer
- Panic
- Recover
*/

func runtimePanic() {
	a, b := 1, 0
	ans := a / b
	fmt.Println(ans)

	/*
		panic: runtime error: integer divide by zero

		goroutine 1 [running]:
		main.runtimePanic()
		        /home/tan/tanveer/golang/series_two/go_11_panic.go:9 +0x11
		main.main()
		        /home/tan/tanveer/golang/series_two/go_11_panic.go:14 +0x22
		exit status 2
	*/
}

func doPanic() {
	fmt.Println("start")
	panic("ooooops")
	fmt.Println("end")

	/*
		start
		panic: ooooops

		goroutine 1 [running]:
		main.doPanic()
		        /home/tan/tanveer/golang/series_two/go_11_panic.go:26 +0x79
		main.main()
		        /home/tan/tanveer/golang/series_two/go_11_panic.go:33 +0x20
		exit status 2
	*/
}

func simpleWebApp() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World !!!"))
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error())
	}

	/*
		$ go run go_11_panic.go
		panic: listen tcp :8080: bind: address already in use

		goroutine 1 [running]:
		main.simpleWebApp()
		        /home/tan/tanveer/golang/series_two/go_11_panic.go:51 +0xd0
		main.main()
		        /home/tan/tanveer/golang/series_two/go_11_panic.go:60 +0x20
		exit status 2
	*/
}

func deferExecutesBeforePanic() {
	// If panic occurs then before panicing all defer statements will get executed
	// But if defer is mentioned after panic statmenet then it will not get executed, panic will make an exit from there
	// meaning
	// panic("will panic here")
	// defer fmt.Println("this was deferred")
	// In this case defer will not get executed
	fmt.Println("start")
	defer fmt.Println("this was deferred")
	panic("will panic here")
	fmt.Println("end")

	/*
		start
		this was deferred
		panic: will panic here

		goroutine 1 [running]:
		main.deferExecutesBeforePanic()
		        /home/tan/tanveer/golang/series_two/go_11_panic.go:70 +0xd5
		main.main()
		        /home/tan/tanveer/golang/series_two/go_11_panic.go:81 +0x20
		exit status 2
	*/
}

func main() {
	// runtimePanic()

	// doPanic()

	// simpleWebApp()

	deferExecutesBeforePanic()
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
