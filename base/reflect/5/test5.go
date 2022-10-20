package main

import (
	"fmt"
	"reflect"
)

/*
IsNil()常被用于判断指针是否为空；IsValid()常被用于判定返回值是否有效。
*/
func main() {

	var a *int
	fmt.Println("var a *int IsNil:", reflect.ValueOf(a).IsNil()) // var a *int IsNil: true
	// nil值
	fmt.Println("nil IsValid:", reflect.ValueOf(nil).IsValid()) // nil IsValid: false
	// 实例化一个匿名结构体
	b := struct{}{}
	// 尝试从结构体中查找"abc"字段
	fmt.Println("不存在的结构体成员:", reflect.ValueOf(b).FieldByName("abc").IsValid()) // 不存在的结构体成员: false
	// 尝试从结构体中查找"abc"方法
	fmt.Println("不存在的结构体方法:", reflect.ValueOf(b).MethodByName("abc").IsValid()) // 不存在的结构体方法: false
	// map
	c := map[string]int{}
	// 尝试从map中查找一个不存在的键
	fmt.Println("map中不存在的键：", reflect.ValueOf(c).MapIndex(reflect.ValueOf("娜扎")).IsValid()) // map中不存在的键： false
}
