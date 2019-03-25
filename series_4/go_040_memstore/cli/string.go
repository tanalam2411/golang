package main

import "fmt"

func(m MemStore) String() string {
	s := fmt.Sprintf("%T (isClosed: %v, size: %v, buf: %v)", m, m.isClosed, m.Size(), m.buf)
	return s
}
