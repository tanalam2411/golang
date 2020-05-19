
##### Go and the Operating System

- Advantages of Go
  - Go does not have a preprocessor. It does high-speed compilation.
  - Go uses `static linking` by default, which means that the binary files produced can be easily transferred to other machines with same OS.
    - Once a Go program is compiled successfully and an executable file is generated, the developer does not need to worry about libraries, dependencies and different library versions anymore.
    
- Preprocessor - An extra step in the compilation process (source code -> Preprocessor -> compiler)

---

- Godoc

- ```$ go get -v  golang.org/x/tools/cmd/godoc```
- `godoc fmt Printf`
- `godoc -http=:8000` or `godoc -http :8000`

---
- Compiling Go code

- `go build <xyz>.go`
  - Will create a single statically linked executable file, in the same location.


```bash
$ ./aSourceFile 
Say hello...

$ file aSourceFile
aSourceFile: ELF 64-bit LSB executable, x86-64, version 1 (SYSV), statically linked, not stripped

$ ls -l aSourceFile
-rwxrwxr-x 1 afour afour 2008801 May 19 16:25 aSourceFile
```

- The build file size is big because it is statically linked, which means that it does not require any external libraries to run.

---

- Bypass unused import error

```go

import (
    _ "os"
)
```

---

- Wrong
```go
func main()
{
}
```

- Right
```go
func main() {
}
```

- Go requires the use of semicolons as statement terminators in many contexts, so adding `{` in its own line will make Go compiler insert a semicolon after `func main();` which will produce error.

---

- Downloading Go Package

- Go modules
```bash
$ go mod init github.com/tanalam2411/packt/mastering_go/go_and_the_OS
go: creating new go.mod: module github.com/tanalam2411/packt/mastering_go/go_and_the_OS

$ go get -v github.com/mactsouk/go/simpleGitHub
go: finding github.com/mactsouk/go latest
go: finding github.com/mactsouk/go/simpleGitHub latest

$ ls -l ~/go/pkg/mod/github.com/mactsouk/go@v0.0.0-20180603081621-6a282087f7bd/
total 12
-r--r--r-- 1 afour afour   98 May 19 16:45 README.md
dr-xr-xr-x 5 afour afour 4096 May 19 16:45 RT
dr-xr-xr-x 2 afour afour 4096 May 19 16:45 simpleGitHub

```

- Delete intermediate files of downloaded Go pkg.
```bash
$ go clean -i -v -x  github.com/mactsouk/go/simpleGitHub
cd /home/afour/go/pkg/mod/github.com/mactsouk/go@v0.0.0-20180603081621-6a282087f7bd/simpleGitHub
rm -f simpleGitHub.test simpleGitHub.test.exe

```

- Delete entire pkg
```bash
$ sudo rm -rf ~/go/pkg/mod/github.com/mactsouk/go@v0.0.0-20180603081621-6a282087f7bd/
```

---
