package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// context.WithTimeout

var wg sync.WaitGroup

func worker(ctx context.Context) {
LOOP:
	for {
		fmt.Println("db connecting ...")
		time.Sleep(time.Millisecond * 10) // 假设正常连接数据库耗时10毫秒
		select {
		case <-ctx.Done(): // 50毫秒后自动调用
			fmt.Println("<-ctx.Done()被调用")
			fmt.Println(ctx.Err()) // context deadline exceeded
			break LOOP
		default:
		}
	}
	fmt.Println("worker done!")
	wg.Done()
}

func main() {
	// 设置一个50毫秒的超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	defer cancel() // 通知子goroutine结束
	wg.Add(1)
	go worker(ctx)
	wg.Wait()
	fmt.Println("over")
}
