package main

import (
	"fmt"
	"hello/concurrent/unsafe/unsafe_pointer/p"
	"unsafe"
)

// 意图很明显，我是想通过unsafe包来实现对V的成员i和j赋值，然后通过PutI()和PutJ()来打印观察输出结果。

// 当然会有些限制，比如需要知道结构体V的成员布局，要修改的成员大小以及成员的偏移量。
// 我们的核心思想就是：结构体的成员在内存中的分配是一段连续的内存，结构体中第一个成员的地址就是这个结构体的地址，
// 您也可以认为是相对于这个结构体偏移了0。相同的，这个结构体中的任一成员都可以相对于这个结构体的偏移来计算出它在内存中的绝对地址。

func main() {
	var v *p.V = new(p.V)
	var i *int32 = (*int32)(unsafe.Pointer(v))
	*i = int32(98)
	//var j *int64 = (*int64)(unsafe.Pointer(uintptr(unsafe.Pointer(v)) + uintptr(unsafe.Sizeof(int32(0))))) // 应考虑内存对齐
	//var j *int64 = (*int64)(unsafe.Pointer(uintptr(unsafe.Pointer(v)) + uintptr(unsafe.Sizeof(int64(0))))) // 正确，应该计算j偏移量，但是j不可见，无法引用
	//var j *int64 = (*int64)(unsafe.Pointer(uintptr(unsafe.Pointer(v)) + uintptr(8))) // 正确
	var j *int64 = (*int64)(unsafe.Add(unsafe.Pointer(v), 8)) // 正确, 调用了1.17新增加的unsafe.Add方法
	*j = int64(763)
	v.PutI()
	v.PutJ()
	fmt.Println("===============testSlice")
	testSlice()
	fmt.Println("===============testSliceData")
	testSliceData()
	fmt.Println("===============testString")
	testString()
	fmt.Println("===============testStringData")
	testStringData()
}

//var i *int32 = (*int32)(unsafe.Pointer(v))
// 将指针v转成通用指针，再转成int32指针。这里就看到了unsafe.Pointer的作用了，您不能直接将v转成int32类型的指针，那样将会panic。
// 刚才说了v的地址其实就是它的第一个成员的地址，所以这个i就很显然指向了v的成员i，
// 通过给i赋值就相当于给v.i赋值了，但是别忘了i只是个指针，要赋值得解引用。

/*
func Slice(ptr *ArbitraryType, len IntegerType) []ArbitraryType
这个函数的主要作用是基于一个指针 ptr 和一个长度 len 创建一个切片。以下是它的详细解释：
ptr: 指向底层数组的起始地址。
len: 切片的长度，必须是一个整数类型或者未命名的常量。

效果等效于：(*[len]ArbitraryType)(unsafe.Pointer(ptr))[:]
这段代码的含义是：
将 ptr 转换为一个指向 [len]ArbitraryType 数组的指针。
使用切片操作符 [:] 将这个数组转换为一个切片。
特殊情况： 如果 ptr 是 nil 并且 len 是零，Slice 函数返回 nil。这种行为在某些情况下是有用的，例如当你想要创建一个空切片而无须分配内存时。
*/
// 从内存指针创建切片（零拷贝）
func testSlice() {
	a := [3]int{2, 3, 4}
	slice := unsafe.Slice(&a[0], 2)
	fmt.Println(slice)
	fmt.Println(len(slice)) // 2
	fmt.Println(cap(slice)) // 2
}

/*
unsafe.Slice(ptr, len) 和常规的切片操作 arr[:] 虽然都创建切片，但在本质、安全性、使用场景上有根本区别：
unsafe针对原始内存指针，对于类型安全没有编译时检查，对于边界检查没有运行时检查
*/

// SliceData(slice []ArbitraryType) *ArbitraryType
// 获取切片底层数组的指针
func testSliceData() {
	// 创建切片
	s := []string{"Go", "Rust", "C++"}

	// 获取底层数组指针
	ptr := unsafe.SliceData(s)

	// 通过指针访问元素
	first := *ptr
	second := *(*string)(unsafe.Add(unsafe.Pointer(ptr), unsafe.Sizeof("")))

	fmt.Println(first, second) // Go Rust

	// 修改指针影响原切片
	*ptr = "Golang"
	fmt.Println(s) // [Golang Rust C++]
}

// String(ptr *byte, len IntegerType) string
// 从字节指针创建字符串（零拷贝）
func testString() {
	// 字节数组
	data := []byte{'H', 'e', 'l', 'l', 'o'}

	// 获取起始指针
	ptr := unsafe.SliceData(data)

	// 创建字符串（共享内存）
	str := unsafe.String(ptr, len(data))

	fmt.Println(str) // Hello

	// 修改原始数据会影响字符串
	data[0] = 'h'
	fmt.Println(str) // hello
}

// StringData(str string) *byte
// 获取字符串底层字节数组的指针
func testStringData() {
	s := "Hello, 世界"

	// 获取字节指针
	ptr := unsafe.StringData(s)

	// 计算长度
	lens := len(s)

	// 重新构造切片（零拷贝）
	bytes := unsafe.Slice(ptr, lens)

	fmt.Printf("%q\n", bytes) // "Hello, 世界"

	// 注意：不能直接修改（字符串不可变）
	// *ptr = 'h' // 运行时错误：写入只读内存
}
