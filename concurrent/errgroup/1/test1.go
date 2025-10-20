package main

import (
	"fmt"
	"time"
)

func main() {
	f1()
}

func f1() {
	// 无法在 main 函数中 recover 另一个goroutine中引发的 panic。
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("main: recover panic:%v\n", e)
		} else {
			fmt.Println("no recover")
		}
	}()
	// 开启一个goroutine执行任务
	go func() {
		// 在同一个goroutine中可recover panic，recover可以不让整个程序直接崩溃
		//defer func() {
		//	if e := recover(); e != nil {
		//		fmt.Printf("children: recover panic:%v\n", e)
		//	}
		//}()

		fmt.Println("in goroutine....")
		// 只能触发当前goroutine中的defer, 只要有任一一个goroutine发送panic，整个程序直接退出
		panic("panic in goroutine")

	}()

	time.Sleep(time.Second * 5)
	fmt.Println("exit")
}
