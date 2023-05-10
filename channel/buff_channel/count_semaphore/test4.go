package main

import (
	"fmt"
	"sync"
	"time"
)

// 带缓冲 channel 用作计数信号量的例子
/*
Go 并发设计的一个惯用法，就是将带缓冲 channel 用作计数信号量（counting semaphore）。带缓冲 channel 中的当前数据个数代表的是，
当前同时处于活动状态（处理业务）的 Goroutine 的数量，而带缓冲 channel 的容量（capacity），就代表了允许同时处于活动状态的 Goroutine 的最大数量。
向带缓冲 channel 的一个发送操作表示获取一个信号量，而从 channel 的一个接收操作则表示释放一个信号量。
*/
var active = make(chan struct{}, 3)
var jobs = make(chan int, 10)

func main() {
	// 生产一批任务当到jobs channel中; 生产任务
	go func() {
		for i := 0; i < 8; i++ {
			jobs <- i + 1
		}
		// 任务生产完毕
		close(jobs)
	}()

	var wg sync.WaitGroup
	// 消费任务，当任务生产完毕，退出for循环
	for job := range jobs {
		// 每个任务启动一个goroutine
		go func(j int) {
			wg.Add(1)
			// 虽然每个任务都启动了一个goroutine，但是只会有3个goroutine拿到信号，当释放了信号之后其余的goroutine才有机会获得信号并执行任务
			active <- struct{}{}
			fmt.Printf("handle job: %d\n", j)
			time.Sleep(2 * time.Second)
			// 处理完任务后，释放信号量
			<-active
			wg.Done()
		}(job)
	}
	wg.Wait()
}
