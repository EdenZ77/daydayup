package main

import (
	"fmt"
	"time"
)

// 无缓冲 channel 用作信号传递的时候，
// 有两种情况，分别是 1 对 1 通知信号和 1 对 n 通知信号。
// 我们先来分析下 1 对 1 通知信号这种情况。

type signal struct{}

func worker() {
	println("worker is working...")
	time.Sleep(time.Second)
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
