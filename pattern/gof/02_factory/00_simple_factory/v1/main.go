package main

import "fmt"

/*
参考资料：https://mp.weixin.qq.com/s/MlC6-TDf06LGpF8hxcSV_w
极客时间的设计模式：关于简单工厂与工厂方法的讲解，我觉得讲的很好。


*/

// Printer 简单工厂要返回的接口类型
type Printer interface {
	Print(name string) string
}

// CnPrinter 是 Printer 接口的实现，它说中文
type CnPrinter struct{}

func (*CnPrinter) Print(name string) string {
	return fmt.Sprintf("你好, %s", name)
}

// EnPrinter 是 Printer 接口的实现，它说中文
type EnPrinter struct{}

func (*EnPrinter) Print(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}

// NewPrinter 是简单工厂函数，其实真实生产环境中，绝大部分都是这种简单工厂函数
func NewPrinter(lang string) Printer {
	switch lang {
	case "cn":
		return new(CnPrinter)
	case "en":
		return new(EnPrinter)
	default:
		return new(CnPrinter)
	}
}

func main() {
	printer := NewPrinter("en")
	fmt.Println(printer.Print("Bob"))
}
