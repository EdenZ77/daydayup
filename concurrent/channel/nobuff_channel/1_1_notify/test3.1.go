package main

import (
	"fmt"
	"time"
)

// 无缓冲 channel用法：用作信号传递，下面示例展示1对1通知信号

type signal struct{}

func worker() {
	println("worker is working...")
	time.Sleep(5 * time.Second)
}

func spawn(f func()) <-chan signal {
	c := make(chan signal)
	go func() {
		println("worker start to work...")
		f()
		c <- signal{}
	}()
	return c
}

func main() {
	println("start a worker...")
	c := spawn(worker)
	<-c
	fmt.Println("worker work done!")
}
