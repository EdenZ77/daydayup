package main

import (
	"fmt"
	"time"
)

/*
在高并发环境或需要频繁创建定时器的场景下，更好的选择是使用 time.NewTimer。使用 time.NewTimer，你可以在使用完毕后停止并释放计时器，从而避免资源的浪费。

在 Go 语言中，time.Timer 结构体提供了一个 Stop 方法，该方法用于停止计时器。
如果调用 Stop 方法成功地停止了定时器，它返回 true。如果定时器已经到期或已经被停止了，则返回 false。

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
		// 当自己预料的事件先发生了则需要停止计时器并释放资源
		/*
			为什么需要这个逻辑
			这个逻辑可以确保在以下两种情况下的正确行为：

			定时器成功停止: 如果定时器尚未触发且成功停止，则定时器的通道 timer.C 中不会有任何值，不需要进一步操作。
			定时器已经触发: 如果定时器已经触发，则通道 timer.C 中可能有一个值。通过 <-timer.C 从通道中读取这个值，可以防止通道阻塞，确保程序的稳定性。
		*/
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

func main1() {
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
func main() {
	// 创建一个定时器，2秒后触发
	timer := time.NewTimer(2 * time.Second)

	go func() {
		// 当定时器到的时候，timer.C通道才有值
		// 如果定时器到了，这个时候调用timer.Stop方法，则返回false，这个时候调用timer.Stop的地方需要将<-timer.C通道清空，
		// 这是为了防止如果没有任何地方调用过<-timer.C，
		<-timer.C
		fmt.Println("Timer expired")
	}()

	time.Sleep(3 * time.Second)
	// 这个时候调用Stop返回false，因为定时已经到了
	if !timer.Stop() {
		// 使用 select 语句避免死锁
		select {
		case <-timer.C:
			fmt.Println("Timer already expired and channel drained")
		default:
			// 通道已经被其他 goroutine drained
		}
	} else {
		fmt.Println("Timer stopped")
	}
}
