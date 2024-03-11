package main

import (
	"fmt"
	"math/rand"
	"time"
)

// getLessIndex 返回数组中任意一个局部最小值的位置
func getLessIndex(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return -1
	}
	if len(arr) == 1 || arr[0] < arr[1] {
		return 0
	}
	if arr[len(arr)-1] < arr[len(arr)-2] {
		return len(arr) - 1
	}
	left := 1
	right := len(arr) - 2
	var mid int
	for left < right {
		mid = left + (right-left)/2
		if arr[mid] > arr[mid-1] {
			right = mid - 1
		} else if arr[mid] > arr[mid+1] {
			left = mid + 1
		} else {
			return mid
		}
	}
	return right // 也可以返回right
	//return left
}

// isRight 验证得到的结果是不是局部最小
func isRight(arr []int, index int) bool {
	if len(arr) <= 1 {
		return true
	}
	if index == 0 {
		return arr[index] < arr[index+1]
	}
	if index == len(arr)-1 {
		return arr[index] < arr[index-1]
	}
	return arr[index] < arr[index-1] && arr[index] < arr[index+1]
}

// generateRandomArray 生成相邻不相等的数组
func generateRandomArray(maxSize int, maxValue int) []int {
	rand.Seed(time.Now().UnixNano())
	size := rand.Intn(maxSize) + 1
	arr := make([]int, size)
	arr[0] = rand.Intn(2*maxValue) - maxValue
	for i := 1; i < size; i++ {
		// 使用for循环以确保arr[i]不等于arr[i-1]
		for {
			arr[i] = rand.Intn(2*maxValue) - maxValue
			if arr[i] != arr[i-1] {
				break
			}
		}
	}
	return arr
}

func main() {
	testTime := 500000
	maxSize := 30
	maxValue := 100
	fmt.Println("测试开始")
	for i := 0; i < testTime; i++ {
		arr := generateRandomArray(maxSize, maxValue)
		ans := getLessIndex(arr)
		if !isRight(arr, ans) {
			fmt.Println("出错了！")
			break
		}
	}
	fmt.Println("测试结束")
}
