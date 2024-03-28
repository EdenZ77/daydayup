package main

import (
	"container/heap"
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// 排序数组，数组的任何一个数移动的距离不会超过k
func sortedArrDistanceLessK(arr []int, k int) {
	if k == 0 {
		return
	}

	// 创建一个优先队列（小根堆）
	minHeap := &IntHeap{}
	heap.Init(minHeap)
	index := 0

	// 首先将前k个元素放入堆中
	for ; index <= min(len(arr)-1, k-1); index++ {
		heap.Push(minHeap, arr[index])
	}

	// 遍历数组剩余的元素，并同时从堆中取出元素放回数组
	i := 0
	for ; index < len(arr); i, index = i+1, index+1 {
		heap.Push(minHeap, arr[index])
		arr[i] = heap.Pop(minHeap).(int)
	}

	// 将堆中剩余的元素取出放回数组
	for minHeap.Len() > 0 {
		arr[i] = heap.Pop(minHeap).(int)
		i++
	}
}

// IntHeap 是一个实现了heap.Interface的int类型的最小堆
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// min 返回两个整数中较小的一个
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// randomArrayNoMoveMoreK 生成一个随机数组，数组中任意数的移动距离不超过k
func randomArrayNoMoveMoreK(maxSize, maxValue, k int) []int {
	arr := make([]int, rand.Intn(maxSize+1))
	for i := range arr {
		arr[i] = rand.Intn(maxValue+1) - rand.Intn(maxValue)
	}

	// 先排序
	sort.Ints(arr)

	// 随机交换，保证每个数的移动距离不超过k
	isSwap := make([]bool, len(arr))
	for i := range arr {
		j := min(i+rand.Intn(k+1), len(arr)-1)
		if !isSwap[i] && !isSwap[j] {
			isSwap[i] = true
			isSwap[j] = true
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	return arr
}

// copyArray 复制一个整型数组
func copyArray(arr []int) []int {
	res := make([]int, len(arr))
	copy(res, arr)
	return res
}

// isEqual 判断两个数组是否相等
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

// printArray 打印一个整型数组
func printArray(arr []int) {
	for _, v := range arr {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func main() {
	fmt.Println("test begin")
	rand.Seed(time.Now().UnixNano())
	testTime := 500000
	maxSize := 100
	maxValue := 100
	succeed := true

	for i := 0; i < testTime; i++ {
		k := rand.Intn(maxSize) + 1
		arr := randomArrayNoMoveMoreK(maxSize, maxValue, k)
		arr1 := copyArray(arr)
		arr2 := copyArray(arr)

		sortedArrDistanceLessK(arr1, k)
		sort.Ints(arr2) // Go的sort.Ints实现了完整的排序

		if !isEqual(arr1, arr2) {
			succeed = false
			fmt.Println("K:", k)
			printArray(arr)
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
