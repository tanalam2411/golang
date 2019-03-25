package main

import "errors"

func (m *MemStore) Read(b []byte) (n int, err error) {
	if m.isClosed {
		return 0, errors.New("Invalid read opertion on closed buffer.")
	}

	if m.readOffset == len(m.buf) {
		return n, errors.New("EOF")
	}

	n = copy(b, m.buf[m.readOffset:])
	m.readOffset += n

	return n, nil
}
