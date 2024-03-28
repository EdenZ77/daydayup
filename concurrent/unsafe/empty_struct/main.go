package main

import (
	"fmt"
	"unsafe"
)

// 编译器会优化它们的存储，使得所有这些为0大小的变量共享同一内存地址
// 然而，Go语言规范并不保证这一行为，因此在不同的Go实现或不同的编译条件下，结果也可能不同。这种优化主要是出于性能和内存利用的考虑。

func main() {
	// 空结构体实例
	var s1 struct{}
	var s2 struct{}

	// 空数组实例
	var a1 [0]int
	var a2 [0]int

	// 显示空结构体和空数组的大小
	fmt.Println("Size of empty struct:", unsafe.Sizeof(s1)) // 0
	fmt.Println("Size of empty array:", unsafe.Sizeof(a1))  // 0

	// 显示空结构体和空数组的地址
	fmt.Printf("Address of s1: %p\n", &s1) // 0xd0a5e0
	fmt.Printf("Address of s2: %p\n", &s2) // 0xd0a5e0
	fmt.Printf("Address of a1: %p\n", &a1) // 0xd0a5e0
	fmt.Printf("Address of a2: %p\n", &a2) // 0xd0a5e0

	// 检查两个空结构体是否有相同的地址
	if &s1 == &s2 {
		fmt.Println("s1 and s2 point to the same address.")
	} else {
		fmt.Println("s1 and s2 point to different addresses.")
	}

	// 检查两个空数组是否有相同的地址
	if &a1 == &a2 {
		fmt.Println("a1 and a2 point to the same address.")
	} else {
		fmt.Println("a1 and a2 point to different addresses.")
	}
}
