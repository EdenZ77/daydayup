package main

import (
	"fmt"
	"sync"
)

// 用无缓冲channel替代锁，之前我们学过用Mutex和atomic来实现累加这个案例，这里使用无缓冲channel来实现。
// 这个核心思想就是利用了无缓冲channel的同步阻塞特性

type counter struct {
	c chan int
	i int
}

func NewCounter() *counter {
	cter := &counter{
		c: make(chan int),
	}
	go func() {
		for {
			cter.i++
			cter.c <- cter.i
		}
	}()
	return cter
}

func (cter *counter) Increase() int {
	return <-cter.c
}

func main() {
	cter := NewCounter()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			v := cter.Increase()
			fmt.Printf("goroutine-%d: current counter value is %d\n", i, v)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
