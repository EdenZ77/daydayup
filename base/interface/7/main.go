package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

//func main() {
//	var n = new(int)
//	*n = 11
//	var ei interface{} = n
//	*n = 62                         // n的值已经改变
//	fmt.Println("data in box:", ei) // 输出仍是61
//	//var a = new(int)
//	//a = ei
//	//fmt.Println()
//}

func main() {
	// 示例1：处理空接口
	var a interface{} = 42

	// 使用反射获取空接口的数据指针
	efaceValue := reflect.ValueOf(a)
	efaceDataPtr := unsafe.Pointer(efaceValue.Pointer())

	// 打印数据指针
	fmt.Printf("eface data pointer: %v\n", efaceDataPtr)

	// 使用数据指针获取实际值
	value := *(*int)(efaceDataPtr)
	fmt.Printf("Value: %d\n", value)

	// 示例2：处理具体接口
	type Stringer interface {
		String() string
	}

	type MyString string
	//
	//func (s MyString) String() string {
	//	return string(s)
	//}
	//
	//var x Stringer = MyString("hello")
	//
	//// 使用反射获取具体接口的数据指针
	//ifaceValue := reflect.ValueOf(x)
	//// 由于具体接口的结构稍复杂，通过两步来获取数据指针
	//ifaceDataPtr := unsafe.Pointer(ifaceValue.Elem().UnsafeAddr())
	//
	//// 打印数据指针
	//fmt.Printf("iface data pointer: %v\n", ifaceDataPtr)
	//
	//// 使用数据指针获取实际值
	//strValue := *(*MyString)(ifaceDataPtr)
	//fmt.Printf("Value: %s\n", strValue)
}
