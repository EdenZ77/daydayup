package main

import "fmt"

func main() {
	fmt.Println(Add[int](1, 2))

	// 错误，匿名函数不能自己定义类型实参
	//fnGeneric := func[T comparable | float32](a, b T) T {
	//	return a + b
	//}
	//
	//fmt.Println(fnGeneric(1, 2))
}

func Add[T int | float32 | float64](a T, b T) T {
	return a + b
}

type MyError interface { // 接口中只有方法，所以是基本接口
	Error() string
}

// 用法和 Go1.18之前保持一致
var err MyError = fmt.Errorf("hello world")
