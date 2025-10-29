package main

import (
	"fmt"
	"golang.org/x/sync/singleflight"
	"log"
	"time"
)

var sfGroup singleflight.Group

func getDataFromDB(key string) (interface{}, error) {
	log.Printf(">> 模拟数据库查询: %s", key)
	time.Sleep(100 * time.Millisecond) // 模拟耗时
	return fmt.Sprintf("Data for %s", key), nil
}

func main() {
	key := "user_123"
	ch := sfGroup.DoChan(key, func() (interface{}, error) {
		return getDataFromDB(key)
	})

	// 设置超时
	timeout := time.After(200 * time.Millisecond)
	select {
	case result := <-ch:
		if result.Err != nil {
			// 处理错误
			fmt.Printf("result err: %+v\n", result.Err)
			return
		}
		fmt.Printf("Success: %v, Shared: %t\n", result.Val, result.Shared)
	case <-timeout:
		fmt.Println("Request timeout!")
	}
}
