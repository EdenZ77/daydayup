package main

import (
	"fmt"
	"time"
)

/*
注意：每次调用 time.After 都会创建一个新的计时器，并且它们在到期前不会被垃圾回收。如果计时器的数量非常多，这可能会导致内存消耗增加。
在高并发环境或需要频繁创建定时器的场景下，更好的选择是使用 time.NewTimer。使用 time.NewTimer，你可以在使用完毕后停止并释放计时器，从而避免资源的浪费。
代码见v1
*/

// select 实现超时机制
func main() {
	ch := make(chan int)
	quit := make(chan bool)

	go func() {
		for {
			select {
			//如果有数据，下面打印。但是有可能ch一直没数据
			case num := <-ch:
				fmt.Println("num = ", num)
			//上面的ch如果一直没数据会阻塞，那么select也会检测其他case条件，发现所有case都不满足，且时间超过3秒钟则触发该case
			case <-time.After(3 * time.Second):
				fmt.Println("超时")
				quit <- true //写入
			}
		}
	}()

	for i := 0; i < 5; i++ {
		ch <- i
		time.Sleep(time.Second)
	}

	// 使用了无缓冲channel的1：1信号通知机制
	<-quit
	fmt.Println("程序结束")
}
