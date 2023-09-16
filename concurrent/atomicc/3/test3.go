package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// https://juejin.cn/post/6844904053042839560
// Go语言在1.4版本的时候向sync/atomic包中添加了新的类型Value，
// 此类型相当于一个容器，被用来"原子地"存储（Store）和加载任意类型的值

// ^uint(0) ：最大无符号整数
// int(^uint(0) >> 1) ：最大有符号整数
// int(^(^uint(0) >> 1)) ：最小负整数

// 其中runtime_procPin方法可以将一个goroutine死死占用当前使用的P (此处参考Goroutine调度器(一)：P、M、G关系)
// 不允许其他的goroutine抢占，而runtime_procUnpin则是释放方法

func main() {
	// 此处依旧选用简单的数据类型，因为代码量少
	config := atomic.Value{}
	config.Store(22)

	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			defer wg.Done()
			// 在某一个goroutine中修改配置
			if i == 0 {
				//time.Sleep(time.Second)
				config.Store(23)
			}
			// 输出中夹杂22，23
			fmt.Println(config.Load())
		}(i)
	}
	wg.Wait()
}
