package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("===== Go context.WithDeadline 深入解析 =====")
	fmt.Println("按回车键继续...")
	fmt.Scanln()

	demoBasicDeadline()

	fmt.Println("\n按回车键继续...")
	fmt.Scanln()
	demoAdvancedFeatures()

	fmt.Println("\n按回车键继续...")
	fmt.Scanln()
	demoHttpClientTimeout()

	fmt.Println("\n按回车键继续...")
	fmt.Scanln()
	demoTaskScheduler()

	fmt.Println("\n按回车键退出程序...")
	fmt.Scanln()
}

// ====== 基本使用演示 ======
func demoBasicDeadline() {
	fmt.Println("\n=== 1. context.WithDeadline 基本使用 ===")

	// 1.1 创建特定绝对时间的截止上下文
	deadline := time.Now().Add(2 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel() // 重要：释放资源

	fmt.Printf("✅ 创建截止时间上下文: %v\n", deadline.Format("15:04:05.000"))

	// 检查截止时间
	if dl, ok := ctx.Deadline(); ok {
		fmt.Printf("上下文截止时间: %v\n", dl.Format("15:04:05.000"))
	} else {
		fmt.Println("上下文未设置截止时间")
	}

	// 模拟操作
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("❌ 操作完成 (但已超时)")
	case <-ctx.Done():
		fmt.Printf("✅ 上下文取消: %v\n", ctx.Err())
	}

	// 1.2 设置过去的时间
	fmt.Println("\n设置过去时间为截止时间:")
	pastDeadline := time.Now().Add(-1 * time.Hour)
	pastCtx, pastCancel := context.WithDeadline(context.Background(), pastDeadline)
	defer pastCancel()

	fmt.Printf("设置截止时间: %v\n", pastDeadline.Format("15:04:05.000"))

	select {
	case <-time.After(100 * time.Millisecond):
		fmt.Println("❌ 不应等待")
	case <-pastCtx.Done():
		if errors.Is(pastCtx.Err(), context.DeadlineExceeded) {
			fmt.Println("✅ 立即超时: 设置过去时间会立即触发超时")
		}
	}
}

// ====== 高级功能演示 ======
func demoAdvancedFeatures() {
	fmt.Println("\n=== 2. context.WithDeadline 高级特性 ===")

	// 2.1 嵌套上下文截止时间传播
	/*
		子上下文可以设置比父上下文更早的截止时间
		子上下文到期不会自动取消父上下文
		父上下文到期会自动取消所有子上下文
	*/
	fmt.Println("\n2.1 嵌套上下文截止时间传播:")
	parentDeadline := time.Now().Add(3 * time.Second)
	parentCtx, parentCancel := context.WithDeadline(context.Background(), parentDeadline)
	defer parentCancel()

	fmt.Printf("父上下文截止时间: %v\n", parentDeadline.Format("15:04:05.000"))

	childDeadline := time.Now().Add(5 * time.Second)
	childCtx, childCancel := context.WithDeadline(parentCtx, childDeadline)
	defer childCancel()

	fmt.Printf("尝试设置子上下文截止时间: %v\n", childDeadline.Format("15:04:05.000"))

	if dl, ok := childCtx.Deadline(); ok {
		fmt.Printf("实际子上下文截止时间: %v (取更早的父截止时间)\n", dl.Format("15:04:05.000"))
	} else {
		fmt.Println("子上下文未设置截止时间")
	}

	start := time.Now()
	<-childCtx.Done()
	fmt.Printf("子上下文取消类型: %v (耗时: %v)\n", childCtx.Err(), time.Since(start))

	// 2.2 剩余时间检查
	fmt.Println("\n2.2 剩余时间检查:")
	remainingCtx, remainingCancel := context.WithDeadline(context.Background(),
		time.Now().Add(2*time.Second))
	defer remainingCancel()

	// 间隔检查剩余时间
	for i := 0; i < 3; i++ {
		if dl, ok := remainingCtx.Deadline(); ok {
			remaining := time.Until(dl)
			fmt.Printf("剩余时间: %v\n", remaining.Round(time.Millisecond))
		}
		time.Sleep(600 * time.Millisecond)
	}

	// 2.3 手动取消先于截止时间
	fmt.Println("\n2.3 手动取消先于截止时间:")
	manualCtx, manualCancel := context.WithDeadline(context.Background(),
		time.Now().Add(2*time.Second))

	go func() {
		time.Sleep(800 * time.Millisecond)
		fmt.Println("⏱️ 提前手动取消...")
		manualCancel()
	}()

	select {
	case <-manualCtx.Done():
		if errors.Is(manualCtx.Err(), context.Canceled) {
			fmt.Println("✅ 手动取消成功")
		}
	case <-time.After(1 * time.Second):
		fmt.Println("❌ 应被手动取消")
	}
}

