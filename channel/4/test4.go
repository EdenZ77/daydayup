package main

import (
	"fmt"
	"sync"
	"time"
)

// 带缓冲 channel 用作计数信号量的例子

var active = make(chan struct{}, 3)
var jobs = make(chan int, 10)

func main() {
	// 生产一批任务当到jobs channel中
	go func() {
		for i := 0; i < 8; i++ {
			jobs <- i + 1
		}
		close(jobs)
	}()

	var wg sync.WaitGroup
	for job := range jobs {
		// 每个任务启动一个goroutine
		go func(j int) {
			wg.Add(1)
			// 虽然每个任务都启动了一个goroutine，但是只会有3个goroutine拿到信号，当释放了信号之后其余的goroutine才有机会获得信号并执行任务
			active <- struct{}{}
			fmt.Printf("handle job: %d\n", j)
			time.Sleep(2 * time.Second)
			<-active
			wg.Done()
		}(job)
	}
	wg.Wait()
}
