package main

import (
	"fmt"
	"unsafe"
)

type Args struct {
	num1 int
	num2 int
}

type Flag struct {
	num1 int16
	num2 int32
}

type T struct {
	b byte

	i int64
	u uint16
}

var name = make([]*Args, 10)

// 为什么叫unsafe包呢？顾名思义，它可能很危险，应该非常谨慎使用，
// 因为主要是通过计算内存大小、对齐系数、偏移量等信息，然后操作指针。直接操作指针当然很危险，这就是我推测叫这个名字的原因
func main() {
	//var a T
	fmt.Println(unsafe.Sizeof(Args{})) // 占用内存字节大小 16
	fmt.Println(unsafe.Sizeof(Flag{})) // 占用内存字节大小 8
	//fmt.Println(unsafe.Sizeof(name)) 后期研究关于切片等类型的内存字节大小问题，先目前先简单研究结构体的内存大小

	// 对于 struct 结构体类型的变量 x，计算 x 每一个字段 f 的unsafe.Alignof(x.f), unsafe.Alignof(x) 等于其中的最大值。
	fmt.Println(unsafe.Alignof(Args{}.num1)) // 8
	fmt.Println(unsafe.Alignof(Args{}.num2)) // 8
	fmt.Println(unsafe.Alignof(Args{}))      // 返回一个类型的对齐系数或者对齐倍数 8

	fmt.Println(unsafe.Alignof(Flag{}.num1)) // 2
	fmt.Println(unsafe.Alignof(Flag{}.num2)) // 4
	fmt.Println(unsafe.Alignof(Flag{}))      // 返回一个类型的对齐系数或者对齐倍数 4

	// 计算结构体类型某个字段偏移量
	fmt.Println(unsafe.Offsetof(Flag{}.num2)) // 4
}
