package main

import (
	"fmt"
	"unsafe"
)

/*
参考资料：https://geektutu.com/post/hpg-struct-alignment.html


*/

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

// 为什么叫unsafe包呢？顾名思义，它可能很危险，应该非常谨慎使用，
// 因为主要是通过计算内存大小、对齐系数、偏移量等信息，然后操作指针。直接操作指针当然很危险，这就是我推测叫这个名字的原因
func main() {
	//var a T
	fmt.Println(unsafe.Sizeof(Args{})) // 占用内存字节大小 16
	fmt.Println(unsafe.Sizeof(Flag{})) // 占用内存字节大小 8

	// 对于 struct 结构体类型的变量 x，计算 x 每一个字段 f 的unsafe.Alignof(x.f), unsafe.Alignof(x) 等于其中的最大值。
	fmt.Println(unsafe.Alignof(Args{}.num1)) // 8
	fmt.Println(unsafe.Alignof(Args{}.num2)) // 8
	fmt.Println(unsafe.Alignof(Args{}))      // 返回一个类型的对齐系数或者对齐倍数 8

	fmt.Println(unsafe.Alignof(Flag{}.num1)) // 2
	fmt.Println(unsafe.Alignof(Flag{}.num2)) // 4
	fmt.Println(unsafe.Alignof(Flag{}))      // 返回一个类型的对齐系数或者对齐倍数 4

	// 计算结构体类型某个字段偏移量
	fmt.Println(unsafe.Offsetof(Flag{}.num2)) // 4

	fmt.Println("==============")
	// unsafe.Sizeof返回值不包括x可能引用的任何内存
	// 例如：如果x是一个切片，Sizeof将返回切片的描述符大小，而不是切片引用的内存大小，所谓的切片描述符大小就是切片在运行时表示的结构体
	/*
		// runtime/slice.go
		type slice struct {
		    array unsafe.Pointer // 元素指针，在32位机器占用4个字节，64位机器占用8个字节
		    len   int // 长度，同上
		    cap   int // 容量，同上
		}
		// 所以对于任何类型的切片来说，Sizeof返回值都是24个字节
	*/
	fmt.Println(unsafe.Sizeof(make([]int32, 10, 10)))   // 24
	fmt.Println(unsafe.Alignof(make([]int32, 10, 10)))  // 8，因为slice struct中三个字段都是8字节
	fmt.Println(unsafe.Sizeof(make([]Person, 10, 10)))  // 24
	fmt.Println(unsafe.Alignof(make([]Person, 10, 10))) // 8，这与什么类型的切片无关
	fmt.Println(unsafe.Sizeof([]int32{}))               // 24

	fmt.Println("==============")
	var a [10]int32
	var b [10]struct {
		c int16
		d int8
	}

	fmt.Println("Alignment of array a:", unsafe.Alignof(a))              // 数组a的对齐倍数 4，因为与int32的对齐倍数一致
	fmt.Println("Alignment of array b:", unsafe.Alignof(b))              // 数组b的对齐倍数 2，与构成数组的struct一致
	fmt.Println("Alignment of int32:", unsafe.Alignof(a[0]))             // int32的对齐倍数 4
	fmt.Println("Alignment of struct in array b:", unsafe.Alignof(b[0])) // 数组b中结构体的对齐倍数 2
}

type Person struct {
	Name string
	Age  int
}
