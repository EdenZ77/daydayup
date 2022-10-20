package main

import (
	"fmt"
	"reflect"
)

/*
在使用反射时，需要首先理解类型（Type）和种类（Kind）的区别。编程中，使用最多的是类型，但在反射中，当需要区分一个大品种的类型时，就会用到种类（Kind）。
例如需要统一判断类型中的指针时，使用种类（Kind）信息就较为方便。

Go语言程序中的类型（Type）指的是系统原生数据类型，如 int、string、bool、float32 等类型，以及使用 type 关键字定义的类型，这些类型的名称就是其类型本身的名称。
例如使用 type A struct{} 定义结构体时，A 就是 struct{} 的类型。
*/

type Enum int

const (
	Zero Enum = 0
)

func main() {
	// 声明一个空结构体
	type cat struct {
	}
	// 获取结构体实例的反射类型对象
	typeOfCat := reflect.TypeOf(cat{})
	typeOfCatPtr := reflect.TypeOf(&cat{})
	// 显示反射类型对象的名称和种类
	fmt.Println(typeOfCat.Name(), typeOfCat.Kind())                        // cat struct
	fmt.Println("name:"+typeOfCatPtr.Name(), "kind:", typeOfCatPtr.Kind()) // name: kind: ptr
	// 获取Zero常量的反射类型对象
	typeOfA := reflect.TypeOf(Zero)
	// 显示反射类型对象的名称和种类
	fmt.Println(typeOfA.Name(), typeOfA.Kind()) // Enum int
}
