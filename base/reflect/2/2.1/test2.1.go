package main

import (
	"fmt"
	"reflect"
)

func main() {
	pointerTest()
}

func pointerTest() {
	// 声明一个空结构体
	type cat struct {
	}
	// ins 是一个指向 cat 结构体的指针（类型为 *cat）
	// ins 不是结构体本身，而是一个指针（内存地址）。
	ins := &cat{}
	// 获取 ins 的反射类型
	typeOfCat := reflect.TypeOf(ins)
	// Name(): 指针类型（如 *cat）没有名称，所以返回空字符串 ""。
	// Kind(): 指针的种类是 reflect.Ptr（输出为 ptr）。
	fmt.Printf("name:'%v' kind:'%v'\n", typeOfCat.Name(), typeOfCat.Kind()) // name:'' kind:'ptr'
	// 获取指针指向的实际类型
	typeOfCat = typeOfCat.Elem()
	// 显示反射类型对象的名称和种类
	fmt.Printf("element name: '%v', element kind: '%v'\n", typeOfCat.Name(), typeOfCat.Kind())
}

func noPointerTest() {
	// 声明一个空结构体
	type cat struct {
	}
	// 创建cat的实例
	ins := cat{}
	// 获取结构体实例的反射类型对象
	typeOfCat := reflect.TypeOf(ins)
	// 显示反射类型对象的名称和种类
	fmt.Printf("name:'%v' kind:'%v'\n", typeOfCat.Name(), typeOfCat.Kind()) // name:'' kind:'ptr'
	// 取类型的元素
	typeOfCat = typeOfCat.Elem()
	// 显示反射类型对象的名称和种类
	fmt.Printf("element name: '%v', element kind: '%v'\n", typeOfCat.Name(), typeOfCat.Kind())
}
