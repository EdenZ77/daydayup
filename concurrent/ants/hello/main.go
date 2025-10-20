package main

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"sync"
	"time"
)

func main() {
	// 1. 创建协程池（容量10）
	pool, _ := ants.NewPool(10)
	defer pool.Release() // 使用完毕后释放

	var wg sync.WaitGroup

	// 2. 提交任务
	for i := 0; i < 100; i++ {
		wg.Add(1)
		task := func() {
			defer wg.Done()
			fmt.Printf("Processing task %d\n", i)
			time.Sleep(100 * time.Millisecond)
		}
		_ = pool.Submit(task) // 提交任务到池
	}

	wg.Wait()
	fmt.Printf("Running goroutines: %d\n", pool.Running())
}
