package main

import (
	"fmt"
	"golang.org/x/sync/singleflight"
	"log"
	"sync"
	"time"
)

var sfGroup singleflight.Group

func getDataFromDB(key string) (interface{}, error) {
	log.Printf(">> 模拟数据库查询: %s", key)
	time.Sleep(100 * time.Millisecond)
	return fmt.Sprintf("Data for %s", key), nil
}

func main() {
	var wg sync.WaitGroup

	// 完全独立的调用 - 没有并发竞争
	wg.Add(1)
	go func() {
		defer wg.Done()
		value, _, shared := sfGroup.Do("unique_key", func() (interface{}, error) {
			return getDataFromDB("unique_key")
		})
		log.Printf("独立调用: Value: %v, Shared: %t", value, shared)
	}()
	wg.Wait()
}
