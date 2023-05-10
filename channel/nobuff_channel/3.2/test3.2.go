package main

import (
	"fmt"
	"sync"
	"time"
)

// 无缓冲channel 实现1对n的信号通知机制

// 多个业务工作
func worker(i int) {
	fmt.Printf("worker %d: is working...\n", i)
	time.Sleep(1 * time.Second)
	fmt.Printf("worker %d: works done\n", i)
}

type signal struct{}

func spawnGroup(f func(i int), num int, groupSignal <-chan signal) <-chan signal {
	// 此处初始化了一个1：1的信号通知
	c := make(chan signal)
	var wg sync.WaitGroup

	for i := 0; i < num; i++ {
		wg.Add(1)
		go func(i int) {
			<-groupSignal
			fmt.Printf("worker %d: start to work...\n", i)
			f(i)
			wg.Done()
		}(i + 1)
	}
	// 当main 协程打开通知，所有业务goroutine跑完，将发送信号给c，main goroutine得到通知，跑完整个进程。
	go func() {
		wg.Wait()
		c <- signal{}
	}()
	return c
}

func main() {
	fmt.Println("start a group of workers...")
	groupSignal := make(chan signal)
	c := spawnGroup(worker, 5, groupSignal)
	time.Sleep(5 * time.Second)
	fmt.Println("the group of workers start to work...")
	// 一组业务被阻塞在无缓冲channel接受数据的地方，当这里close这个channel的话，那批业务的channel将可以执行接受操作，从而实现1：n的信号通知效果
	close(groupSignal)
	// 其实这个模型就是，main发信号让一组业务goroutine统一开始执行，这组业务goroutine都执行完毕之后再给main发送信号，main就可以接着执行后面的逻辑
	<-c
	fmt.Println("the group of workers work done!")
}
