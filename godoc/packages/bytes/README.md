##### Package bytes

`/usr/local/go/src/bytes`

`import "bytes"`

- Package files - 
```bash
https://github.com/golang/go/tree/master/src/bytes
- buffer.go
- bytes.go
- reader.go
```

---

##### type Buffer
- [ref](https://github.com/golang/go/blob/master/src/bytes/buffer.go)
- [ref](https://golang.org/pkg/bytes/#Buffer)
- A Buffer is a variable-sized buffer of bytes with Read and Write methods.
- The zero value for Buffer is an empty buffer ready to use.

```go
type Buffer struct {
	buf      []byte // contents are the bytes buf[off : len(buf)]
	off      int    // read at &buf[off], write at &buf[len(buf)]
	lastRead readOp // last read operation, so that Unread* can work correctly.
}

```