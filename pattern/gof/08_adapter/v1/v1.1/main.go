package main

import "fmt"

/*
参考资料：https://time.geekbang.com/column/article/205912

在Go语言中，由于缺少继承机制，我们不能直接使用类似Java中的类适配器模式。但我们可以通过嵌入（匿名结构体）来模拟出类似的效果。
嵌入可以使得嵌入的类型（Adaptee）的方法被外部类型（Adapter）所拥有，就像它们是外部类型自己的方法一样。
这为模拟类适配器模式提供了便利，因为外部类型可以重写嵌入类型的方法，从而实现适配。


*/

// ITarget 是客户所期望的接口
type ITarget interface {
	F1()
	F2()
	Fc()
}

// Adaptee 是被适配的类，有一些自己的方法
type Adaptee struct{}

func (a *Adaptee) Fa() {
	fmt.Println("Adaptee method Fa")
}

func (a *Adaptee) Fb() {
	fmt.Println("Adaptee method Fb")
}

func (a *Adaptee) Fc() {
	fmt.Println("Adaptee method Fc")
}

// Adapter 嵌入了Adaptee，可以重写Fa和Fb，还可以实现ITarget接口
type Adapter struct {
	Adaptee
}

func (a *Adapter) F1() {
	a.Fa() // 调用Adaptee的Fa方法
}

func (a *Adapter) F2() {
	// Adapter的F2方法实现
	fmt.Println("Adapter method F2")
}

// Fc方法直接继承自Adaptee，无需重新实现

func main() {
	adapter := Adapter{}
	adapter.F1()
	adapter.F2()
	adapter.Fc()
}
