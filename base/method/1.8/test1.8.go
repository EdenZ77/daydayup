package main

import (
	"fmt"
	"reflect"
)

type I interface {
	M1()
	M2()
}

type T struct {
	I
}

func (T) M3() {
	fmt.Println("T的M3方法")
}

func (*T) M4() {
	fmt.Println("T的M4方法")
}

func main() {
	var t T
	var p *T
	var newP = new(T)
	dumpMethodSet(t)
	dumpMethodSet(p)
	t.M3()
	t.M4()
	newP.M3()
	newP.M4()
	newP.M1()
	newP.M2()
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
