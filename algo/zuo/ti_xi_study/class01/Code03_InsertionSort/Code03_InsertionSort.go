package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"time"
)

// 插入排序
func insertionSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	for i := 1; i < len(arr); i++ {
		for j := i - 1; j >= 0 && arr[j] > arr[j+1]; j-- {
			swap(arr, j, j+1)
		}
	}
}

// 使用异或操作符进行元素交换
// 注意：如果i和j是同一个位置，会导致该位置值变为0
func swap(arr []int, i, j int) {
	if i != j {
		arr[i] = arr[i] ^ arr[j]
		arr[j] = arr[i] ^ arr[j]
		arr[i] = arr[i] ^ arr[j]
	}
}

// comparator 为测试用的辅助比较函数
func comparator(arr []int) {
	sort.Ints(arr)
}

// generateRandomArray 生成随机数组
// 长度[0, maxSize]，元素值为-maxValue到maxValue
func generateRandomArray(maxSize int, maxValue int) []int {
	rand.Seed(time.Now().UnixNano())
	length := rand.Intn(maxSize + 1) // rand.Intn(n) 随机返回[0,n)之间的整数，其实数组长度没有必要=0
	arr := make([]int, length)
	for i := 0; i < length; i++ {
		arr[i] = rand.Intn(maxValue+1) - rand.Intn(maxValue+1)
	}
	return arr
}

// isEqual 检查两个数组是否相等
func isEqual(arr1, arr2 []int) bool {
	return reflect.DeepEqual(arr1, arr2)
}

// printArray 打印数组
func printArray(arr []int) {
	if arr == nil {
		return
	}
	for _, v := range arr {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}

func main() {
	testTime := 500000
	maxSize := 100  // 随机数组的长度0～100
	maxValue := 100 // 值：-100～100
	succeed := true
	for i := 0; i < testTime; i++ {
		arr := generateRandomArray(maxSize, maxValue)
		arr1 := make([]int, len(arr))
		arr2 := make([]int, len(arr))
		copy(arr1, arr)
		copy(arr2, arr)
		insertionSort(arr1)
		comparator(arr2)
		if !isEqual(arr1, arr2) {
			succeed = false
			printArray(arr)
			break
		}
	}
	if succeed {
		fmt.Println("Nice!")
	} else {
		fmt.Println("Fucking fucked!")
	}

	arr := generateRandomArray(maxSize, maxValue)
	printArray(arr)
	insertionSort(arr)
	printArray(arr)
}
