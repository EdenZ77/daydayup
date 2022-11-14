package main

import (
	"fmt"
	"reflect"
)

// 结构体类型中嵌入接口类型
// 由于嵌入某接口类型的结构体类型的方法集合包含了这个接口类型的方法集合，这就意味着，这个结构体类型也是它嵌入的接口类型的一个实现。
// 即便结构体类型自身并没有实现这个接口类型的任意一个方法，也没有关系。

type I interface {
	M1()
	M2()
}

type T struct {
	I
}

func (T) M3() {
	fmt.Println("T M3 ")
}

func main() {
	t := T{}
	var p = &t
	dumpMethodSet(t)
	dumpMethodSet(p)

	t.M3()
	p.M3()
	fmt.Println("调用接口的方法，这些方法结构体T都没有实现，运行的时候会报错，但是编译不会报错")
	t.M1()
	t.M2()
	p.M1()
	p.M2()
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
