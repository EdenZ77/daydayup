package main

import "fmt"

// 直接赋值 (=):
// 对于值类型 (struct, array)，这是有效的浅拷贝（因为它们本身不含引用，所以等同于独立副本）。
// 对于引用类型 (slice, map, pointer, etc.)，这只是复制了引用（指针）。

func main() {
	// 数组 (值类型) - 赋值创建独立副本
	originalArr := [3]int{1, 2, 3}
	copyArr := originalArr // 独立副本
	copyArr[0] = 100
	fmt.Println(originalArr) // [1 2 3]
	fmt.Println(copyArr)     // [100 2 3]

	// 切片 (引用类型) - 赋值复制引用
	originalSlice := []int{1, 2, 3}
	copySlice := originalSlice // 指向同一个底层数组
	copySlice[0] = 100
	fmt.Println(originalSlice) // [100 2 3]
	fmt.Println(copySlice)     // [100 2 3] (共享修改)

	// 结构体包含切片字段
	type MyStruct struct {
		ID   int
		Data []string
	}

	original := MyStruct{ID: 1, Data: []string{"a", "b"}}
	shallowCopy := original                       // 结构体字段被复制，包括 Data 切片(指针)
	shallowCopy.ID = 2                            // 修改值类型字段，不影响 original
	shallowCopy.Data[0] = "x"                     // 修改共享切片的元素
	fmt.Println(original.ID, original.Data)       // 1 [x b]  -> ID没变，Data变了!
	fmt.Println(shallowCopy.ID, shallowCopy.Data) // 2 [x b]
}
