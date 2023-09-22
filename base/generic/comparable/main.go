package main

/*
https://time.geekbang.org/column/article/601128?screen=full
*/
type foo struct {
	a int
	s string
}

type bar struct {
	a  int
	sl []string
}

func doSomething[T comparable](t T) T {
	var a T
	if a == t {
	}

	if a != t {
	}
	return a
}

/*
根据其注释说明，所有可比较的类型都实现了 comparable 这个接口，
包括：布尔类型、数值类型、字符串类型、指针类型、channel 类型、元素类型实现了 comparable 的数组和成员类型均实现了 comparable 接口的结构体类型。
*/

func main() {
	doSomething(true)
	doSomething(3)
	doSomething(3.14)
	doSomething(3 + 4i)
	doSomething("hello")
	var p *int
	doSomething(p)
	doSomething(make(chan int))
	doSomething([3]int{1, 2, 3})
	doSomething(foo{})
	//doSomething(bar{}) //  bar does not implement comparable
}
