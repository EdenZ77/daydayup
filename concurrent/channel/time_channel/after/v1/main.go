package main

import (
	"fmt"
	"time"
)

/*
在高并发环境或需要频繁创建定时器的场景下，更好的选择是使用 time.NewTimer。使用 time.NewTimer，你可以在使用完毕后停止并释放计时器，从而避免资源的浪费。

在 Go 语言中，time.Timer 结构体提供了一个 Stop 方法，该方法用于停止计时器。
如果成功停止计时器且计时器还没有过期（即定时器的 C 通道中还没有值），Stop 方法将返回 true。
如果调用 Stop 方法时计时器已经过期（即定时器的 C 通道中已经有值了），Stop 方法将返回 false，此时通道中的值需要被接收并丢弃，以便释放底层的资源。
*/

func worker(c <-chan bool) {
	// 创建一个定时器，设置超时时间为30秒
	timer := time.NewTimer(30 * time.Second)

	select {
	case signal := <-c:
		// 如果从通道 c 中接收到消息，则处理消息
		if signal {
			fmt.Println("Received signal from channel, doing some work.")
		}
		// 停止计时器并释放资源
		if !timer.Stop() {
			<-timer.C
		}
	case <-timer.C:
		// 如果计时器超时，则执行超时后的逻辑
		fmt.Println("Timeout! Ending worker.")
		return // 退出函数
	}

	// 如果函数中还有其他逻辑，确保在退出前停止并释放计时器
	// 防止资源泄露
	if !timer.Stop() {
		<-timer.C
	}

	// 其他逻辑...
}

func main() {
	// 创建一个布尔类型的通道
	done := make(chan bool)

	// 启动 worker 协程
	go worker(done)

	// 模拟在一段时间后发送信号通知 worker
	time.Sleep(5 * time.Second) // 等待5秒
	done <- true                // 发送信号

	// 给 worker 一些时间来处理接收的信号
	time.Sleep(2 * time.Second)

	// 结束主程序
	fmt.Println("Main program ending.")
}
