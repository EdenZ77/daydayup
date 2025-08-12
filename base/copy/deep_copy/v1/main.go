package main

import "fmt"

type Person struct {
	Name    string
	Address *Address
}

type Address struct {
	Street string
	City   string
}

func DeepCopyPerson(p *Person) *Person {
	if p == nil {
		return nil
	}
	copyVar := &Person{
		Name: p.Name, // string 是值类型（不可变），直接复制安全
		Address: &Address{ // 为嵌套引用类型 *Address 创建新对象并复制字段
			Street: p.Address.Street,
			City:   p.Address.City,
		},
	}
	return copyVar
}

// 手动深拷贝 (Manual Deep Copy):
// 为复杂结构（尤其是包含引用字段的 struct) 显式编写代码，为新结构分配内存并递归复制所有数据。

func main() {
	// 使用
	original := &Person{Name: "Alice", Address: &Address{"123 Main", "Springfield"}}
	deepCopy := DeepCopyPerson(original)
	deepCopy.Address.City = "Metropolis" // 仅修改副本的地址
	fmt.Println(original.Address.City)   // Springfield (未变)
	fmt.Println(deepCopy.Address.City)   // Metropolis
}
