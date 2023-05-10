package main

import "fmt"

// doSomethingOption 定义一个内部使用的配置项结构体
// 类型名称及字段的首字母小写（包内私有）
type doSomethingOption struct {
	a string
	b int
	c bool
	// ...
}

// 此时，同样是使用函数选项模式，但我们这一次通过使用接口类型来“隐藏”内部的逻辑。

// IOption 定义一个接口类型
type IOption interface {
	apply(*doSomethingOption)
}

// funcOption 定义funcOption类型，实现 IOption 接口
type funcOption struct {
	f func(*doSomethingOption)
}

func (fo funcOption) apply(o *doSomethingOption) {
	fo.f(o)
}

func newFuncOption(f func(*doSomethingOption)) IOption {
	return &funcOption{
		f: f,
	}
}

// WithB 将b字段设置为指定值的函数
func WithB(b int) IOption {
	return newFuncOption(func(o *doSomethingOption) {
		// 这里实现赋值逻辑
		o.b = b
	})
}

// DoSomething 包对外提供的函数
func DoSomething(a string, opts ...IOption) {
	o := &doSomethingOption{a: a}
	for _, opt := range opts {
		opt.apply(o)
	}
	// 在包内部基于o实现逻辑...
	fmt.Printf("o:%#v\n", o)
}

/*
如此一来，我们只需对外提供一个DoSomething的功能函数和一系列WithXxx函数。对于调用方来说，使用起来也很方便。

其实我个人看来，实现接口意义不大，增加了理解难度
*/
func main() {

	DoSomething("q1mi")
	DoSomething("q1mi", WithB(100))
}
