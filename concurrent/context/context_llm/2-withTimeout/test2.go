package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

// 模拟一个可能耗时的API调用
func callSlowAPI(ctx context.Context, apiName string) (string, error) {
	// 模拟不同的处理时间
	var delay time.Duration
	switch apiName {
	case "user":
		// 1秒=1000毫秒
		delay = 1200 * time.Millisecond
	case "order":
		delay = 800 * time.Millisecond
	case "product":
		delay = 300 * time.Millisecond
	default:
		delay = 2 * time.Second
	}

	// 模拟API处理时间
	select {
	case <-time.After(delay):
		return fmt.Sprintf("%s API 响应 (耗时: %v)", apiName, delay), nil
	case <-ctx.Done():
		return "", fmt.Errorf("%s API 超时取消: %v", apiName, ctx.Err())
	}
}

// 带有超时控制的HTTP请求
func fetchWithTimeout(ctx context.Context, url string) (string, error) {
	// 创建HTTP请求
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return "", err
	}

	// 执行请求
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Do err")
		return "", err
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func main() {
	// 示例1: 基本超时控制
	fmt.Println("=== 示例1: 基本超时控制 ===")
	ctx1, cancel1 := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel1()

	result, err := callSlowAPI(ctx1, "user")
	if err != nil {
		fmt.Println("错误:", err)
	} else {
		fmt.Println("结果:", result)
	}

	// 示例2: 多个并行请求
	fmt.Println("\n=== 示例2: 多个并行请求 ===")
	ctx2, cancel2 := context.WithTimeout(context.Background(), 800*time.Millisecond)
	defer cancel2()

	var wg sync.WaitGroup
	apis := []string{"user", "order", "product"}

	results := make(chan string, len(apis))
	errors := make(chan error, len(apis))

	for _, api := range apis {
		wg.Add(1)
		go func(name string) {
			defer wg.Done()
			res, err1 := callSlowAPI(ctx2, name)
			if err1 != nil {
				errors <- err1
			} else {
				results <- res
			}
		}(api)
	}

	// 等待所有goroutine完成
	go func() {
		wg.Wait()
		close(results)
		close(errors)
	}()

	// 收集结果
	fmt.Println("成功结果:")
	for res := range results {
		fmt.Println("-", res)
	}
	fmt.Println("\n错误信息:")
	for err1 := range errors {
		fmt.Println("-", err1)
	}

	// 示例3: HTTP请求超时
	fmt.Println("\n=== 示例3: HTTP请求超时 ===")
	// 将500毫秒改为5000毫秒，这个接口就可以正常调通啦
	ctx3, cancel3 := context.WithTimeout(context.Background(), 800*time.Millisecond)
	defer cancel3()

	// 注意: 这个URL实际会超时，因为设置的超时时间很短
	url := "https://httpbin.org/delay/1" // 这个URL会延迟1秒响应
	start := time.Now()
	body, err := fetchWithTimeout(ctx3, url)
	if err != nil {
		fmt.Printf("HTTP请求错误 (%v): %v\n", time.Since(start), err)
	} else {
		fmt.Printf("HTTP请求成功 (%v): %d字节\n", time.Since(start), len(body))
	}

	fmt.Println("main out!!!")
}
