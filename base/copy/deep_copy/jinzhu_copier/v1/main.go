package main

import (
	"fmt"
	"github.com/jinzhu/copier"
)

// 第三方库 (Third-Party Libraries):
// 社区中有专为深拷贝设计的库，它们通常提供更简洁的接口和更好的性能，能处理循环引用等问题。
// 流行的库： github.com/jinzhu/copier：功能强大，支持选项配置。

// copier.Copy默认执行的是浅拷贝，但对指针字段特殊处理：
// 当遇到指针字段时 *Address，copier会创建新指针并拷贝指针指向的内容
// 最终效果相当于深度复制了指针指向的结构体

type Person struct {
	Name    string
	Address *Address
}

type Address struct {
	Street string
	City   string
}

func main() {
	// 原始数据
	srcAddr := &Address{Street: "123 Main St", City: "Springfield"}
	srcPerson := Person{Name: "John", Address: srcAddr}
	//srcPerson := &Person{Name: "John", Address: srcAddr}

	// 复制
	var destPerson Person
	//destPerson := &Person{}
	copier.Copy(&destPerson, &srcPerson) // 浅拷贝

	// 修改复制后的地址
	destPerson.Address.Street = "456 Oak Ave"
	fmt.Println("Original Street:", srcPerson.Address.Street) // 输出: 123 Main St
	fmt.Println("Copied Street: ", destPerson.Address.Street) // 输出: 456 Oak Ave
}
