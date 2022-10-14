package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 之前我们使用Mutex加锁来解决累加问题，但是感觉太重了，这里我们可以使用atomic就可以了
var count int32

func add(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		if atomic.CompareAndSwapInt32(&count, count, count+1) {
			break
		}
	}
}

// 修改方式2
func add2(wg *sync.WaitGroup) {
	defer wg.Done()
	atomic.AddInt32(&count, 1)
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go add(&wg)
	}
	wg.Wait()
	fmt.Println(count)
}
