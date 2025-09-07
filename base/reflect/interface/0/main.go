package main

import (
	"fmt"
	"reflect"
)

type User struct {
	PublicName string // 导出字段
	privateID  int    // 未导出字段
}

func main() {
	printValue()
}

/*
场景 1：安全处理结构体字段
CanInterface() bool方法详解
安全检查：判断是否可以对 reflect.Value安全调用 Interface() 方法
防止 panic：如果返回 false，调用 Interface() 会引发 panic
导出性检查：核心作用是检查值是否来自未导出（私有）的结构体字段
*/
func canTest() {
	u := User{PublicName: "Alice", privateID: 123}
	v := reflect.ValueOf(u)

	// 处理导出字段
	publicField := v.FieldByName("PublicName")
	if publicField.CanInterface() {
		value := publicField.Interface().(string)
		fmt.Println("Public field:", value) // 输出: Alice
	}

	// 处理未导出字段
	privateField := v.FieldByName("privateID")
	if privateField.CanInterface() {
		// 这里不会执行，因为 CanInterface() 返回 false
		value := privateField.Interface().(int)
		fmt.Println("Private field:", value)
	} else {
		fmt.Println("Cannot access private field directly")
		// 替代方案：使用反射方法读取值
		if privateField.Kind() == reflect.Int {
			fmt.Println("Private field value:", privateField.Int())
		}
	}
}

func PrintValue(v reflect.Value) {
	if !v.IsValid() {
		fmt.Println("<invalid value>")
		return
	}

	if v.CanInterface() {
		fmt.Printf("Value: %v (type: %T)\n", v.Interface(), v.Interface())
	} else {
		fmt.Printf("Value: %v (unexported field, kind: %s)\n", v, v.Kind())
	}
}

func printValue() {
	type Example struct {
		A string
		b int
	}

	e := Example{A: "Public", b: 42}
	ev := reflect.ValueOf(e)

	PrintValue(ev.FieldByName("A")) // Value: Public (type: string)
	PrintValue(ev.FieldByName("b")) // Value: 42 (unexported field, kind: int)
}
