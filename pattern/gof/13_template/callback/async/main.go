package main

import (
	"fmt"
	"time"
)

// Callback 定义一个回调函数类型
type Callback func(result string)

// 异步操作函数，接受回调函数作为参数
func doAsyncTask(callback Callback) {
	fmt.Println("回调函数")
	// 模拟耗时操作
	go func() {
		fmt.Println("异步任务开始...")
		time.Sleep(2 * time.Second) // 模拟异步工作
		callback("任务完成")            // 调用回调函数
	}()
}

func main() {
	// 定义一个回调函数
	myCallback := func(result string) {
		fmt.Println(result)
	}

	// 调用异步操作函数，传入回调函数
	doAsyncTask(myCallback)

	// 等待异步任务和回调完成
	time.Sleep(4 * time.Second)
	fmt.Println("程序结束")
}
