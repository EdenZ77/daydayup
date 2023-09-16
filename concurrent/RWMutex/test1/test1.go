package main

import (
	"sync"
	"time"
)

/*
我以计数器为例，来说明一下，如何使用 RWMutex 保护共享资源。计数器的 count++操作是写操作，而获取 count 的值是读操作，这个场景非常适合读写锁，
因为读操作可以并行执行，写操作时只允许一个线程执行，这正是 readers-writers 问题。

在这个例子中，使用 10 个 goroutine 进行读操作，每读取一次，sleep 1 毫秒，同时，
还有一个 gorotine 进行写操作，每一秒写一次，这是一个 1 writer-n reader 的读写场景，而且写操作还不是很频繁（一秒一次）
*/

func main() {
	var counter Counter
	for i := 0; i < 10; i++ { // 10个reader
		go func() {
			for {
				counter.Count() // 计数器读操作
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for { // 一个writer
		counter.Incr() // 计数器写操作
		time.Sleep(time.Second)
	}
}

// Counter 一个线程安全的计数器
type Counter struct {
	mu    sync.RWMutex
	count uint64
}

// Incr 使用写锁保护
func (c *Counter) Incr() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

// Count 使用读锁保护
func (c *Counter) Count() uint64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.count
}
