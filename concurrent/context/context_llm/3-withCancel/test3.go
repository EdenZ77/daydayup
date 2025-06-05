package main

import (
	"context"
	"fmt"
	"time"
)

// 任务函数
func runTask(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("[%s] 收到取消信号: %v\n", name, ctx.Err())
			return
		default:
			fmt.Printf("[%s] 正在工作...\n", name)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

// 1、取消信号自动向下传播到所有子上下文
// 2、可以单独取消某个子上下文而不影响同级或父级
/*
根上下文 (rootCtx)
├── 子上下文1 (childCtx1)
│   └── 孙子上下文 (grandchildCtx)
└── 子上下文2 (childCtx2)
*/

func main() {
	// 创建根上下文
	rootCtx, rootCancel := context.WithCancel(context.Background())
	defer rootCancel() // 确保根上下文最终被取消

	// 从根上下文创建两个子上下文
	childCtx1, cancelChild1 := context.WithCancel(rootCtx)
	childCtx2, childCancel2 := context.WithCancel(rootCtx)
	defer childCancel2() // 确保子上下文2最终被取消

	// 从 childCtx1 创建孙子上下文
	grandchildCtx, _ := context.WithCancel(childCtx1)

	// 启动任务
	go runTask(rootCtx, "根任务")
	go runTask(childCtx1, "子任务1")
	go runTask(childCtx2, "子任务2")
	go runTask(grandchildCtx, "孙子任务1")

	// 场景1: 取消子上下文1
	time.Sleep(2 * time.Second)
	fmt.Println("\n--- 取消子上下文1 ---")
	cancelChild1()

	// 等待效果显现
	time.Sleep(1 * time.Second)

	// 场景2: 取消根上下文
	time.Sleep(1 * time.Second)
	fmt.Println("\n--- 取消根上下文 ---")
	rootCancel()

	// 等待所有任务响应
	time.Sleep(1 * time.Second)
	fmt.Println("\n所有任务完成")
}
