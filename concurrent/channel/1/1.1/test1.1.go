package main

import (
	"sync"
	"time"
)

/*
演示：使用操作符<-，声明只发送 channel 类型（send-only）和只接收 channel 类型（recv-only）

在这个例子中，我们启动了两个 Goroutine，分别代表生产者（produce）与消费者（consume）。
生产者只能向 channel 中发送数据，我们使用chan<- int作为 produce 函数的参数类型；消费者只能从 channel 中接收数据，我们使用<-chan int作为 consume 函数的参数类型。

在消费者函数 consume 中，我们使用了 for range 循环语句来从 channel 中接收数据，for range 会阻塞在对 channel 的接收操作上，
直到 channel 中有数据可接收或 channel 被关闭循环，才会继续向下执行。channel 被关闭后，for range 循环也就结束了。

*/

func produce(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i + 1
		time.Sleep(time.Second)
	}
	close(ch)
}

func consume(ch <-chan int) {
	for n := range ch {
		println(n)
	}
}

func main() {
	ch := make(chan int, 5)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		produce(ch)
	}()

	go func() {
		defer wg.Done()
		consume(ch)
	}()

	wg.Wait()
}
