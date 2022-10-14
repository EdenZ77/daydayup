package main

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/petermattis/goid"
)

// 有一点，我要提醒你一句，尽管拥有者可以多次调用 Lock，但是也必须调用相同次数的 Unlock，这样才能把锁释放掉。
// 这是一个合理的设计，可以保证 Lock 和 Unlock 一一对应。

type RecursiveMutex struct {
	sync.Mutex
	owner int64
	// 这个goroutine重入的次数
	recursive int32
}

func (m *RecursiveMutex) Lock() {
	gid := goid.Get()
	if atomic.LoadInt64(&m.owner) == gid {
		m.recursive++
		return
	}
	m.Mutex.Lock()
	atomic.StoreInt64(&m.owner, gid)
	m.recursive = 1
}

func (m *RecursiveMutex) Unlock() {
	gid := goid.Get()
	if atomic.LoadInt64(&m.owner) != gid {
		panic(fmt.Sprintf("wrong the owner, expected:%d but found:%d", m.owner, gid))
	}
	m.recursive--
	if m.recursive != 0 {
		return
	}
	atomic.StoreInt64(&m.owner, -1)
	m.Mutex.Unlock()
}

func main() {

}
