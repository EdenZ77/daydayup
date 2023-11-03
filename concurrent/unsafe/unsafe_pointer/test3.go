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
	fmt.Println("===============")
	//testSlice()
	slice2arrayWithHack()
}

//var i *int32 = (*int32)(unsafe.Pointer(v))
// 将指针v转成通用指针，再转成int32指针。这里就看到了unsafe.Pointer的作用了，您不能直接将v转成int32类型的指针，那样将会panic。
// 刚才说了v的地址其实就是它的第一个成员的地址，所以这个i就很显然指向了v的成员i，
// 通过给i赋值就相当于给v.i赋值了，但是别忘了i只是个指针，要赋值得解引用。

func testSlice() {
	a := [3]int{2, 3, 4}
	//a := []int{3, 4, 5}
	slice := unsafe.Slice(&a[0], 3)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}

func slice2arrayWithHack() {
	var b = []int{11, 12, 13}
	var a = *(*[3]int)(unsafe.Pointer(&b[0]))
	a[1] += 10
	fmt.Printf("%v\n", b) // [11 12 13]
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))

}
