package main

import (
	"fmt"
	"time"
)

// 闭包既可以通过函数参数使用外部函数变量，也可以直接使用，两者有什么区别呢？
func main() {
	// case 1 打印：i = 2 说明闭包内捕获外部函数变量是取的地址,而不是调用闭包时刻的参数值*******************非常重要的结论
	//i := 1
	//
	//go func() {
	//	fmt.Println("before i=", i) // before i= ？
	//	time.Sleep(2000 * time.Millisecond)
	//	fmt.Println("after i =", i) // after i = ？
	//}()
	//time.Sleep(1000 * time.Millisecond)
	//i = 2
	//time.Sleep(5000 * time.Millisecond)

	// case 1.1
	//strArr := []string{"11", "22", "33"}
	//// 函数参数是值传递。对于函数参数是切片，必须深刻理解，切片就是一个结构体
	//go func(ss []string) {
	//	time.Sleep(100 * time.Millisecond)
	//	fmt.Println("arr =", ss) // arr = [xx 22 33]
	//}(strArr)
	//strArr[0] = "xx"
	//strArr = append(strArr, "1212")
	//time.Sleep(1000 * time.Millisecond)

	// case 1.2
	//strArr1 := []string{"11", "22", "33"}
	//// 闭包是引用传递
	//go func() {
	//	time.Sleep(100 * time.Millisecond)
	//	fmt.Println("arr1 =", strArr1) // arr1 = [11 22 33 1212]
	//}()
	//strArr1 = append(strArr1, "1212")
	//time.Sleep(1000 * time.Millisecond)

	//slice := new([]int)
	// case 2 打印：i = 1
	// 所以我们在使用go func的时候最好把可能改变的值通过值传递的方式传入到闭包之中,避免在协程运行的时候参数值改变导致结果不可预期

	// 其实在《go语言第一课》的21节说明了，在Go语言中，函数参数传递采用的是值传递。
	// 所谓“值传递”，就是将实际参数在内存中的表示逐位拷贝（Bitwise Copy）到形式参数中。对于像整型、数组、结构体这类类型，
	// 它们的内存表示就是它们自身的数据内容，因此当这些类型作为实参类型时，值传递拷贝的就是它们自身，传递的开销也与它们自身的大小成正比。
	// 但是像 string、切片、map 这些类型就不是了，它们的内存表示对应的是它们数据内容的“描述符”。
	// 当这些类型作为实参类型时，值传递拷贝的也是它们数据内容的“描述符”，不包括数据内容本身，所以这些类型传递的开销是固定的，与数据内容大小无关。
	// 这种只拷贝“描述符”，不拷贝实际数据内容的拷贝过程，也被称为“浅拷贝”。
	i := 1
	go func(i int) {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("i =", i) // 输出为：i = 1
	}(i) // 通过匿名函数参数将值传入闭包

	i++
	time.Sleep(1000 * time.Millisecond)
}
