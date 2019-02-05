package main

import (
	"fmt"
	"log"
)

/* AGENDA
- Defer
- Panic
- Recover
*/

func panicker() {
	// defer func() - anonymous function
	// defer takes function call not just function
	// recover() function exits from the current function but continues within the function (parent) its being called.

	fmt.Println("about tp panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error: ", err)
		}
	}()

	panic("Some exception occured")
	fmt.Println("Done panicking")

	/*
		start
		about tp panic
		2019/01/24 12:42:42 Error:  Some exception occured
		end
	*/
}

func repanicker() {
	// In case where you think you cannot recover or you can't handle the panic, you can panic again
	// which should be handled by the upper stack
	fmt.Println("about tp panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error: ", err)
			panic(err)
		}
	}()

	panic("Some exception occured")
	fmt.Println("Done panicking")

	// notice - "repanicker : end" didn't got printed in the main function, because the program exited
	/*
		repanicker : start
		about tp panic
		2019/01/24 12:48:19 Error:  Some exception occured
		panic: Some exception occured [recovered]
				panic: Some exception occured

		goroutine 1 [running]:
		main.repanicker.func1()
				/home/tan/tanveer/golang/series_two/go_12_panic_recover.go:37 +0xb9
		panic(0x49f020, 0x4d3e30)
				/usr/local/go/src/runtime/panic.go:513 +0x1b9
		main.repanicker()
				/home/tan/tanveer/golang/series_two/go_12_panic_recover.go:41 +0x99
		main.main()
				/home/tan/tanveer/golang/series_two/go_12_panic_recover.go:58 +0xeb
		exit status 2
	*/
}

func main() {

	fmt.Println("panicker : start")
	panicker()
	fmt.Println("panicker : end")

	fmt.Println("repanicker : start")
	repanicker()
	fmt.Println("repanicker : end")
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
