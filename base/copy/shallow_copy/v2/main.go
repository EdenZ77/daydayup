package main

import "fmt"

// 函数传参:
// Go 函数的参数传递都是值传递。
// 传递值类型(int, struct, array)：传递一个副本，函数内修改不影响原始值。
// 传递引用类型(slice, map, pointer, etc.)：传递的是该类型的值（即指针或包含指针的结构的副本），函数内可以通过这个副本指针修改原始数据。

func modifyValue(x int) {
	x = 100 // 修改的是副本，不影响外部
}

func modifySlice(s []int) {
	s[0] = 100 // 通过副本指针修改了共享的底层数组
}

func main() {
	a := 10
	modifyValue(a)
	fmt.Println(a) // 10 (未变)

	sl := []int{1, 2, 3}
	modifySlice(sl)
	fmt.Println(sl) // [100 2 3] (已变)
}
