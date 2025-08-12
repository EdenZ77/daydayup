package main

import (
	"fmt"
	"github.com/jinzhu/copier"
)

type DeepStructure struct {
	IDs     []int          // 切片
	Info    map[string]int // 映射
	Extra   *Nested        // 指针结构
	Address *Address
}

type Nested struct {
	Tags []string
}

type Address struct {
	Street string
	City   string
}

func main() {
	// 1. 准备源数据
	srcAddr := &Address{Street: "123 Main St", City: "Springfield"}
	src := DeepStructure{
		IDs:  []int{1, 2, 3},
		Info: map[string]int{"A": 10, "B": 20},
		Extra: &Nested{
			Tags: []string{"tag1", "tag2"},
		},
		Address: srcAddr,
	}

	// 2. 创建目标结构
	var shallowDest DeepStructure
	var deepDest DeepStructure

	// 3. 执行不同复制方式
	copier.Copy(&shallowDest, &src) // 默认浅复制
	copier.CopyWithOption(&deepDest, &src, copier.Option{DeepCopy: true})

	// 4. 修改源数据（影响浅复制目标，不影响深复制目标）
	src.IDs[0] = 100              // 修改切片元素
	src.Info["A"] = 100           // 修改映射值
	src.Extra.Tags[0] = "changed" // 修改嵌套指针
	src.Address.Street = "srcxxxx"

	// 5. 修改深复制目标（确保完全独立）
	deepDest.IDs = append(deepDest.IDs, 4) // 添加新元素
	deepDest.Info["C"] = 30                // 添加新键
	deepDest.Extra.Tags = []string{"new"}  // 替换整个切片
	deepDest.Address.Street = "deepppp"

	// 6. 打印结果
	fmt.Println("=== 源数据修改后 ===")
	fmt.Printf("Source IDs: %v\n", src.IDs)          // [100 2 3]
	fmt.Printf("Source Info: %v\n", src.Info)        // map[A:100 B:20]
	fmt.Printf("Source Tags: %v\n", src.Extra.Tags)  // [changed tag2]
	fmt.Printf("Source Address: %#v\n", src.Address) // Source Address: &main.Address{Street:"srcxxxx", City:"Springfield"}

	fmt.Println("\n=== 浅复制目标 ===")
	fmt.Printf("Shallow IDs: %v\n", shallowDest.IDs)          // [100 2 3] → 受源修改影响
	fmt.Printf("Shallow Info: %v\n", shallowDest.Info)        // map[A:100 B:20] → 受源修改影响
	fmt.Printf("Shallow Tags: %v\n", shallowDest.Extra.Tags)  // [changed tag2] → 受源修改影响
	fmt.Printf("Shallow Address: %#v\n", shallowDest.Address) // Shallow Address: &main.Address{Street:"123 Main St", City:"Springfield"}

	fmt.Println("\n=== 深复制目标 ===")
	fmt.Printf("Deep IDs: %v\n", deepDest.IDs)          // [1 2 3 4] → 独立
	fmt.Printf("Deep Info: %v\n", deepDest.Info)        // map[A:10 B:20 C:30] → 独立
	fmt.Printf("Deep Tags: %v\n", deepDest.Extra.Tags)  // [new] → 独立
	fmt.Printf("Deep Address: %#v\n", deepDest.Address) // Deep Address: &main.Address{Street:"deepppp", City:"Springfield"}
}
