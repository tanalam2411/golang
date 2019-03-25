package main

import "fmt"

/*
Todo: close.go, ms.go, read.go, reset.go, string.go, write.go // move this to new package

*/
func main() {


	var w0 = []byte{1, 2, 3, 4, 5, 6}
	var w1 = []byte{7, 8, 9, 10}

	var m MemStore

	m.Write(w0)
	m.Write(w1)

	fmt.Println(m)

	var r0 = make([]byte, 4)
	var n int
	var err error

	for err == nil {
		n, err = m.Read(r0)
		fmt.Println(n, err, r0)
	}


	m.Reset()
	n, err = m.Read(r0)
	fmt.Println(n, err, r0)

	m.Close()

	n, err = m.Read(r0)

	if err != nil {
		fmt.Println(n, err)
		return
	}

	fmt.Println(n, err, r0)
}



/*
golang/series_4/go_040_memstore/cli$ go build

golang/src/golang/series_4/go_040_memstore/cli$ ll
total 1924
drwxrwxr-x 2 tan tan    4096 Mar 25 02:37 ./
drwxrwxr-x 4 tan tan    4096 Mar 24 23:54 ../
-rwxrwxr-x 1 tan tan 1932458 Mar 25 02:37 cli*
-rw-rw-r-- 1 tan tan      65 Mar 25 02:13 close.go
-rw-rw-r-- 1 tan tan     465 Mar 25 02:35 main.go
-rw-rw-r-- 1 tan tan     390 Mar 25 02:18 ms.go
-rw-rw-r-- 1 tan tan     312 Mar 25 02:22 read.go
-rw-rw-r-- 1 tan tan     175 Mar 25 02:23 reset.go
-rw-rw-r-- 1 tan tan     165 Mar 25 02:25 string.go
-rw-rw-r-- 1 tan tan     303 Mar 25 02:37 write.go

golang/series_4/go_040_memstore/cli$ ./cli
main.MemStore (isClosed: false, size: 10, buf: [1 2 3 4 5 6 7 8 9 10])
4 <nil> [1 2 3 4]
4 <nil> [5 6 7 8]
2 <nil> [9 10 7 8]
0 EOF [9 10 7 8]
4 <nil> [1 2 3 4]
0 Invalid read opertion on closed buffer.

golang/src/golang/series_4/go_040_memstore/cli$ go run main.go
# command-line-arguments
./main.go:11:8: undefined: MemStore

 */