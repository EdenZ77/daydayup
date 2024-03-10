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
	for L < R { // 当 L..R 至少有两个数时
		mid = L + ((R - L) >> 1)
		if sortedArr[mid] == num {
			return true
		} else if sortedArr[mid] > num {
			R = mid - 1
		} else {
			L = mid + 1
		}
	}
	return sortedArr[L] == num
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
