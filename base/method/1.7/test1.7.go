package main

import (
	"fmt"
	"reflect"
)

type T struct{}

func (T) M1() {
	fmt.Println("T M1 ")
}
func (T) M2() {
	fmt.Println("T M2 ")
}

func (*T) M3() {
	fmt.Println("*T M3 ")
}
func (*T) M4() {
	fmt.Println("*T M4 ")
}

func main() {
	/*
		两者方法集合不同==============
	*/
	var n int
	dumpMethodSet(n)
	dumpMethodSet(&n)

	var t T
	dumpMethodSet(t)
	dumpMethodSet(&t)

	/*
		方法表达式与各自类型的方法集合相同=============
	*/
	var t1 T
	f1 := (*T).M1
	f2 := (*T).M2
	f3 := (*T).M3
	f4 := (*T).M4
	fmt.Printf("the type of f1 is %T\n", f1) // the type of f1 is func(*main.T)
	fmt.Printf("the type of f2 is %T\n", f2) // the type of f2 is func(*main.T)
	fmt.Printf("the type of f3 is %T\n", f3) // the type of f3 is func(*main.T)
	fmt.Printf("the type of f4 is %T\n", f4) // the type of f4 is func(*main.T)
	f1(&t1)
	f2(&t1)
	f3(&t1)
	f4(&t1)

	f11 := T.M1
	f12 := T.M2
	//f13 := T.M3 // cannot call pointer method M3 on T
	//f14 := T.M4 // cannot call pointer method M4 on T
	fmt.Printf("the type of f11 is %T\n", f11)
	fmt.Printf("the type of f12 is %T\n", f12)
	//fmt.Printf("the type of f13 is %T\n", f13)
	//fmt.Printf("the type of f14 is %T\n", f14)
	f11(t1)
	f12(t1)
	//f13(t1)
	//f14(t1)

	/*
		方法调用没有区别==============
	*/
	t1.M1()
	t1.M2()
	t1.M3()
	t1.M4()

	t2 := &T{}
	t2.M1()
	t2.M2()
	t2.M3()
	t2.M4()
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
