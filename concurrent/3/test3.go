package main

import (
	"fmt"
	"sync"
)

// Counter 更多的时候我们是将Mutex放入到需要加锁的struct属性上（可采用嵌入字段的方式，或者 mu sync.Mutex 这种普通属性的方式）
// 如果嵌入的 struct 有多个字段，我们一般会把 Mutex 放在要控制的字段上面，然后使用空格把字段分隔开来。即使你不这样做，代码也可以正常编译，只不过，用这种风格去写的话，逻辑会更清晰，也更易于维护。
// 甚至，你还可以把获取锁、释放锁、计数加一的逻辑封装成一个方法，对外不需要暴露锁等逻辑 见test3.1

type Counter struct {
	sync.Mutex
	Count uint64
}

func main() {
	// 计数器的值
	var counter Counter
	// 使用WaitGroup等待10个goroutine完成
	var wg sync.WaitGroup
	wg.Add(10)

	// 启动10个goroutine
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			// 对变量count执行10万次加1
			for j := 0; j < 100000; j++ {
				counter.Lock()
				counter.Count++
				counter.Unlock()
			}
		}()
	}
	// 等待10个goroutine完成
	wg.Wait()
	fmt.Println(counter.Count)
}
