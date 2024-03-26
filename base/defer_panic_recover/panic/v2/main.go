package main

import "fmt"

/*
如果F的defer中无recover捕获，则将panic抛到G中，G函数会立刻终止，不会执行G函数内后面的内容，但不会立刻return，而调用G的defer...以此类推
*/
func main() {
	G()
	fmt.Println("main")
}

// 输出：
//F start
//F defer
//G 捕获异常: F a
//G defer
//main

func G() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("G 捕获异常:", err)
		}
		fmt.Println("G defer")
	}()
	F()
	fmt.Println("G 继续执行")
}

func F() {
	fmt.Println("F start")
	defer func() {
		fmt.Println("F defer")
	}()
	panic("F a")
}
