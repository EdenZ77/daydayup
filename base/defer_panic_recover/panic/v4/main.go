package main

import (
	"fmt"
	"time"
)

/*
recover都是在当前的goroutine里进行捕获的，这就是说，对于创建goroutine的外层函数，如果goroutine内部发生panic并且内部没有用recover，
外层函数是无法用recover来捕获的，这样会造成程序崩溃
*/
func main() {
	G()
	fmt.Println("main")
}

// 输出：
//F defer
//panic: F panic
//
//goroutine 6 [running]:
//main.F()
//D:/workspace/go_project/study/daydayup/base/defer_panic_recover/panic/v4/main.go:36 +0x49
//created by main.G
//D:/workspace/go_project/study/daydayup/base/defer_panic_recover/panic/v4/main.go:26 +0x46
//
//Process finished with the exit code 2

func G() {
	defer func() {
		// goroutine外进行recover
		if err := recover(); err != nil {
			fmt.Println("G 捕获异常:", err)
		}
		fmt.Println("G defer")
	}()
	// 创建goroutine调用F函数
	go F()
	time.Sleep(time.Second)
	fmt.Println("G 继续执行")
}

func F() {
	defer func() {
		fmt.Println("F defer")
	}()
	//goroutine内部抛出panic
	panic("F panic")
}
