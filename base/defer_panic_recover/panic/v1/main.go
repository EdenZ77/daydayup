package main

import "fmt"

/*
F中出现panic时，F函数会立刻终止，不会执行F函数内panic后面的内容，但不会立刻return，而是调用F的defer，如果F的defer中有recover捕获，则F在执行完defer后正常返回，调用函数F的函数G继续正常执行

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
		if err := recover(); err != nil {
			fmt.Println("捕获异常:", err)
		}
		fmt.Println("b")
	}()
	panic("a")
}
