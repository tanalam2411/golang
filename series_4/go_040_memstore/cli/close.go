package main


func (m *MemStore) Close() {
	m.isClosed = true
}
