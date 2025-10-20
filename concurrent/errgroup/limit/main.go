package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"time"
)

func main() {
	g, _ := errgroup.WithContext(context.Background())
	g.SetLimit(2) // 最大并发数为2

	for i := 0; i < 500; i++ {
		i := i // 闭包捕获
		g.Go(func() error {
			fmt.Printf("Task %d started\n", i)
			time.Sleep(time.Duration(100*i) * time.Millisecond)
			fmt.Printf("Task %d completed\n", i)
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		fmt.Println("Error:", err)
	}
}
