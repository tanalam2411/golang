package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

func main() {

	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}

	defer trace.Stop()

	x, y := 2, 3
	fmt.Printf("Addition of %v and %v is equal to: %v \n", x, y, x+y)

}

/*
~/golang/go_tool_trace$ go run main.go
Addition of 2 and 3 is equal to: 5

tan@tan-ThinkPad-E450:~/tanveer/golang/go_tool_trace$ ll
total 16
drwxrwxr-x  2 tan tan 4096 Jan 27 23:42 ./
drwxrwxr-x 20 tan tan 4096 Jan 27 23:37 ../
-rw-rw-r--  1 tan tan  318 Jan 27 23:42 main.go
-rw-rw-r--  1 tan tan 1345 Jan 27 23:42 trace.out

~/golang/go_tool_trace$ go tool trace trace.out
2019/01/27 23:43:05 Parsing trace...
2019/01/27 23:43:05 Splitting trace...
2019/01/27 23:43:05 Opening browser. Trace viewer is listening on http://127.0.0.1:35331
*/
