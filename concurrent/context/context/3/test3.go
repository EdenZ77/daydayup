package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go goTest(ctx)
	//time.Sleep(4 * time.Second)
	//for range time.Tick(time.Second) {
	//	select {
	//	case <-ctx.Done():
	//		fmt.Println("收到来自子节点的超时信号")
	//		break
	//	default:
	//		fmt.Println("default")
	//	}
	//}

	//cancel()
	time.Sleep(12 * time.Second)
	fmt.Println("main over ====")
}

func goTest(ctx context.Context) {
	withCancel, cancelFunc := context.WithCancel(ctx)
	// 这里调用了cancel，导致go Speak的ctx被cancel
	defer cancelFunc()
	go Speak(withCancel)
	fmt.Println("goTest goroutine over =========")
}

func Speak(ctx context.Context) {
	fmt.Println("Speak =======")
	timeout, cancelFunc := context.WithTimeout(ctx, 7*time.Second)
	defer cancelFunc()
	//time.Sleep(3 * time.Second)
	//cancelFunc()
	for range time.Tick(time.Second) {
		select {
		case <-timeout.Done():
			fmt.Println("我要闭嘴了")
			return
		default:
			fmt.Println("balabalabalabala")
		}
	}
}
