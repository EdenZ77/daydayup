package main

import "fmt"

func main() {
	a := 16
	b := 603

	fmt.Println(a)
	fmt.Println(b)

	// 交换 a 和 b 的值
	a = a ^ b
	b = a ^ b
	a = a ^ b

	fmt.Println(a)
	fmt.Println(b)

	arr := []int{3, 1, 100}
	i := 0
	j := 0

	// 尝试交换相同位置的元素，理论上这里的交换是无意义的，因为 i 和 j 都是 0
	arr[i] = arr[i] ^ arr[j]
	arr[j] = arr[i] ^ arr[j]
	arr[i] = arr[i] ^ arr[j]

	fmt.Println(arr[i], ",", arr[j])

	fmt.Println(arr[0])
	fmt.Println(arr[2])

	// 使用函数交换 arr 中索引为 0 和 0 的元素
	swap(arr, 0, 0)

	fmt.Println(arr[0])
	fmt.Println(arr[2])
}

func swap(arr []int, i int, j int) {
	arr[i] = arr[i] ^ arr[j]
	arr[j] = arr[i] ^ arr[j]
	arr[i] = arr[i] ^ arr[j]
}
