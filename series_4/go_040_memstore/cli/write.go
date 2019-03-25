package main

import "errors"

func(m *MemStore) Write(b []byte) (n int, err error) {
	if m.isClosed {
		return 0, errors.New("Invalid write opertion on closed buffer.")
	}

	l := len(b)

	temp := make([]byte, len(m.buf)+l)
	n = copy(temp, m.buf)
	n = copy(temp[n:], b)

	m.buf = temp
	return n, nil

}
