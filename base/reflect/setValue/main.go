package main

import (
	"fmt"
	"reflect"
)

/*
reflect.Value是通过reflect.ValueOf(X)获得的，只有当X是指针的时候，才可以通过reflec.Value修改实际变量X的值，
即：要修改反射类型的对象就一定要保证其值是“addressable”的。
*/

func main() {

	var num float64 = 1.2345
	fmt.Println("old value of pointer:", num) // old value of pointer: 1.2345

	// 需要传入的参数是* float64这个指针，然后可以通过pointer.Elem()去获取所指向的Value，注意一定要是指针。
	pointer := reflect.ValueOf(&num)
	// 通过reflect.ValueOf获取num中的reflect.Value，注意，参数必须是指针才能修改其值
	fmt.Println("old type of pointer:", pointer.Type()) // old type of pointer: *float64
	newValue := pointer.Elem()

	fmt.Println("new type of pointer:", newValue.Type()) // new type of pointer: float64
	// newValue.CantSet()表示是否可以重新设置其值，如果输出的是true则可修改，否则不能修改
	fmt.Println("settability of pointer:", newValue.CanSet()) // settability of pointer: true

	// 重新赋值
	newValue.SetFloat(77)
	fmt.Println("new value of pointer:", num) // new value of pointer: 77

	////////////////////
	// 如果reflect.ValueOf的参数不是指针，会如何？
	pointer = reflect.ValueOf(num)
	//newValue = pointer.Elem() // 如果非指针，这里直接panic，“panic: reflect: call of reflect.Value.Elem on float64 Value”
}

/*
func main() {
    i := 1
    v := reflect.ValueOf(&i)
    v.Elem().SetInt(10)
    fmt.Println(i)
}

$ go run reflect.go
10

这种获取指针对应的 reflect.Value 并通过 Elem 方法迂回的方式就能够获取到可以被设置的变量，这一复杂的过程主要也是因为 Go 语言的函数调用都是值传递的，我们可以将上述代码理解成：
func main() {
    i := 1
    v := &i
    *v = 10
}

如果不能直接操作 i 变量修改其持有的值，我们就只能获取 i 变量所在地址并使用 *v 修改所在地址中存储的整数。


*/
