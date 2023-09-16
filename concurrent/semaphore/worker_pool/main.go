package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"log"
	"runtime"
	"time"
)

var (
	maxWorkers = runtime.GOMAXPROCS(0)                    // worker数量 8
	sema       = semaphore.NewWeighted(int64(maxWorkers)) //信号量
	task       = make([]int, maxWorkers*4)                // 任务数，是worker的四倍 32
)

func main() {
	ctx := context.Background()

	for i := range task {
		// 如果没有worker可用，会阻塞在这里，直到某个worker被释放
		if err := sema.Acquire(ctx, 1); err != nil {
			log.Printf("获取worker失败: %v", err)
			break
		}

		// 启动worker goroutine，由于sema的容量是8，所以最多只能有8个goroutine同时执行
		// 如果不使用信号量，那么就会启动32个goroutine，这样会导致同时运行的goroutine太多，从而导致内存占用过高
		// 使用信号量，可以控制同时运行的goroutine数量 信号量的容量，就是最大的同时运行goroutine的数量
		// 当这一批goroutine执行完毕后，会释放sema的容量，上面的Acquire就会获取到容量，从而继续执行
		go func(i int) {
			defer sema.Release(1)
			fmt.Println(i + 1)
			time.Sleep(4 * time.Second) // 模拟一个耗时操作
			task[i] = i + 1
		}(i)
	}

	// 请求所有的worker,这样能确保前面的worker都执行完
	if err := sema.Acquire(ctx, int64(maxWorkers)); err != nil {
		log.Printf("获取所有的worker失败: %v", err)
	}

	fmt.Println(task)
}
