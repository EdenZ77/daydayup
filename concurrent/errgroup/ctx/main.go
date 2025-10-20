package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"time"
)

// 上下文感知任务
func main() {
	g, ctx := errgroup.WithContext(context.Background())

	g.Go(func() error {
		select {
		case <-time.After(2 * time.Second):
			fmt.Println("Long task completed")
			return nil
		case <-ctx.Done():
			fmt.Println("Long task canceled")
			return ctx.Err()
		}
	})

	g.Go(func() error {
		time.Sleep(100 * time.Millisecond)
		return fmt.Errorf("trigger error")
	})

	if err := g.Wait(); err != nil {
		fmt.Println("Group error:", err)
	}
}
