package main

import (
	"fmt"
	"math/rand"
	"time"
)

func reversePairNumber(arr []int) int {
	if arr == nil || len(arr) < 2 {
		return 0
	}
	return process(arr, 0, len(arr)-1)
}

func process(arr []int, l, r int) int {
	if l == r {
		return 0
	}
	mid := l + (r-l)/2
	return process(arr, l, mid) + process(arr, mid+1, r) + merge(arr, l, mid, r)
}

func merge(arr []int, L, m, r int) int {
	help := make([]int, r-L+1)
	// 从右向左进行统计逆序对
	i := len(help) - 1
	p1 := m
	p2 := r
	res := 0
	for p1 >= L && p2 > m {
		if arr[p1] > arr[p2] {
			res += p2 - m
			help[i] = arr[p1]
			p1--
		} else {
			help[i] = arr[p2]
			p2--
		}
		i--
	}
	for p1 >= L {
		help[i] = arr[p1]
		p1--
		i--
	}
	for p2 > m {
		help[i] = arr[p2]
		p2--
		i--
	}
	for j := 0; j < len(help); j++ {
		arr[L+j] = help[j]
	}
	return res
}

// Test functions below

func comparator(arr []int) int {
	ans := 0
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				ans++
			}
		}
	}
	return ans
}

func generateRandomArray(maxSize, maxValue int) []int {
	arr := make([]int, rand.Intn(maxSize+1))
	for i := range arr {
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
	for i := range arr1 {
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
	fmt.Println("测试开始")
	for i := 0; i < testTime; i++ {
		arr1 := generateRandomArray(maxSize, maxValue)
		arr2 := copyArray(arr1)
		if reversePairNumber(arr1) != comparator(arr2) {
			fmt.Println("Oops!")
			printArray(arr1)
			printArray(arr2)
			break
		}
	}
	fmt.Println("测试结束")
}
