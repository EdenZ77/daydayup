package main

import (
	"fmt"
	"reflect"
)

type T1 struct{}

func (T1) T1M1()   { println("T1's M1") }
func (*T1) PT1M2() { println("PT1's M2") }

type T2 struct{}

func (T2) T2M1()   { println("T2's M1") }
func (*T2) PT2M2() { println("PT2's M2") }

// T 垂直组合本质上是一种“能力继承”，采用嵌入方式定义的新类型继承了嵌入类型的能力

// T Go 还有一种常见的组合方式，叫水平组合。和垂直组合的能力继承不同，水平组合是一种能力委托（Delegate），我们通常使用接口类型来实现水平组合。
// Go 语言中的接口是一个创新设计，它只是方法集合，并且它与实现者之间的关系无需通过显式关键字修饰，
// 它让程序内部各部分之间的耦合降至最低，同时它也是连接程序各个部分之间“纽带”。

type T struct {
	T1
	*T2
}

// 不使用嵌套
//type T struct {
//	t1 T1
//	t2 *T2
//}

func main() {
	//t := T{
	//	T1: T1{},
	//	T2: &T2{},
	//}

	// 当不使用嵌套类型的时候
	//t := T{
	//	t1: T1{},
	//	t2: &T2{},
	//}
	// 当不使用嵌套类型的时候, 方法集合为空
	//dumpMethodSet(t)
	//dumpMethodSet(&t)

	//t.T1M1()     // T1's M1
	//t.PT1M2()    // PT1's M2
	//t.T2M1()     // T2's M1
	//t.PT2M2()    // PT2's M2
	//(&t).T1M1()  // T1's M1
	//(&t).PT1M2() // PT1's M2
	//(&t).T2M1()  // T2's M1
	//(&t).PT2M2() // PT2's M2

}

func dumpMethodSet(i interface{}) {
	dynTyp := reflect.TypeOf(i)

	if dynTyp == nil {
		fmt.Printf("there is no dynamic type\n")
		return
	}

	n := dynTyp.NumMethod()
	if n == 0 {
		fmt.Printf("%s's method set is empty!\n", dynTyp)
		return
	}

	fmt.Printf("%s's method set:\n", dynTyp)
	for j := 0; j < n; j++ {
		fmt.Println("-", dynTyp.Method(j).Name)
	}
	fmt.Printf("\n")
}
