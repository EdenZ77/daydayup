package main

import "fmt"

// stringify_without_zero.go

func StringifyWithoutZero[T fmt.Stringer](s []T) (ret []string) {
	var zero T
	for _, v := range s {
		if v == zero { // 编译器报错：invalid operation: v == zero (incomparable types in type set)
			continue
		}
		ret = append(ret, v.String())
	}
	return ret
}

type MyString string

func (s MyString) String() string {
	return string(s)
}

func main() {
	sl := StringifyWithoutZero([]MyString{"I", "love", "golang"})
	fmt.Println(sl) // 输出：[I love golang]
}
