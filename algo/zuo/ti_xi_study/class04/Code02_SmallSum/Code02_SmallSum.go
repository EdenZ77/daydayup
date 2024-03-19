package main

import (
	"fmt"
	"math/rand"
	"time"
)

func smallSum(arr []int) int {
	if arr == nil || len(arr) < 2 {
		return 0
	}
	return process(arr, 0, len(arr)-1)
}

// arr[L..R]既要排好序，也要求小和返回
// 所有merge时，产生的小和，累加
// 左 排序   merge
// 右 排序  merge
// merge
func process(arr []int, l, r int) int {
	if l == r {
		return 0
	}
	mid := l + (r-l)/2
	return process(arr, l, mid) +
		process(arr, mid+1, r) +
		merge(arr, l, mid, r)
}

func merge(arr []int, L, m, r int) int {
	help := make([]int, r-L+1)
	i := 0
	p1 := L
	p2 := m + 1
	res := 0
	for p1 <= m && p2 <= r {
		// 如果arr[p1] < arr[p2]，产生小和
		if arr[p1] < arr[p2] {
			// 统计比arr[p2]大的数有多少个*arr[p2]
			res += (r - p2 + 1) * arr[p1]
			help[i] = arr[p1]
			p1++
		} else {
			// 如果arr[p1] >= arr[p2]，则先将arr[p2]放入help中
			help[i] = arr[p2]
			p2++
		}
		i++
	}
	for p1 <= m {
		help[i] = arr[p1]
		p1++
		i++
	}
	for p2 <= r {
		help[i] = arr[p2]
		p2++
		i++
	}
	for j := 0; j < len(help); j++ {
		arr[L+j] = help[j]
	}
	return res
}

// 接下来是测试用的函数

func comparator(arr []int) int {
	if arr == nil || len(arr) < 2 {
		return 0
	}
	res := 0
	for i := 1; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if arr[j] < arr[i] {
				res += arr[j]
			}
		}
	}
	return res
}

func generateRandomArray(maxSize, maxValue int) []int {
	arr := make([]int, rand.Intn(maxSize+1))
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(maxValue+1) - rand.Intn(maxValue)
	}
	return arr
}

func copyArray(arr []int) []int {
	if arr == nil {
		return nil
	}
	res := make([]int, len(arr))
	copy(res, arr)
	return res
}

func isEqual(arr1, arr2 []int) bool {
	if (arr1 == nil && arr2 != nil) || (arr1 != nil && arr2 == nil) {
		return false
	}
	if arr1 == nil && arr2 == nil {
		return true
	}
	if len(arr1) != len(arr2) {
		return false
	}
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

func printArray(arr []int) {
	if arr == nil {
		return
	}
	for _, num := range arr {
		fmt.Printf("%d ", num)
	}
	fmt.Println()
}

func main() {
	rand.Seed(time.Now().UnixNano())
	testTime := 500000
	maxSize := 100
	maxValue := 100
	succeed := true

	for i := 0; i < testTime; i++ {
		arr1 := generateRandomArray(maxSize, maxValue)
		arr2 := copyArray(arr1)
		if smallSum(arr1) != comparator(arr2) {
			succeed = false
			printArray(arr1)
			printArray(arr2)
			break
		}
	}
	if succeed {
		fmt.Println("Nice!")
	} else {
		fmt.Println("Fucking fucked!")
	}
}
