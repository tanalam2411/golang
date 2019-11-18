
**Modules** - Dependency management system

- Module - is a collection of Go packages with `go.mod` file at its root.
- `go.mod` - defines module's module path(import path)


1) Creating a new module:

```bash
modules$ mkdir hello
modules$ touch hello/hello.go
modules$ touch hello/hello_test.go

-----------hello.go----------------
package hello

func Hello() string {
	return "Hello, world."
}

----------------------------------


-----------hello_test.go----------------------------------
package hello

import "testing"


func TestHello(t *testing.T) {
	want := "Hello, world."
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}

--------------------------------------------------------

modules/hello$ go test
PASS
ok      golang/series_5/modules/hello   0.002s

```
Uses complete path `golang/series_5/modules/hello`

Let's make the current directory(hello) the root of a module.

```bash
modules/hello$ go mod init example.com/hello
go: creating new go.mod: module example.com/hello

```

```bash
modules/hello$ go test
PASS
ok      example.com/hello       0.002s

```
Uses `hello` as root directory because of module initialization which makes it as root dir.

- `go.mod` - only appears in the root of the module.

```bash
modules/hello$ cat go.mod 
module example.com/hello

go 1.12

```


2) Adding a dependency - 


Update `hello.go`

```bash
package hello

import "rsc.io/quote"


func Hello() string {
	return quote.Hello()
}
```

```bash
modules/hello$ go test
go: finding rsc.io/quote v1.5.2
go: finding rsc.io/sampler v1.3.0
go: finding golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
go: downloading rsc.io/quote v1.5.2
go: extracting rsc.io/quote v1.5.2
go: downloading rsc.io/sampler v1.3.0
go: extracting rsc.io/sampler v1.3.0
go: downloading golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
go: extracting golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
PASS
ok      example.com/hello       0.003s

```
The go command resolves imports by using the specific dependency module versions listed in go.mod.
- If import package not  provided in `go.mod` it automatically  downloads the package and updates the `go.mod` file using the latest(tagged) version.
- It also downloaded sub dependencies (rsc.io/sampler, golang.org/x/text)
- Only direct dependencies are recorded in the go.mod file.
```bash
modules/hello$ cat go.mod 
module example.com/hello

go 1.12

require rsc.io/quote v1.5.2

```

- Running `go test` again will not download the dependencies as it is cached locally (`$GOPATH/pkg/mod` )

- listing current module(main module) and all its dependencies(including sub dep..):
```bash
modules/hello$ go list -m all
example.com/hello
golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
rsc.io/quote v1.5.2
rsc.io/sampler v1.3.0

```

v0.0.0-20170915032832-14c0d48ead0c - is called `pseudo-version` UTC and suffix is the prefix of the commit hash(untagged commit).

- `go.sum` - Cryptographic hashes of the content of specific module versions:
```bash
modules/hello$ cat go.sum 
golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c h1:qgOY6WgZOaTkIIMiVjBQcw93ERBE4m30iBm00nkL0i8=
golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c/go.mod h1:NqM8EUOU14njkJ3fqMW+pc6Ldnwhi/IjpwHt7yyuwOQ=
rsc.io/quote v1.5.2 h1:w5fcysjrx7yqtD/aO+QwRjYZOKnaM9Uh2b40tElTs3Y=
rsc.io/quote v1.5.2/go.mod h1:LzX7hefJvL54yjefDEDHNONDjII0t9xZLPXsUe+TKr0=
rsc.io/sampler v1.3.0 h1:7uVkIFmeBqHfdjD+gZwtXXI+RODJ2Wc4O7MPEh/QiW4=
rsc.io/sampler v1.3.0/go.mod h1:T1hPZKmBbMNahiBKFy5HrXp6adAjACjK9JXDnKaTXpA=

```

- Both `go.mod` and `go.sum` should be checked into version control.


3) Upgrading dependencies - 

- Semantic version tags - X.X.X [Major(Breaker), Minor(Feature), Patch]

```bash
modules/hello$ go get golang.org/x/text
go: finding golang.org/x/text v0.3.2
go: finding golang.org/x/tools v0.0.0-20180917221912-90fa682c2a6e
go: downloading golang.org/x/text v0.3.2
go: extracting golang.org/x/text v0.3.2

```

```bash
modules/hello$ go test
--- FAIL: TestHello (0.00s)
    hello_test.go:9: Hello() = "", want "Hello, world."
FAIL
exit status 1
FAIL    example.com/hello       0.004s

```

// indirect - indicates dependency is not used directly by this module, only indirectly by other module dep...
```bash
modules/hello$ cat go.mod 
module example.com/hello

go 1.12

require (
        golang.org/x/text v0.3.2 // indirect
        rsc.io/quote v1.5.2
)

```


