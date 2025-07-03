package main

import "fmt"

type Slice[T int | float32 | float64] []T

func main() {
	// 这里传入了类型实参int，泛型类型Slice[T]被实例化为具体的类型 Slice[int]
	var a Slice[int] = []int{1, 2, 3}
	fmt.Printf("Type Name: %T\n", a) //输出：Type Name: Slice[int]

	// 传入类型实参float32, 将泛型类型Slice[T]实例化为具体的类型 Slice[string]
	var b Slice[float32] = []float32{1.0, 2.0, 3.0}
	fmt.Printf("Type Name: %T\n", b) //输出：Type Name: Slice[float32]

	// ✗ 错误。因为变量a的类型为Slice[int]，b的类型为Slice[float32]，两者类型不同
	//a = b

	// ✗ 错误。string不在类型约束 int|float32|float64 中，不能用来实例化泛型类型
	//var c Slice[string] = []string{"Hello", "World"}

	// ✗ 错误。Slice[T]是泛型类型，不可直接使用必须实例化为具体的类型
	//var x Slice[T] = []int{1, 2, 3}
}
