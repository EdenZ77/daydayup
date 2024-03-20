package main

import (
	"fmt"
	"math/rand"
	"time"
)

func reversePairs(arr []int) int {
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
	// [L....M] [M+1....R]
	ans := 0
	// 目前囊括进来的数，是从[M+1, windowR)
	windowR := m + 1
	// 遍历左组
	for i := L; i <= m; i++ {
		for windowR <= r && (arr[i] > arr[windowR]*2) {
			windowR++
		}
		// 统计右组窗口中，比arr[i]*2小的数有多少个
		ans += windowR - (m + 1)
	}
	// 上面的统计完了，下面是merge的过程，这两部分的时间复杂度都是O(N)，所以整体的时间复杂度就是O(N)
	help := make([]int, r-L+1)
	i := 0
	p1 := L
	p2 := m + 1
	for p1 <= m && p2 <= r {
		if arr[p1] <= arr[p2] {
			help[i] = arr[p1]
			p1++
		} else {
			help[i] = arr[p2]
			p2++
		}
		i++
	}
	for p1 <= m {
		help[i] = arr[p1]
		i++
		p1++
	}
	for p2 <= r {
		help[i] = arr[p2]
		i++
		p2++
	}
	for i, val := range help {
		arr[L+i] = val
	}
	return ans
}

func comparator(arr []int) int {
	ans := 0
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j]<<1 {
				ans++
			}
		}
	}
	return ans
}

func generateRandomArray(maxSize int, maxValue int) []int {
	rand.Seed(time.Now().UnixNano())
	length := rand.Intn(maxSize + 1)
	arr := make([]int, length)
	for i := range arr {
		arr[i] = rand.Intn(maxValue+1) - rand.Intn(maxValue+1)
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
	if (arr1 == nil) != (arr2 == nil) {
		return false
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
	for _, val := range arr {
		fmt.Printf("%d ", val)
	}
	fmt.Println()
}

func main() {
	testTime := 500000
	maxSize := 100
	maxValue := 100
	fmt.Println("测试开始")
	for i := 0; i < testTime; i++ {
		arr1 := generateRandomArray(maxSize, maxValue)
		arr2 := copyArray(arr1)
		if reversePairs(arr1) != comparator(arr2) {
			fmt.Println("Oops!")
			printArray(arr1)
			printArray(arr2)
			break
		}
	}
	fmt.Println("测试结束")
}