- list all available tagged versions of the module.
```bash
modules/hello$ go list  -m -versions golang.org/x/text
golang.org/x/text v0.1.0 v0.2.0 v0.3.0 v0.3.1 v0.3.2
```

```bash
modules/hello$ go get golang.org/x/text@v0.3.1
go: finding golang.org/x/text v0.3.1
go: downloading golang.org/x/text v0.3.1
go: extracting golang.org/x/text v0.3.1
modules/hello$ go test
--- FAIL: TestHello (0.00s)
    hello_test.go:9: Hello() = "", want "Hello, world."
FAIL
exit status 1
FAIL    example.com/hello       0.004s
modules/hello$ go get golang.org/x/text@v0.3.0
go: finding golang.org/x/text v0.3.0
go: downloading golang.org/x/text v0.3.0
go: extracting golang.org/x/text v0.3.0
modules/hello$ go test
PASS
ok      example.com/hello       0.003s

```

- By default version is `@latest`.


4) Adding a dependency on a new major version:

Edit `hello.go`:

add:
```bash
import quoteV3 "rsc.io/quote/v3"


func Proverb() string {
    return quoteV3.Concurrency()
}
```

Add test:
```bash
import "testing"


func TestProverb(t *testing.T) {
    want := "Concurrency is not parallelism."
    if got := Proverb(); gpt != want {
        t.Errorf("Proverb() = %q, want %q", got, want)
    }
}

```

```bash
/modules/hello$ go test
go: finding rsc.io/quote/v3 v3.1.0
go: downloading rsc.io/quote/v3 v3.1.0
go: extracting rsc.io/quote/v3 v3.1.0
PASS
ok      example.com/hello       0.003s

```


Now our module depends on both `rsc.io/quote` and `rsc.io/quote/v3`:
```bash
modules/hello$ go list -m rsc.io/q...
rsc.io/quote v1.5.2
rsc.io/quote/v3 v3.1.0
```

- Each major version(v1, v2., v3 ...) of a module uses different module path.
- Starting at `v2`, the path must end in the major version.
- `rsc.io/quote/v3` - This is called `semantic import versioning`
- At most one major version of any particular module path can be included. `rsc.io/quote` `rsc.io/quote/v2` `rsc.io/quote/v3` ...


5) Upgrading a dependency to a new major version

Upgrade from `rsc.io/quote` to `rsc.io/quote/v3`

```bash
modules/hello$ go doc rsc.io/quote/v3
package quote // import "rsc.io/quote/v3"

Package quote collects pithy sayings.

func Concurrency() string
func GlassV3() string
func GoV3() string
func HelloV3() string
func OptV3() string
```

We can update from `quote.Hello()` to `quoteV3.HelloV3()`:

```bash
package hello

import (
	"rsc.io/quote/v3"
)

func Hello() string {
	return quote.HelloV3()
}

func Proverb() string {
	return quote.Concurrency()
}
```

```bash
modules/hello$ go test
PASS
ok      example.com/hello       0.003s

```

6) Removing unused dependencies:

So we have removed use of `rsc.io/quote` package but its still mentioned in `go.mod` file.

```bash
modules/hello$ go list -m all
example.com/hello
golang.org/x/text v0.3.0
golang.org/x/tools v0.0.0-20180917221912-90fa682c2a6e
rsc.io/quote v1.5.2
rsc.io/quote/v3 v3.1.0
rsc.io/sampler v1.3.0

```

`go build` or `go test` can tell when something is missing and needs to be added, but not when something can safely be removed.

- `go mod tidy` cleans up these unused dependencies.

```bash
modules/hello$ go mod tidy
modules/hello$ go list -m all
example.com/hello
golang.org/x/text v0.3.0
rsc.io/quote/v3 v3.1.0
rsc.io/sampler v1.3.0

```

```bash
modules/hello$ cat go.mod 
module example.com/hello

go 1.12

require (
        golang.org/x/text v0.3.0 // indirect
        rsc.io/quote/v3 v3.1.0
)
```

```bash
modules/hello$ go test
PASS
ok      example.com/hello       0.003s
```





________________________________________________________________________________

**Publishing Go Modules**

- Tag - Alias of git hash
 - Lightweight Tag - Simple pointer to a git commit hash
 - Annotated Tag - Full object (tagger's name, msg, etc)

________________________________________________________________________________
Refs:
- https://blog.golang.org/using-go-modules
- https://medium.com/rungo/anatomy-of-modules-in-go-c8274d215c16
- https://git-scm.com/book/en/v2/Git-Basics-Tagging
- *** https://github.com/golang/go/wiki/Modules