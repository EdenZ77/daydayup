package main

import (
	"fmt"
	"reflect"
)

func main() {
	pointerChain()
}

func interfaceTest() {
	var i interface{} = 42
	v := reflect.ValueOf(&i)
	fmt.Printf("kind: %v\n", v.Kind()) // kind: ptr
	elem := v.Elem()
	fmt.Printf("elem kind: %v\n", elem.Kind()) // elem kind: interface
}

func interfaceNoTest() {
	var i interface{} = 42
	v := reflect.ValueOf(i)
	fmt.Printf("kind: %v\n", v.Kind()) // kind: int
	// panic: reflect: call of reflect.Value.Elem on int Value
	elem := v.Elem()
	fmt.Printf("elem kind: %v\n", elem.Kind())
}

func modifyInterface() {
	var i interface{} = 42

	v := reflect.ValueOf(&i).Elem() // 接口本身的反射值
	if v.Kind() == reflect.Interface {
		// 创建新值并设置
		newValue := reflect.ValueOf("hello")
		v.Set(newValue)
	}

	fmt.Println(i) // 输出: hello
}

func pointerChain() {
	x := 42
	p := &x
	pp := &p

	v := reflect.ValueOf(pp)          // **int
	fmt.Printf("v: %v\n", v.Kind())   // v: ptr
	v1 := v.Elem()                    // *int
	fmt.Printf("v1: %v\n", v1.Kind()) // v1: ptr
	v2 := v1.Elem()                   // int
	fmt.Printf("v2: %v\n", v1.Kind()) // v2: ptr
	fmt.Println(v2.Int())             // 42
}

func processInterface(i interface{}) {
	v := reflect.ValueOf(&i).Elem() // 获取接口本身的反射值

	if v.Kind() != reflect.Interface {
		fmt.Println("参数不是接口类型")
		return
	}

	elem := v.Elem()
	fmt.Printf("接口包含: 类型=%v, 值=%v\n", elem.Kind(), elem.Interface())
}

func processInterfaceTest() {
	processInterface(42)         // 接口包含: 类型=int, 值=42
	processInterface("hello")    // 接口包含: 类型=string, 值=hello
	processInterface(struct{}{}) // 接口包含: 类型=struct, 值={}
}
