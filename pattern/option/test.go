package main

import "fmt"

// Option
/*
我们可能需要为Option的字段指定默认值
Option的字段成员可能会发生变更

可阅读grpc源码中dialOptions使用的option模式, 它引入了接口，和示例稍有不同
*/
type Option struct {
	A string
	B string
	C int
}

func newOption(a, b string, c int) *Option {
	return &Option{
		A: a,
		B: b,
		C: c,
	}
}

type OptionFunc func(*Option)

// WithA 利用闭包为每个字段编写一个设置值的With函数
func WithA(a string) OptionFunc {
	return func(o *Option) {
		o.A = a
	}
}

func WithB(b string) OptionFunc {
	return func(o *Option) {
		o.B = b
	}
}

func WithC(c int) OptionFunc {
	return func(o *Option) {
		o.C = c
	}
}

var (
	defaultOption = &Option{
		A: "A",
		B: "B",
		C: 100,
	}
)

func newOption2(opts ...OptionFunc) (opt *Option) {
	opt = defaultOption
	for _, o := range opts {
		o(opt)
	}
	return
}

func main() {
	x := newOption("nazha", "小王子", 10)
	fmt.Println(x)
	x = newOption2()
	fmt.Println(x)
	x = newOption2(
		WithA("沙河娜扎"),
		WithC(250),
	)
	fmt.Println(x)
}
