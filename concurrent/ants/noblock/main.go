package main

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"time"
)

func main() {
	// 非阻塞模式（队列满时立即返回错误）
	pool, _ := ants.NewPool(5, ants.WithNonblocking(true))
	defer pool.Release()

	for i := 0; i < 10; i++ {
		err := pool.Submit(func() {
			time.Sleep(1 * time.Second)
		})

		if err != nil {
			fmt.Printf("Task %d submit failed: %v\n", i, err)
		}
	}
}
