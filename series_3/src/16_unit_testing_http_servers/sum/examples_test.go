package sum_test

import (
	"fmt"
	"golang/series_3/src/16_unit_testing_http_servers/sum"
)

func ExampleInts() {
	s := sum.Ints(1, 2, 3, 4, 5)
	fmt.Printf("sum of 1 to 5 is: %v", s)
	// Output:
	// sum of 1 to 5 is: 15
}
