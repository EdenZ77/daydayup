package main

import "fmt"

// 参考资料：GPT

// ITemplate 定义一个接口，包含模板方法和需要子类实现的方法
type ITemplate interface {
	TemplateMethod()
	PrimitiveOperation1()
	PrimitiveOperation2()
}

// BaseTemplate 定义一个基本的模板结构体，实现了ITemplate接口的TemplateMethod方法
type BaseTemplate struct {
	ITemplate
}

func (t *BaseTemplate) TemplateMethod() {
	// 这种写法将先掉用具体子类的重写方法，当没有重写方法时才会调用base的默认方法
	t.ITemplate.PrimitiveOperation1()
	// 这种写法将只调用Base的默认方法，无法调用具体子类的重写方法；只有当base没有默认方法时才会调用具体子类的重写方法
	t.PrimitiveOperation2()
}

// PrimitiveOperation1 Base提供默认实现
func (t *BaseTemplate) PrimitiveOperation1() {
	fmt.Println("BaseTemplate PrimitiveOperation1===")
}

// ConcreteTemplate1 定义一个子类，继承BaseTemplate，并实现PrimitiveOperation1和PrimitiveOperation2方法
type ConcreteTemplate1 struct {
	BaseTemplate
}

// 由于base提供了默认实现，所以具体类不提供该方法时运行将不再报错
//func (t *ConcreteTemplate1) PrimitiveOperation1() {
//	fmt.Println("ConcreteTemplate1 PrimitiveOperation1")
//}

func (t *ConcreteTemplate1) PrimitiveOperation2() {
	fmt.Println("ConcreteTemplate1 PrimitiveOperation2")
}

// ConcreteTemplate2 定义另一个子类，继承BaseTemplate，并实现PrimitiveOperation1和PrimitiveOperation2方法
type ConcreteTemplate2 struct {
	BaseTemplate
}

func (t *ConcreteTemplate2) PrimitiveOperation1() {
	fmt.Println("ConcreteTemplate2 PrimitiveOperation1")
}

func (t *ConcreteTemplate2) PrimitiveOperation2() {
	fmt.Println("ConcreteTemplate2 PrimitiveOperation2")
}

func main() {
	// 创建ConcreteTemplate1和ConcreteTemplate2的实例，并调用模板方法
	concreteTemplate1 := &ConcreteTemplate1{}
	concreteTemplate1.ITemplate = concreteTemplate1
	concreteTemplate1.TemplateMethod()

	concreteTemplate2 := &ConcreteTemplate2{}
	concreteTemplate2.ITemplate = concreteTemplate2
	concreteTemplate2.TemplateMethod()
}
