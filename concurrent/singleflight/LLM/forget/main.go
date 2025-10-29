package main

import (
	"database/sql"
	"errors"
	"fmt"
	"golang.org/x/sync/singleflight"
	"log"
	"sync"
	"time"
)

var sfGroup singleflight.Group

func getDataFromDB(key string) (interface{}, error) {
	log.Printf(">> 模拟数据库查询: %s", key)
	time.Sleep(100 * time.Millisecond) // 模拟耗时
	return fmt.Sprintf("Data for %s", key), nil
}

func main() {
	var wg sync.WaitGroup
	key := "user_123"

	// 启动 5 个并发请求
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			// 使用 singleflight
			value, err, shared := sfGroup.Do(key, func() (interface{}, error) {
				data, err := getDataFromDB(key)
				if errors.Is(err, sql.ErrNoRows) {
					// 如果数据不存在，Forget 这个 key，让下一次调用可以重新查询
					sfGroup.Forget(key)
				}
				return data, err
			})

			if err != nil {
				log.Printf("Goroutine %d: Error: %v", id, err)
				return
			}
			log.Printf("Goroutine %d: Value: %v, Shared: %t", id, value, shared)
		}(i)
	}

	wg.Wait()
	log.Println("main end!")
}
