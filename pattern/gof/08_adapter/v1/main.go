package main

import "fmt"

/*
参考资料：https://time.geekbang.com/column/article/205912
以下对象适配器模式的实现，在v1.1中实现了类适配器模式
*/

// ITarget 是目标接口，客户期望的接口
type ITarget interface {
	F1()
	F2()
	Fc()
}

// Adaptee 是被适配的目标类，有自己的方法集合
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

// Adapter 实现了ITarget接口，通过内部持有Adaptee的实例，来实现方法的适配
type Adapter struct {
	adaptee *Adaptee
}

func NewAdapter(adaptee *Adaptee) *Adapter {
	return &Adapter{
		adaptee: adaptee,
	}
}

func (a *Adapter) F1() {
	a.adaptee.Fa() // 调用Adaptee的Fa方法来实现F1方法
}

func (a *Adapter) F2() {
	// Adapter重新实现F2方法
	fmt.Println("Adapter method F2")
}

func (a *Adapter) Fc() {
	a.adaptee.Fc() // 直接委托给Adaptee的Fc方法
}

func main() {
	adaptee := &Adaptee{}
	adapter := NewAdapter(adaptee)
	adapter.F1()
	adapter.F2()
	adapter.Fc()
}
