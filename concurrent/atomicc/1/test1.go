package main

import (
	"fmt"
	"sync"
)

// 多个goroutine 累加count，导致并发问题
var count int

func add(wg *sync.WaitGroup) {
	defer wg.Done()
	count++
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go add(&wg)
	}
	wg.Wait()
	fmt.Println(count)
}
