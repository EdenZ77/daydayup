package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	interfaceTest()
}

/*
IsValid()判断 Value 本身是否持有数据。
在调用 IsNil() 或 IsZero() 之前，必须检查 IsValid()。无效的 Value 调用它们会导致 panic。
*/
func isValidTest() {
	// 情况 1: 有效的值 (非nil)
	var x int = 42
	vx := reflect.ValueOf(x)
	fmt.Println("vx.IsValid():", vx.IsValid()) // true: vx 持有整数 42

	// 情况 2: 有效的值 (nil指针)
	var p *int = nil
	vp := reflect.ValueOf(p)
	fmt.Println("vp.IsValid():", vp.IsValid()) // true: vp 持有 *int 类型的 nil 值

	// 情况 3: 无效的值 (ValueOf(nil))
	vnil := reflect.ValueOf(nil)
	fmt.Println("vnil.IsValid():", vnil.IsValid()) // false: 这是通过 ValueOf(nil) 得到的无效 Value

	// 情况 4: 从无效操作获取的值 (尝试获取不存在的结构体字段)
	type S struct{ F int }
	s := S{F: 1}
	vs := reflect.ValueOf(s)
	vBadField := vs.FieldByName("NonExistent")
	fmt.Println("vBadField.IsValid():", vBadField.IsValid()) // false: 字段不存在

}

/*
IsNil()只检查通道变量本身是否为 nil
已初始化或已关闭的通道都不为 nil
*/
func chanIsNilTest() {
	// nil 通道
	var nilChan chan int
	vNilChan := reflect.ValueOf(nilChan)
	fmt.Println("nilChan is nil:", vNilChan.IsNil()) // true

	// 已初始化的通道
	realChan := make(chan int)
	vRealChan := reflect.ValueOf(realChan)
	fmt.Println("realChan is nil:", vRealChan.IsNil()) // false

	// 已关闭的通道
	closedChan := make(chan struct{})
	close(closedChan)
	vClosedChan := reflect.ValueOf(closedChan)
	fmt.Println("closedChan is nil:", vClosedChan.IsNil()) // false
}

type MyStruct struct{}

func (m *MyStruct) Method() {}

func funcTest() {
	// nil 函数
	var nilFunc func()
	vNilFunc := reflect.ValueOf(nilFunc)
	fmt.Println("nilFunc is nil:", vNilFunc.IsNil()) // true

	// 普通函数
	normalFunc := func() { fmt.Println("Hello") }
	vNormalFunc := reflect.ValueOf(normalFunc)
	fmt.Println("normalFunc is nil:", vNormalFunc.IsNil()) // false

	// 方法值
	var m *MyStruct
	methodValue := m.Method
	vMethod := reflect.ValueOf(methodValue)
	fmt.Println("methodValue is nil:", vMethod.IsNil()) // false
}

func interfaceTest() {
	// 未初始化的接口（本身为 nil）
	var nilInterface interface{}
	vNilInterface := reflect.ValueOf(nilInterface)
	fmt.Println("vNilInterface.IsValid():", vNilInterface.IsValid()) // false
	// vNilInterface.IsNil() // panic: 无效值

	// 接口持有 nil 指针
	var nilPtr *int
	var interfaceWithNil interface{} = nilPtr
	vInterfaceWithNil := reflect.ValueOf(interfaceWithNil)
	fmt.Println("interfaceWithNil is nil:", vInterfaceWithNil.IsNil()) // true

}

func pointerTest() {
	// nil 指针
	var nilPtr *int
	vNilPtr := reflect.ValueOf(nilPtr)
	fmt.Println("nilPtr is nil:", vNilPtr.IsNil()) // true

	// 指向有效值的指针
	x := 10
	ptr := &x
	vPtr := reflect.ValueOf(ptr)
	fmt.Println("ptr is nil:", vPtr.IsNil()) // false

	// 指向 nil 的指针
	var target *int = nil
	ptrToNil := &target
	vPtrToNil := reflect.ValueOf(ptrToNil)
	fmt.Println("ptrToNil is nil:", vPtrToNil.IsNil()) // false
}

func sliceTest() {
	// nil 切片
	var nilSlice []int
	vNilSlice := reflect.ValueOf(nilSlice)
	fmt.Println("nilSlice is nil:", vNilSlice.IsNil()) // true

	// 空切片（非 nil）
	emptySlice := []int{}
	vEmptySlice := reflect.ValueOf(emptySlice)
	fmt.Println("emptySlice is nil:", vEmptySlice.IsNil()) // false

	// 有内容的切片
	fullSlice := []int{1, 2, 3}
	vFullSlice := reflect.ValueOf(fullSlice)
	fmt.Println("fullSlice is nil:", vFullSlice.IsNil()) // false

	// 从数组创建的切片
	arr := [3]int{1, 2, 3}
	sliceFromArr := arr[:]
	vSliceFromArr := reflect.ValueOf(sliceFromArr)
	fmt.Println("sliceFromArr is nil:", vSliceFromArr.IsNil()) // false
}

func unsafePointerTest() {
	// nil 不安全指针
	var nilUnsafePtr unsafe.Pointer
	vNilUnsafePtr := reflect.ValueOf(nilUnsafePtr)
	fmt.Println("nilUnsafePtr is nil:", vNilUnsafePtr.IsNil()) // true

	// 指向有效地址的不安全指针
	var x int = 42
	unsafePtr := unsafe.Pointer(&x)
	vUnsafePtr := reflect.ValueOf(unsafePtr)
	fmt.Println("unsafePtr is nil:", vUnsafePtr.IsNil()) // false

	// 指向 nil 的不安全指针
	var nilTarget *int = nil
	unsafePtrToNil := unsafe.Pointer(nilTarget)
	vUnsafePtrToNil := reflect.ValueOf(unsafePtrToNil)
	fmt.Println("unsafePtrToNil is nil:", vUnsafePtrToNil.IsNil()) // true
}
