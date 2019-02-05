package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

/* AGENDA

- Creating goroutines
- Synchronization
  - WaitGroups
  - Mutexes
- Parallelism
- Best practices

*/

// WaitGroupp - is designed to synchronize multiple goroutines
var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{} // Read Write mutex - multiple Reads single write. When reading will not be able to write and when writing will not be able to read

func main() {
	fmt.Println("----------withOutWG---------------")
	withOutWG()

	fmt.Println("----------anonymousFunc---------------")
	anonymousFunc()

	fmt.Println("----------anonymousFuncWG---------------")
	anonymousFuncWG()

	fmt.Println("----------multiGoRouWithSameData---------------")
	multiGoRouWithSameData()

	fmt.Println("----------goMaxProcs---------------")
	goMaxProcs()
}
func sayHello() {
	fmt.Println("Hello Go")
}

func withOutWG() {
	go sayHello()
	// try commenting time.Sleep line
	time.Sleep(100 * time.Millisecond)
}

func anonymousFunc() {
	// race condition
	var msg = "One"
	go func() {
		fmt.Println("One :- ", msg)
	}()
	msg = "two"
	// time.Sleep(100 * time.Millisecond)

	go func(msg string) {
		fmt.Println("Two :- ", msg)
	}(msg)
	msg = "three"
	time.Sleep(100 * time.Microsecond)
}

func anonymousFuncWG() {
	var msg = "One"
	wg.Add(1) // add one to wg
	go func(msg string) {
		fmt.Println("msg: ", msg)
		wg.Done() // decrement 1 within wg
	}(msg)
	msg = "Two"
	wg.Wait()
}

func multiGoRouWithSameData() {

	// spwaning 20 goroutines
	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go showInc()
		m.Lock()
		go increment()
	}
	wg.Wait()
}

func showInc() {
	// m.RLock()
	fmt.Println("Incr is: ", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	// m.Lock()
	counter++
	m.Unlock()
	wg.Done()
}

func goMaxProcs() {
	// returns default number of OS threads
	fmt.Println("Threads: ", runtime.GOMAXPROCS(-1))

	// We can limit number of threads utilized by our application like this
	runtime.GOMAXPROCS(1)
	fmt.Println("Threads: ", runtime.GOMAXPROCS(-1))

	// General practise one OS thread per core
	// Assiging larger number to GOMAXPROCS can give you better performance,
	// But needs to be tested the limit by which its giving you best performance.
	// Because if its larger than expected then it will give negative result as most of the time will be spent in scheduling and managing threads
	runtime.GOMAXPROCS(100)
	fmt.Println("Threads: ", runtime.GOMAXPROCS(-1))
}

/* BEST PRACTICES

- Don't create goroutines in libraries
  - Let consumer control concurrency
- When creating a goroutine, know how it will end
  - Avoids subtle memory leaks
- Check for race condition at compile time

*/

/* SUMMARY

- Creating goroutines
  - Use go keyword in front of function call
  - When using anonymous functions, pass data as local variables

- Synchrounization
  - Use sync.WaitGroup to wait for groups of goroutines to complete
  - USe sync.Mutex and sync.RWMutex to protect data access

- Parallelism
  - By default, Go will use CPU threads equal to available cores
  - Change with runtime.GOMAXPROCS
  - More threads can increase performance, but too many can slow it down

- Main function itself is a goroutine
*/
