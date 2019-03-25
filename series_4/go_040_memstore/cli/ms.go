package main

import (
	"fmt"
	"io"
)

// MemStore interface
type MemStoreOperation interface {
	Size() int
	IsEmpty() bool
	Write(b []byte) (n int, err error)
	io.Reader
	fmt.Stringer
}


// MemStore struct
type MemStore struct {
	buf []byte
	readOffset int
	isClosed bool
}

func (m MemStore) Size() int {
	return len(m.buf)
}


func (m MemStore) IsEmpty() bool {
	return m.Size() == 0
}
