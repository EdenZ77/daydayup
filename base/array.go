package main

import (
	"fmt"
)

func main() {
	//slice := make([]int, 5, 10) // 长度为5，容量为10
	//slice[2] = 2                // 索引为2的元素赋值为2
	//fmt.Println(slice)

	// 定义一个数组
	array := [5]int{0, 1, 2, 3, 4}

	// 对数组进行切片操作
	slice := array[:]
	fmt.Println(len(slice)) // 5
	fmt.Println(cap(slice)) // 5
}
