package main

import (
	"fmt"
	"time"
)

type Field struct {
	name string
}

//func (p *Field) print() {
//	fmt.Println(p.name)
//}

// 解决办法二：修改方法的receiver
func (p Field) print() {
	fmt.Println(p.name)
}

func main() {
	data1 := []*Field{{"one"}, {"two"}, {"three"}}
	for _, v := range data1 {
		//v := v // 最佳的解决办法
		//go v.print()  // 等价于 go Field.print(v)
		go Field.print(*v)
	}

	data2 := []Field{{"four"}, {"five"}, {"six"}}
	for _, v := range data2 {
		//v := v // 最佳的解决办法
		//go v.print() // 等价于 go Field.print(&v)
		go Field.print(v)
	}

	time.Sleep(3 * time.Second)
}
