**Go Tooling in Action**

1. Go code formatting:

```bash
$ cat <<EOF >main.go
> package main
> import "fmt"
> func main(){fmt.Println("Hello, there.")}
> EOF
``` 

 Look `gotmft --help`
 ```bash
$ gofmt --help
usage: gofmt [flags] [path ...]
  -cpuprofile string
    	write cpu profile to this file
  -d	display diffs instead of rewriting files
  -e	report all errors (not just the first 10 on different lines)
  -l	list files whose formatting differs from gofmt's
  -r string
    	rewrite rule (e.g., 'a[b:len(a)] -> a[b:]')
  -s	simplify code
  -w	write result to (source) file instead of stdout
```

Println the formatted version of go code on terminal(stdout):
```bash
$ gofmt main.go 
package main

import "fmt"

func main() { fmt.Println("Hello, there.") }
```

Show the difference
```bash
$ gofmt -d main.go
diff -u main.go.orig main.go
--- main.go.orig	2019-11-15 16:17:26.736988097 +0530
+++ main.go	2019-11-15 16:17:26.736988097 +0530
@@ -1,3 +1,5 @@
 package main
+
 import "fmt"
-func main(){fmt.Println("Hello, there.")}
\ No newline at end of file
+
+func main() { fmt.Println("Hello, there.") }
```

Simplify and apply the formatting to `main.go`. You can run this on whole dir by doing `gofmt -s -w .`

```bash
$ gofmt -s -w main.go

$ cat main.go 
package main

import "fmt"

func main() { fmt.Println("Hello, there.") }
```

2. Cross compiling using `GOOS`
```bash
$ GOOS=windows go build

$ ll
total 2076
drwxrwxr-x 2 tan tan    4096 Nov 15 16:26 ./
drwxrwxr-x 4 tan tan    4096 Nov 15 15:54 ../
-rwxrwxr-x 1 tan tan 2106880 Nov 15 16:26 go_tooling_in_action.exe*
-rw-rw-r-- 1 tan tan      73 Nov 15 16:20 main.go
-rw-rw-r-- 1 tan tan    1381 Nov 15 16:26 README.md


$ file go_tooling_in_action.exe 
go_tooling_in_action.exe: PE32+ executable (console) x86-64 (stripped to external PDB), for MS Windows
```

3. `go install` - It will build the binary and move it to `GOPATH/din`.
```bash
go_tooling_in_action$ go install

go_tooling_in_action$ $GOPATH/bin/go_tooling_in_action 
Hello, there.
```

4. `go get` - Downloads, compiles and installs the binary.

```bash
go_tooling_in_action$ go get github.com/golang/example/hello
go: finding github.com/golang/example latest
go: downloading github.com/golang/example v0.0.0-20170904185048-46695d81d1fa
go: extracting github.com/golang/example v0.0.0-20170904185048-46695d81d1fa


go_tooling_in_action$ $GOPATH/bin/hello 
Hello, Go examples!
```

5. `go doc and list` - extracting info from Go packages

Show import path of the current pkg
```bash
go_tooling_in_action$ go list
golang/series_3/src/go_tooling_in_action

```

Show package name
```bash
go_tooling_in_action$ go list -f '{{ .Name }}'
main
```

Show package name and pkg level doc: `go doc` will also show you the pkg level doc
```bash
go_tooling_in_action$ go list -f '{{ .Name }}: {{ .Doc }}'
main: package related documentation which you can see on `$ go list -f '{{ .Doc }}'`
```

Show other pkg doc:
```bash
go_tooling_in_action$ go list -f '{{ .Doc }}' fmt
Package fmt implements formatted I/O with functions analogous to C's printf and scanf.

OR

go_tooling_in_action$ go doc fmt
package fmt // import "fmt"

Package fmt implements formatted I/O with functions analogous to C's printf
and scanf. The format 'verbs' are derived from C's but are simpler.

...
```

Get doc of specific func of a pkg:
```bash
go_tooling_in_action$ go doc fmt Println
package fmt // import "fmt"

func Println(a ...interface{}) (n int, err error)
    Println formats using the default formats for its operands and writes to
    standard output. Spaces are always added between operands and a newline is
    appended. It returns the number of bytes written and any write error
    encountered.
```

Get complete doc on http local server:
```bash
go_tooling_in_action$ go get -v  golang.org/x/tools/cmd/godoc
go_tooling_in_action$ godoc -http :8000

# now visit :- http://localhost:8000/
```

List all imports within this `main` pkg:

```bash
$ go list -f '{{ .Imports }}'
[fmt]
```

List all packages on which `fmt` package depends on:
```bash
go_tooling_in_action$ go list -f '{{ join .Imports "\n" }}' fmt
errors
internal/fmtsort
io
math
os
reflect
strconv
sync
unicode/utf8
```

-----

Running http server:
```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}


func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Hello, world !!!")
}

```            
Here when we run it, the is not running but not throwing any exceptions too, because go doesn't throws exception instead we have to explicitly get the error and handle it.
```bash
go_tooling_in_action$ go run main.go
```

But how to know at which line caused the error, so that we can get it and handle it explicitly.

For that we can use tool [`errcheck`](https://github.com/kisielk/errcheck):
```bash
go_tooling_in_action$ go get -u github.com/kisielk/errcheck
go: finding github.com/kisielk/errcheck v1.2.0
go: downloading github.com/kisielk/errcheck v1.2.0
go: extracting github.com/kisielk/errcheck v1.2.0
go: finding golang.org/x/tools latest

```

After running it, it shows the error occurred at last execution:
```bash
go_tooling_in_action$ errcheck 
main.go:11:21:	http.ListenAndServe(":8000", nil)
```
So we can see that we need to handle error at line no. `11`

```go
	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		log.Fatal(err)
	}
```

And it works fine.