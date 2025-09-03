package main

import (
	"fmt"
	"reflect"
)

// MyStruct 1. 非接口类型示例
type MyStruct struct{}

func (m MyStruct) ExportedMethod()   {} // 导出方法
func (m *MyStruct) PointerMethod()   {} // 指针接收者方法
func (m MyStruct) unexportedMethod() {} // 未导出方法

// MyInterface 2. 接口类型示例
type MyInterface interface {
	InterfaceMethod() // 接口方法
	unexported()      // 未导出接口方法（不常见但允许）
}

func main() {
	fmt.Println("=== 非接口类型（结构体）===")
	analyzeMethods(MyStruct{})

	fmt.Println("\n=== 非接口类型（结构体指针）===")
	analyzeMethods(&MyStruct{})

	fmt.Println("\n=== 接口类型 ===")
	analyzeInterfaceMethods()
}

func analyzeMethods(v interface{}) {
	t := reflect.TypeOf(v)
	fmt.Printf("类型: %s, 方法数: %d\n", t, t.NumMethod())

	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		fmt.Printf(" 方法 %d: %s\n", i, method.Name)
		fmt.Printf("  PkgPath: '%s'\n", method.PkgPath)
		fmt.Printf("  Type: %s\n", method.Type)
		fmt.Printf("  Func: %v\n", method.Func)
		fmt.Printf("  Index: %d\n", method.Index)
	}

	// 按名称查找方法
	if method, found := t.MethodByName("ExportedMethod"); found {
		fmt.Println("\n找到 'ExportedMethod':")
		fmt.Printf("  签名: %s\n", method.Type)
	}

}

func analyzeInterfaceMethods() {
	// 获取接口类型
	t := reflect.TypeOf((*MyInterface)(nil)).Elem()

	fmt.Printf("接口类型: %s, 方法数: %d\n", t, t.NumMethod())

	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		fmt.Printf(" 方法 %d: %s\n", i, method.Name)
		fmt.Printf("  PkgPath: '%s'\n", method.PkgPath) // 未导出方法的PkgPath显示包路径
		fmt.Printf("  Type: %s\n", method.Type)
		fmt.Printf("  Func: %v\n", method.Func) // Func字段总是 nil（接口只有签名没有实现）
		fmt.Printf("  Index: %d\n", method.Index)
	}
}
