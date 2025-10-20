package main

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"sync"
	"time"
)

// 带参数的任务函数
func taskWithParams(i int) {
	fmt.Printf("Task %d started\n", i)
	time.Sleep(200 * time.Millisecond)
	fmt.Printf("Task %d finished\n", i)
}

func main() {
	pool, _ := ants.NewPool(5)
	defer pool.Release()

	var wg sync.WaitGroup

	for i := 0; i < 20; i++ {
		wg.Add(1)
		// 使用闭包捕获参数
		_ = pool.Submit(func() {
			defer wg.Done()
			taskWithParams(i)
		})
	}

	wg.Wait()
}
