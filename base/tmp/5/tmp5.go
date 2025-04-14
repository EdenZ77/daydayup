package main

import "fmt"

// 定义一个接受函数作为参数的函数
func process(num int, callback func(int)) {
	fmt.Println("Processing number:", num)
	callback(num * 2) // 调用传入的函数
}

func out(result int) {
	fmt.Println("The out result is:", result)
}

func main() {
	// 定义一个匿名函数作为回调
	//doublePrint := func(result int) {
	//	fmt.Println("The result is:", result)
	//}

	process(5, out)

	// 也可以直接传入匿名函数
	process(10, func(x int) {
		fmt.Println("Direct callback with result:", x)
	})
}
