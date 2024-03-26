package main

import (
	"errors"
	"fmt"
	"io"
)

func main() {
	err1 := errors.New("EOF")
	err2 := errors.New("EOF")

	fmt.Println(err1 == io.EOF)   // false
	fmt.Println(err1 == err2)     // false
	fmt.Println(io.EOF == io.EOF) // true
	printNonEmptyInterface()
	//printNonEmptyInterface1()
}

type T struct {
	name string
}

func (t T) Error() string {
	return "bad error"
}
func printNonEmptyInterface1() {
	var err1 error    // 非空接口类型
	var err1ptr error // 非空接口类型
	var err2 error    // 非空接口类型
	var err2ptr error // 非空接口类型

	err1 = T{"eden"}
	err1ptr = &T{"eden"}

	err2 = T{"eden"}
	err2ptr = &T{"eden"}

	println("err1:", err1)
	println("err2:", err2)
	println("err1 = err2:", err1 == err2)             // true
	println("err1ptr = err2ptr:", err1ptr == err2ptr) // false
}

func printNonEmptyInterface() {
	var err1 error    // 非空接口类型
	var err1ptr error // 非空接口类型
	var err2 error    // 非空接口类型
	var err2ptr error // 非空接口类型
	err1 = (*T)(nil)
	println("err1:", err1)
	println("err1 = nil:", err1 == nil)

	err1 = T{"eden"}
	err1ptr = &T{"eden"}
	t1 := T{"eden"}
	t1ptr := &T{"eden"}

	err2 = T{"eden"}
	err2ptr = &T{"eden"}
	t2 := T{"eden"}
	t2ptr := &T{"eden"}

	println("err1:", err1)
	println("err2:", err2)
	println("err1 = err2:", err1 == err2)             // true
	println("err1ptr = err2ptr:", err1ptr == err2ptr) // false
	// 结构体比较：自动对比内部的属性是否相等，但是如果结构体内部含有map,slice,Function这类，使用==编译会报错
	println("t1 = t2:", t1 == t2)             // true
	println("t1ptr = t2ptr:", t1ptr == t2ptr) // false

	//err2 = fmt.Errorf("%d\n", 5)
	//println("err1:", err1)
	//println("err2:", err2)
	//println("err1 = err2:", err1 == err2)
}

func printNilInterface() {
	// nil接口变量
	var i interface{} // 空接口类型
	var err error     // 非空接口类型
	println(i)
	println(err)
	println("i = nil:", i == nil)
	println("err = nil:", err == nil)
	println("i = err:", i == err)
}
