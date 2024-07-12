package main

import (
	"fmt"
	"sync"
	"time"
)

type SharedData struct {
	mu   sync.Mutex
	data int64
}

func (s *SharedData) Modify() {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i := int64(0); i < 10000000; i++ {
		s.data += 1
		if i == 500 {
			time.Sleep(1 * time.Millisecond)
		}
	}
	s.data = 42
}

func (s *SharedData) Read() int64 {
	return s.data
}

func main() {
	shared := &SharedData{data: int64(0)}

	var wg sync.WaitGroup
	wg.Add(2)

	// 协程A：修改数据
	go func() {
		defer wg.Done()
		shared.Modify()
	}()

	// 协程B：读取数据
	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Millisecond) // 确保协程A开始修改
		fmt.Println("Read data: ", shared.Read())
	}()

	wg.Wait()
	fmt.Println("Final data: ", shared.data)
}
