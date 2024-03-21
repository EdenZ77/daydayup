package main

import (
	"fmt"
	"math/rand"
	"time"
)

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

// arr[L..R]上，以arr[R]位置的数做划分值
// <=X放左边  >X放右边，并且<=X区域的最后一个值等于X。
// 思路就是，arr[R]作为划分值，处理arr[L..R-1]，然后将arr[R]放到>=X区域的第一个位置，这样就完成了划分。
func partition(arr []int, L, R int) int {
	if L > R {
		return -1
	}
	if L == R {
		return L
	}
	lessEqual := L - 1
	index := L
	for index < R {
		if arr[index] <= arr[R] {
			lessEqual++
			swap(arr, index, lessEqual)
		}
		index++
	}
	swap(arr, lessEqual+1, R)
	return lessEqual + 1
}

// arr[L...R] 玩荷兰国旗问题的划分，以arr[R]做划分值
// <arr[R]放左边 ==arr[R]放中间 > arr[R]放右边
// 返回等于区域的左右边界
func netherlandsFlag(arr []int, L, R int) []int {
	if L > R {
		return []int{-1, -1}
	}
	if L == R {
		return []int{L, R}
	}
	less := L - 1 // < 区 右边界
	more := R     // > 区 左边界
	index := L
	for index < more { // 当前位置，不能和 >区的左边界撞上
		if arr[index] < arr[R] {
			less++
			swap(arr, index, less)
			index++
		} else if arr[index] == arr[R] {
			index++
		} else {
			more--
			swap(arr, index, more)
		}
	}
	swap(arr, more, R)
	return []int{less + 1, more}
}

func quickSort1(arr []int) {
	if len(arr) < 2 {
		return
	}
	process1(arr, 0, len(arr)-1)
}

func process1(arr []int, L, R int) {
	if L >= R {
		return
	}
	M := partition(arr, L, R)
	process1(arr, L, M-1)
	process1(arr, M+1, R)
}

func quickSort2(arr []int) {
	if len(arr) < 2 {
		return
	}
	process2(arr, 0, len(arr)-1)
}

func process2(arr []int, L, R int) {
	if L >= R {
		return
	}
	equalArea := netherlandsFlag(arr, L, R)
	process2(arr, L, equalArea[0]-1)
	process2(arr, equalArea[1]+1, R)
}

func quickSort3(arr []int) {
	if len(arr) < 2 {
		return
	}
	process3(arr, 0, len(arr)-1)
}

func process3(arr []int, L, R int) {
	if L >= R {
		return
	}
	rand.Seed(time.Now().UnixNano())
	swap(arr, L+rand.Intn(R-L+1), R)
	equalArea := netherlandsFlag(arr, L, R)
	process3(arr, L, equalArea[0]-1)
	process3(arr, equalArea[1]+1, R)
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

func main() {
	testTime := 500000
	maxSize := 100
	maxValue := 100
	succeed := true
	for i := 0; i < testTime; i++ {
		arr1 := generateRandomArray(maxSize, maxValue)
		arr2 := copyArray(arr1)
		arr3 := copyArray(arr1)
		quickSort1(arr1)
		quickSort2(arr2)
		quickSort3(arr3)
		if !isEqual(arr1, arr2) || !isEqual(arr2, arr3) {
			succeed = false
			break
		}
	}
	if succeed {
		fmt.Println("Nice!")
	} else {
		fmt.Println("Oops!")
	}
}
