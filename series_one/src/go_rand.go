package main

import (
	"fmt"
	"math"
	"math/rand"
		"time"
)

/*


Math Rand - https://golang.org/pkg/math/rand/
Math Complete Package - https://golang.org/pkg/math/

----------------------------------------------------------------

Godoc - https://blog.golang.org/godoc-documenting-go-code
Godoc commands - https://godoc.org/golang.org/x/tools/cmd/godoc
init function - https://golang.org/doc/effective_go.html#init

----------------------------------------------------------------

: Rand always generate same number
https://stackoverflow.com/questions/39529364/go-rand-intn-same-number-value
https://github.com/golang/tour/issues/267
https://stackoverflow.com/questions/12321133/golang-random-number-generator-how-to-seed-properly

----------------------------------------------------------------

: annotation [0, n) -> [ means including and ) means not including. So [0,n) means 0 till n-1

----------------------------------------------------------------

: Formats your code 
gofmt -s -w .           // formats all files in current directory
gofmt -s -w gotut.go    // formats provided file



*/

func random_generator() {

	// godoc fmt Println
	// godoc math Sqrt
	fmt.Println("Sqaure root of 4: ", math.Sqrt(4))

	// godoc math/rand Intn
	fmt.Println("Random number between 1-100:", rand.Intn(100))

	//cgodoc math/rand Float64
	fmt.Println("Random number between 0.0<= f < 1.0: ", rand.Float64())

	fmt.Println("Random number between 5.0<= f < 10.0: ", float64(rand.Intn(10-5)+5)+rand.Float64())
}

func main() {

	// Math Rand - https://golang.org/pkg/math/rand/
	// Seed : is a factor that determines the randomness that you'll get
	// it is usually implemented this way so that results across the board are same if seed is not randomized
	rand.Seed(time.Now().UTC().UnixNano())

	random_generator()
}
