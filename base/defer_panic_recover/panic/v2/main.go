package main

import "fmt"

/*
如果F的defer中无recover捕获，则将panic抛到G中，G函数会立刻终止，不会执行G函数内后面的内容，但不会立刻return，而调用G的defer...以此类推
*/
func main() {
	G()
	fmt.Println("main")
}
func G() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("捕获异常:", err)
		}
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
