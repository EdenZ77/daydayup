package main

import "fmt"

type student struct {
	name string
	age  int
}

// 指针切片语法糖的运用，以及for循环中变量本质
func main() {
	m := make(map[string]*student)
	stus := []*student{
		{"小王子", 18},
		{"小王子2", 23},
		{"小王子3", 26},
	}

	for _, stu := range stus {
		m[stu.name] = stu
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
}
