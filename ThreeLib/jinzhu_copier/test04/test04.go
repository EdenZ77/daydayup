package main

import (
	"fmt"
	"github.com/jinzhu/copier"
)

/*
切片赋值：
使用一个切片来为另一个切片赋值。如果类型相同，那好办，直接append就行。如果类型不同呢？copier就派上大用场了
*/

type User struct {
	Name string
	Age  int
}

type Employee struct {
	Name string
	Age  int
	Role string
}

// 这个实际上就是将源切片中每个元素分别赋值到目标切片中。
func main() {
	users := []User{
		{Name: "dj", Age: 18},
		{Name: "dj2", Age: 18},
	}
	employees := []Employee{}

	copier.Copy(&employees, &users)
	fmt.Printf("%#v\n", employees)
}
