package main

import "fmt"

// TestOption
/*
我们可能需要为Option的字段指定默认值
Option的字段成员可能会发生变更

可阅读grpc源码中dialOptions使用的option模式, 它引入了接口，和示例稍有不同
*/
type TestOption struct {
	A string
	B string
	C int
}

func newTestOption(a, b string, c int) *TestOption {
	return &TestOption{
		A: a,
		B: b,
		C: c,
	}
}

type OptionFunc func(*TestOption)

// WithA 利用闭包为每个字段编写一个设置值的With函数
func WithA(a string) OptionFunc {
	return func(o *TestOption) {
		o.A = a
	}
}

func WithB(b string) OptionFunc {
	return func(o *TestOption) {
		o.B = b
	}
}

func WithC(c int) OptionFunc {
	return func(o *TestOption) {
		o.C = c
	}
}

var (
	defaultOption = &TestOption{
		A: "A",
		B: "B",
		C: 100,
	}
)

func newOption2(opts ...OptionFunc) (opt *TestOption) {
	opt = defaultOption
	for _, o := range opts {
		o(opt)
	}
	return
}

func main() {
	//x := newTestOption("nazha", "小王子", 10)
	//fmt.Println(x)
	//x = newOption2()
	//fmt.Println(x)

	// 思想就是当不传入任何参数时，会有默认值，当需要修改某个字段时，传入相应的WithXxx(value)
	// 将修改行为封装成了函数，这可能是go的设计哲学吧。
	x := newOption2(
		WithA("沙河娜扎"),
		WithC(250),
	)
	fmt.Println(x)
}
