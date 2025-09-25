package main

import "fmt"

type Demo struct {
	name string
}

func createDemo(name string) *Demo {
	d := new(Demo) // 局部变量 d 逃逸到堆
	d.name = name
	return d
}

func test(demo *Demo) {
	fmt.Println(demo.name)
}

//func main() {
//	demo := createDemo("demo")
//	test(demo)
//}

//func main() {
//	demo := createDemo("demo")
//	fmt.Println(demo)
//}

func Increase() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}

func main() {
	in := Increase()
	fmt.Println(in()) // 1
	fmt.Println(in()) // 2
}
