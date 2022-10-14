package main

import (
	"fmt"
	"time"
)

// 无缓冲 channel用法：用作信号传递，下面示例展示1对1通知信号

type signal struct{}

func worker() {
	println("执行某个业务...")
	time.Sleep(5 * time.Second)
}
func spawn(f func()) <-chan signal {
	c := make(chan signal)
	go func() {
		println("worker start to work...")
		f()
		// 执行完成外部传入的业务之后，传递信号
		c <- signal{}
	}()
	return c
}

func main() {
	println("start a worker...")
	c := spawn(worker)
	// 在任务没有完成的时候，阻塞在这里
	<-c
	fmt.Println("worker work done!")
}
