package main

import (
	"fmt"
)

// A 类型
type A struct{}

// F 是 A 类型的方法，它将用作回调函数
func (a *A) F() {
	fmt.Println("A's F() method is called")
}

// B 类型
type B struct {
	callbackFunc func() // 用于存储回调函数的字段
}

// RegisterCallback 用于注册回调函数到 B 类型
func (b *B) RegisterCallback(f func()) {
	b.callbackFunc = f
}

// P 是 B 类型的方法，它会调用之前注册的回调函数
func (b *B) P() {
	fmt.Println("B's P() method is called")
	if b.callbackFunc != nil {
		b.callbackFunc() // 调用回调函数
	}
}

func main() {
	a := &A{}
	b := &B{}

	// 将 A 的 F 方法注册为 B 的回调函数
	b.RegisterCallback(a.F)

	// 调用 B 的 P 方法，这将触发回调函数
	b.P()
}
