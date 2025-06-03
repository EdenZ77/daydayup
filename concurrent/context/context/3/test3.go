package main

import (
	"context"
	"fmt"
	"time"
)

// 模拟长时间运行的任务
func worker(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done(): // 监听取消信号
			fmt.Printf("Worker %d: 收到取消信号，原因: %v\n", id, ctx.Err())
			return
		default:
			fmt.Printf("Worker %d: 正在处理任务...\n", id)
			time.Sleep(500 * time.Millisecond) // 模拟工作
		}
	}
}

func main() {
	// 创建可取消的Context
	ctx, cancel := context.WithCancel(context.Background())

	// 启动3个worker
	for i := 1; i <= 3; i++ {
		go worker(ctx, i)
	}

	// 模拟运行一段时间后取消
	time.Sleep(2 * time.Second)
	fmt.Println("\n主程序: 发送取消信号...")
	cancel() // 触发取消操作

	// 给worker一些时间响应
	time.Sleep(500 * time.Millisecond)
	fmt.Println("主程序: 退出")
}
