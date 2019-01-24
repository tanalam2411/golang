/*
Failing Forward
1) Strong typed - Once type is assigned to a variable it can't be changed.

2) Statically typed - All of the variables have to be defined at compile time.
   Go has alternative for it called type system.

3) Go has Garbage collection.

4) Built-in concurrency

5) Compile to standalone binaries (Single binary)

https://github.com/Microsoft/vscode-go/issues/441

*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello all...")
}

/*
Different ways to run and build go program

1) $ go run src/github.com/tanalam2411/series_two/go_introduction.go
Hello all...

2) $ go build github.com/tanalam2411/series_two
$ ./series_two
Hello all...

3) $ go install github.com/tanalam2411/series_two
$ ls bin/
series_two
$ bin/series_two
Hello all...
*/
