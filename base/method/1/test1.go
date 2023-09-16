package main

import "fmt"

type T struct {
	a int
}

func (t T) Get() int {
	return t.a
}

func (t *T) Set(a int) int {
	t.a = a
	return t.a
}

func main() {
	var t T
	f1 := (*T).Set                           // f1的类型，也是*T类型Set方法的类型：func (t *T, int)int
	f2 := T.Get                              // f2的类型，也是T类型Get方法的类型：func(t T)int
	fmt.Printf("the type of f1 is %T\n", f1) // the type of f1 is func(*main.T, int) int
	fmt.Printf("the type of f2 is %T\n", f2) // the type of f2 is func(main.T) int
	f1(&t, 3)
	fmt.Println(f2(t)) // 3
}

// Get 类型T的方法Get的等价函数
func Get(t T) int {
	return t.a
}

// Set 类型*T的方法Set的等价函数
func Set(t *T, a int) int {
	t.a = a
	return t.a
}
