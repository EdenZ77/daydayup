package main

import "fmt"

/*
如果一直没有recover，抛出的panic到当前goroutine最上层函数时，程序直接异常终止
*/

func main() {
	G()
	fmt.Println("main")
}

func G() {
	defer func() {
		fmt.Println("c")
	}()
	F()
	fmt.Println("继续执行")
}

func F() {
	defer func() {
		fmt.Println("b")
	}()
	panic("a")
}
