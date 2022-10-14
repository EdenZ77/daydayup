package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

// 如何处理test1中出现的多个goroutine累加导致的并发问题呢？
// 最最常见的方式就是——加锁，这10个goroutine必须使用同一把锁
// 但是这种方式似乎有点太重了，其实我们可以通过atomic的方式来实现，可以看当前目录下面的atomicc包

func main() {
	var mu sync.Mutex
	var count = 0

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				mu.Lock()
				count++
				mu.Unlock()
			}
			fmt.Println(GoID())
		}()
	}
	wg.Wait()
	fmt.Println(count, GoID())
}

// GoID 获取每个goroutine 的id
func GoID() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	// 得到id字符串
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %+v", err))
	}
	return id
}
