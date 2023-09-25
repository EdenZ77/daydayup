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

func G() {
	defer func() {
		//goroutine外进行recover
		if err := recover(); err != nil {
			fmt.Println("捕获异常:", err)
		}
		fmt.Println("c")
	}()
	//创建goroutine调用F函数
	go F()
	time.Sleep(time.Second)
}

func F() {
	defer func() {
		fmt.Println("b")
	}()
	//goroutine内部抛出panic
	panic("a")
}