// ====== HTTP客户端超时控制 ======
func demoHttpClientTimeout() {
	fmt.Println("\n=== 3. HTTP客户端超时控制 (WithDeadline) ===")

	// 创建测试服务器
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if delay := r.URL.Query().Get("delay"); delay != "" {
			d, _ := time.ParseDuration(delay)
			time.Sleep(d)
		}
		w.Write([]byte("服务器响应"))
	}))
	defer server.Close()

	// 3.1 正常请求
	fmt.Println("\n3.1 正常请求 (在截止时间内完成)")
	deadline := time.Now().Add(800 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	req, _ := http.NewRequestWithContext(ctx, "GET", server.URL+"?delay=500ms", nil)

	start := time.Now()
	resp, err := http.DefaultClient.Do(req)

	if err == nil {
		fmt.Printf("✅ 请求成功: 状态码 %d (耗时: %v)\n", resp.StatusCode, time.Since(start))
		resp.Body.Close()
	} else {
		fmt.Printf("❌ 请求失败: %v\n", err)
	}

	// 3.2 超时请求
	fmt.Println("\n3.2 超时请求 (在截止时间内无法完成)")
	shortDeadline := time.Now().Add(300 * time.Millisecond)
	shortCtx, shortCancel := context.WithDeadline(context.Background(), shortDeadline)
	defer shortCancel()

	req, _ = http.NewRequestWithContext(shortCtx, "GET", server.URL+"?delay=500ms", nil)

	start = time.Now()
	_, err = http.DefaultClient.Do(req)

	if errors.Is(err, context.DeadlineExceeded) {
		fmt.Printf("✅ 预期超时: %v (耗时: %v)\n", err, time.Since(start))
	} else if err != nil {
		fmt.Printf("❌ 其他错误: %v\n", err)
	} else {
		fmt.Println("❌ 不应成功但成功了")
	}
}

// ====== 任务调度器应用 ======
func demoTaskScheduler() {
	fmt.Println("\n=== 4. 任务调度器应用 (WithDeadline) ===")

	type Task struct {
		Name     string
		Deadline time.Time
	}

	// 简单的任务执行器
	executeTask := func(ctx context.Context, task Task) {
		taskCtx, cancel := context.WithDeadline(ctx, task.Deadline)
		defer cancel()

		fmt.Printf("\n🟠 启动任务: %-15s 截止时间: %v\n",
			task.Name, task.Deadline.Format("15:04:05.000"))

		// 模拟任务执行时间
		/*
			rand.Intn(1500)  // 产生 [0, 1500) 范围内的随机整数
			rand.Intn(1500) + 500  // 将范围偏移到 [500, 2000)
			time.Duration(...) * time.Millisecond  // 将整数转换为毫秒级时间间隔
			即一个在 500 毫秒到 2 秒之间的随机时间值
		*/
		processingTime := time.Duration(rand.Intn(1500)+500) * time.Millisecond

		// 真实任务启动（在实际中，这可能是数据库查询或API调用）
		resultCh := make(chan string, 1)
		go func() {
			// 模拟实际工作耗时
			time.Sleep(processingTime)
			resultCh <- "任务结果"
		}()

		select {
		case res := <-resultCh:
			if errors.Is(taskCtx.Err(), context.DeadlineExceeded) {
				fmt.Printf("🔴 [超时完成] %-15s 耗时: %v (结果: %s)\n",
					task.Name, processingTime, res)
			} else {
				fmt.Printf("🟢 [按时完成] %-15s 耗时: %v (结果: %s)\n",
					task.Name, processingTime, res)
			}
		case <-taskCtx.Done():
			if errors.Is(taskCtx.Err(), context.DeadlineExceeded) {
				fmt.Printf("🔴 [超时取消] %-15s 耗时: %v\n",
					task.Name, processingTime)
			} else {
				fmt.Printf("🔵 [手动取消] %-15s 原因: %v\n",
					task.Name, taskCtx.Err())
			}
		}
	}

	// 创建一组任务
	now := time.Now()
	tasks := []Task{
		{"短任务", now.Add(800 * time.Millisecond)},
		{"中任务", now.Add(1200 * time.Millisecond)},
		{"长任务", now.Add(1500 * time.Millisecond)},
		{"超长任务", now.Add(2000 * time.Millisecond)},
	}

	fmt.Println("任务列表:")
	for _, task := range tasks {
		fmt.Printf("  - %-15s 截止: %v\n", task.Name, task.Deadline.Format("15:04:05.000"))
	}

	// 创建父上下文 (全局超时)
	parentCtx, parentCancel := context.WithDeadline(context.Background(),
		now.Add(3*time.Second))
	defer parentCancel()

	fmt.Printf("\n全局截止时间: %v\n", now.Add(3*time.Second).Format("15:04:05.000"))

	// 启动所有任务
	var wg sync.WaitGroup
	for _, task := range tasks {
		wg.Add(1)
		go func(t Task) {
			defer wg.Done()
			executeTask(parentCtx, t)
		}(task)
	}

	// 模拟提前取消
	go func() {
		time.Sleep(1 * time.Second)
		if rand.Float32() < 0.5 { // 50% 概率手动取消
			fmt.Println("\n⚠️  手动取消长任务!")
			parentCancel()
		}
	}()

	wg.Wait()
	fmt.Println("\n所有任务执行完成")
}
