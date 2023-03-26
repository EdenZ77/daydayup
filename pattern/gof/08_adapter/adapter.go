package main

import "fmt"

// V5 适配的目标
type V5 interface {
	Use5V()
}

// Phone 业务类，依赖V5接口
type Phone struct {
	v V5
}

func NewPhone(v V5) *Phone {
	return &Phone{v}
}

func (p *Phone) Charge() {
	fmt.Println("Phone进行充电...")
	p.v.Use5V()
}

// V220 被适配的角色，适配者
type V220 struct{}

func (v *V220) Use220V() {
	fmt.Println("使用220V的电压")
}

// Adapter 电源适配器
type Adapter struct {
	v220 *V220
}

func (a *Adapter) Use5V() {
	fmt.Println("使用适配器进行充电")

	//调用适配者的方法
	a.v220.Use220V()
}

func NewAdapter(v220 *V220) *Adapter {
	return &Adapter{v220}
}

// ------- 业务逻辑层 -------
func main() {
	// 目标要使用的接口与适配者提供的无法兼容，这时使适配器实现目标要使用的接口，但是内部将调用适配者提供的接口，然后进行调整，以便最终的返回参数满足目标要求
	iphone := NewPhone(NewAdapter(new(V220)))

	iphone.Charge()
}
