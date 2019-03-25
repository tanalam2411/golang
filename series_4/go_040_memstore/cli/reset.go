package main

import "errors"

func(m *MemStore) Reset() (err error) {
	if m.isClosed {
		return errors.New("Cann't reset closed buffer.")
	}

	m.readOffset = 0
	return nil
}
