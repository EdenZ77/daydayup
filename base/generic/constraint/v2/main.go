package main

import "fmt"

// stringify_new_without_zero.go
type Stringer interface {
	comparable // 编译器会判断想要实现 Stringer 接口的类型是否实现了 comparable 接口
	String() string
}

func StringifyWithoutZero[T Stringer](s []T) (ret []string) {
	var zero T
	for _, v := range s {
		if v == zero {
			continue
		}
		ret = append(ret, v.String())
	}
	return ret
}

type MyString struct {
	//sl []string   如果这里是slice，那么就不是comparable了
	xx string
}

//type MyString string // 如果这里是string，那么就是comparable了

func (s MyString) String() string {
	return string(s.xx)
}

func main() {
	//sl := StringifyWithoutZero([]MyString{"I", "", "love", "", "golang"}) // 输出：[I love golang]
	sl := StringifyWithoutZero([]MyString{{"I"}, {"P"}}) // 输出：[I love golang]
	fmt.Println(sl)
}
