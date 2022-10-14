package main

import "sync"

// RWMap 一个读写锁保护的线程安全的map
type RWMap struct {
	// 读写锁保护下面的map字段
	sync.RWMutex
	m map[int]int
}

func NewRWMap(n int) *RWMap {
	return &RWMap{
		m: make(map[int]int, n),
	}
}

func (m *RWMap) Get(k int) (int, bool) {
	m.RLock()
	defer m.RUnlock()
	v, existed := m.m[k]
	return v, existed
}

func (m *RWMap) Set(k, v int) {
	m.Lock()
	defer m.Unlock()
	m.m[k] = v
}

func (m *RWMap) Delete(k int) {
	m.Lock()
	defer m.Unlock()
	delete(m.m, k)
}

func (m *RWMap) Len() int {
	m.RLock()
	defer m.RUnlock()
	return len(m.m)
}

func (m *RWMap) Each(f func(k, v int) bool) {
	m.RLock()
	defer m.RUnlock()

	for k, v := range m.m {
		if !f(k, v) {
			return
		}
	}

}

func main() {

}
