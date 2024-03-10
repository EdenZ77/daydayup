package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// exist 检查在已排序的数组 sortedArr 中是否存在数字 num
func exist(sortedArr []int, num int) bool {
	if sortedArr == nil || len(sortedArr) == 0 {
		return false
	}
	L := 0
	R := len(sortedArr) - 1
	var mid int
	// L..R
	for L <= R { // 当 L..R 至少有两个数时
		mid = L + ((R - L) >> 1)
		if sortedArr[mid] == num {
			return true
		} else if sortedArr[mid] > num {
			R = mid - 1
		} else {
			L = mid + 1
		}
	}
	// for循环写成 L < R
	// 当循环结束时，如果 num 没有找到，L 和 R 可能会收敛到同一个索引上，这个索引是 num 可能的位置（如果 num 存在于数组中）。
	// 注意：不可写成sortedArr[R] == num，如果num小于数组中最小的数时，R有可能=-1，导致下标越界
	//return sortedArr[L] == num

	// for循环写成 L <= R
	return false
}

// test 是用于测试在未排序数组中是否存在数字 num
func test(sortedArr []int, num int) bool {
	for _, cur := range sortedArr {
		if cur == num {
			return true
		}
	}
	return false
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

func main() {
	testTime := 500000
	maxSize := 10
	maxValue := 100
	succeed := true
	for i := 0; i < testTime; i++ {
		arr := generateRandomArray(maxSize, maxValue)
		sort.Ints(arr) // Go语言标准库中的排序函数
		value := rand.Intn(maxValue+1) - rand.Intn(maxValue)
		if test(arr, value) != exist(arr, value) {
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
