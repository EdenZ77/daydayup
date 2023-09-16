package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mutex sync.Mutex
	cond := sync.Cond{L: &mutex}
	condition := false

	go func() {
		time.Sleep(1 * time.Second)
		cond.L.Lock()
		fmt.Println("子goroutine已经锁定...")
		fmt.Println("子goroutine更改条件数值，并发送通知...")
		condition = true //更改数值
		cond.Signal()    //发送通知：一个goroutine
		fmt.Println("子goroutine...继续...")
		time.Sleep(5 * time.Second)
		fmt.Println("子goroutine已经解锁!")
		cond.L.Unlock()
	}()

	cond.L.Lock()
	fmt.Println("main 已经锁定...")
	//if !condition {
	for !condition { // 应使用for循环，被唤醒之后需要再次检查条件是否满足
		fmt.Println("main 即将等待...")
		//1.wait尝试解锁
		//2.等待-->当前的goroutine进入阻塞状态，等待被唤醒: signal(), broadcast()
		//3.一旦被唤醒后，又被锁定
		cond.Wait()
		fmt.Println("main 被唤醒...")
	}

	fmt.Println("main 继续")
	fmt.Println("main 解锁...")
	cond.L.Unlock()

	time.Sleep(10 * time.Second)
}
