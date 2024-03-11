package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// nearestIndex 在 arr 中找到小于或等于 value 的最右位置
func nearestIndex(arr []int, value int) int {
	if arr == nil || len(arr) == 0 {
		return -1
	}
	L := 0
	R := len(arr) - 1
	index := -1 // 记录最右的对应位置
	for L <= R {
		mid := L + ((R - L) >> 1)
		if arr[mid] <= value {
			index = mid
			L = mid + 1
		} else {
			R = mid - 1
		}
	}
	return index
}

// test 是用于测试的辅助函数，它通过线性搜索在数组中找到小于或等于 value 的最右位置
func test(arr []int, value int) int {
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] <= value {
			return i
		}
	}
	return -1
}

// generateRandomArray 生成一个随机数组
func generateRandomArray(maxSize int, maxValue int) []int {
	rand.Seed(time.Now().UnixNano())
	length := rand.Intn(maxSize + 1)
	arr := make([]int, length)
	for i := 0; i < length; i++ {
		arr[i] = rand.Intn(maxValue+1) - rand.Intn(maxValue)
	}
	return arr
}

// printArray 打印数组
func printArray(arr []int) {
	for _, v := range arr {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}

func main() {
	testTime := 500000
	maxSize := 10
	maxValue := 100
	succeed := true
	for i := 0; i < testTime; i++ {
		arr := generateRandomArray(maxSize, maxValue)
		sort.Ints(arr) // Go 语言的排序函数
		value := rand.Intn(maxValue+1) - rand.Intn(maxValue)
		if test(arr, value) != nearestIndex(arr, value) {
			printArray(arr)
			fmt.Println(value)
			fmt.Println(test(arr, value))
			fmt.Println(nearestIndex(arr, value))
			succeed = false
			break
		}
	}
	if succeed {
		fmt.Println("Nice!")
	} else {
		fmt.Println("Fucking fucked!")
	}
}
