package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	g, _ := errgroup.WithContext(context.Background())

	// 任务1
	g.Go(func() error {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Task 1 completed")
		return nil
	})

	// 任务2
	g.Go(func() error {
		time.Sleep(200 * time.Millisecond)
		fmt.Println("Task 2 completed")
		return nil
	})

	// 任务3（返回错误）
	g.Go(func() error {
		time.Sleep(50 * time.Millisecond)
		return fmt.Errorf("task 3 failed")
	})

	// 等待所有任务完成
	if err := g.Wait(); err != nil {
		fmt.Printf("Error occurred: %v\n", err)
	}
}
