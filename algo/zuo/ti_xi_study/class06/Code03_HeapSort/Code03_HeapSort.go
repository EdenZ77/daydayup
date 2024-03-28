package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// heapSort performs heap sort on the array.
func heapSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	// 使用 heapInsert 方法建立堆
	// 时间复杂度为 O(NlogN)
	//for i := 0; i < len(arr); i++ { // O(N)
	//	heapInsert(arr, i) // O(logN)
	//}
	// 为什么采用下面的方式建堆时间复杂度为 O(N)？
	// 因为这是从下往上建堆，下面层数的节点数更多，且浅；上面层数的节点数更少，且深。
	for i := len(arr) - 1; i >= 0; i-- {
		heapify(arr, i, len(arr))
	}
	heapSize := len(arr)
	// 将堆顶元素与最后一个元素交换，然后重新调整堆
	swap(arr, 0, heapSize-1)
	heapSize--
	for heapSize > 0 {
		heapify(arr, 0, heapSize)
		swap(arr, 0, heapSize-1)
		heapSize--
	}
}

// heapInsert 是插入操作，将索引为 index 的新元素插入堆中。
func heapInsert(arr []int, index int) {
	for arr[index] > arr[(index-1)/2] {
		swap(arr, index, (index-1)/2)
		index = (index - 1) / 2
	}
}

// heapify makes the array satisfy heap properties starting from the index down to the leaves.
func heapify(arr []int, index, heapSize int) {
	left := index*2 + 1
	for left < heapSize {
		largest := left
		if left+1 < heapSize && arr[left+1] > arr[left] {
			largest = left + 1
		}
		if arr[largest] <= arr[index] {
			largest = index
		}
		if largest == index {
			break
		}
		swap(arr, largest, index)
		index = largest
		left = index*2 + 1
	}
}

// swap exchanges elements at indices i and j.
func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

// generateRandomArray creates an array of random integers.
func generateRandomArray(maxSize, maxValue int) []int {
	rand.Seed(time.Now().UnixNano())
	size := rand.Intn(maxSize + 1)
	arr := make([]int, size)
	for i := range arr {
		arr[i] = rand.Intn(maxValue+1) - rand.Intn(maxValue)
	}
	return arr
}

// isEqual checks if two arrays are equal.
func isEqual(arr1, arr2 []int) bool {
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

// printArray prints elements of the array.
func printArray(arr []int) {
	for _, num := range arr {
		fmt.Printf("%d ", num)
	}
	fmt.Println()
}

// main is the entry point for testing heap sort implementation.
func main() {
	testTime := 500000
	maxSize := 100
	maxValue := 100
	succeed := true
	for i := 0; i < testTime; i++ {
		arr1 := generateRandomArray(maxSize, maxValue)
		arr2 := make([]int, len(arr1))
		copy(arr2, arr1)
		heapSort(arr1)
		sort.Ints(arr2)
		if !isEqual(arr1, arr2) {
			succeed = false
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
	heapSort(arr)
	printArray(arr)
}
