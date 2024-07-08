package main

import (
	"fmt"
	"reflect"
)

//

type Interface interface {
	M1()
	M2()
}

type T struct {
	Interface
}

//func (t T) M1() {
//	fmt.Println("Im T struct T M1 method")
//}
//func (t *T) M2() {
//	fmt.Println("Im T struct *T M2 method")
//}

func (t T) M3() {
	fmt.Println("Im T struct T M3 method")
}

func main() {
	var t T
	var pt *T
	//var i Interface

	t.M1()
	t.M2()
	t.M3()

	pt.M1()
	pt.M2()
	pt.M3()

	//i = pt
	//i = t // cannot use t (type T) as type Interface in assignment: T does not implement Interface (M2 method has pointer receiver)
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
