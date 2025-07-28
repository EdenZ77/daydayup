package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Config struct {
	APIEndpoint string
	Timeout     time.Duration
	MaxRetries  int
}

// 如果不适用atomic更新配置，则读取线程可能读到部分更新的值

func main() {
	var config atomic.Value

	// 初始配置
	initial := Config{
		APIEndpoint: "https://api.example.com",
		Timeout:     5 * time.Second,
		MaxRetries:  3,
	}
	config.Store(initial)

	// 模拟配置更新
	go func() {
		for {
			time.Sleep(10 * time.Second)
			newCfg := Config{
				APIEndpoint: "https://new-api.example.com",
				Timeout:     10 * time.Second,
				MaxRetries:  5,
			}
			config.Store(newCfg)
			fmt.Println("配置已更新")
		}
	}()

	// 工作协程使用配置
	for i := 0; i < 5; i++ {
		go func(id int) {
			for {
				cfg := config.Load().(Config)
				fmt.Printf("Worker %d 使用配置: %s, %v, %d\n",
					id, cfg.APIEndpoint, cfg.Timeout, cfg.MaxRetries)
				time.Sleep(2 * time.Second)
			}
		}(i)
	}

	// 保持主程序运行
	select {}
}
