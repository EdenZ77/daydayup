package main

import "sync"

/*
为什么要讨论这个话题呢？因为 Mutex 经常会和其他非线程安全（对于 Go 来说，我们其实指的是 goroutine 安全）的数据结构一起，组合成一个线程安全的数据结构。
新数据结构的业务逻辑由原来的数据结构提供，而 Mutex 提供了锁的机制，来保证线程安全。

比如队列，我们可以通过 Slice 来实现，但是通过 Slice 实现的队列不是线程安全的，出队（Dequeue）和入队（Enqueue）会有 data race 的问题。
这个时候，Mutex 就要隆重出场了，通过它，我们可以在出队和入队的时候加上锁的保护。
*/

type SliceQueue struct {
	data []interface{}
	mu   sync.Mutex
}

func NewSliceQueue(n int) *SliceQueue {
	return &SliceQueue{
		data: make([]interface{}, 0, n),
	}
}

func (q *SliceQueue) Enqueue(v interface{}) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.data = append(q.data, v)
}

func (q *SliceQueue) Dequeue() interface{} {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.data) == 0 {
		return nil
	}
	v := q.data[0]
	q.data = q.data[1:]
	return v
}

func main() {

}
