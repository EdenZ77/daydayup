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

	// 首先将前k+1个元素放入堆中  0...K
	for ; index <= min(len(arr)-1, k); index++ {
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
/*
这个 Interface 是 Go 语言标准库 sort 包中定义的一个接口，它用于提供一种通用的方式来对任何实现了这个接口的集合进行排序。这个接口定义了三个方法，Len(), Less(i, j int) 和 Swap(i, j int)，这些方法提供了排序算法所需的基本操作。

Len() 方法返回集合中元素的数量。
Less(i, j int) 方法报告索引 i 处的元素是否应该在索引 j 处的元素之前。此方法用于定义集合中元素的排序顺序。如果 Less(i, j) 和 Less(j, i) 都返回 false，则认为索引 i 和 j 处的元素是相等的。Sort 方法可以自由地决定相等元素在最终结果中的顺序，而 Stable 方法则会保持相等元素在输入中的原始顺序。
Swap(i, j int) 方法交换索引 i 和 j 处的元素。

Less 方法必须描述一个传递性排序（transitive ordering），即如果 Less(i, j) 和 Less(j, k) 均为 true，则 Less(i, k) 也必须为 true；反之，如果 Less(i, j) 和 Less(j, k) 均为 false，则 Less(i, k) 也必须为 false。

实际上，任何需要排序的数据结构，只要实现了这三个方法，就可以使用 sort 包中的 Sort 函数进行排序。例如，如果你有一个切片类型的数据结构，并且你希望能够对其进行排序，你就可以定义这三个方法，然后使用 sort.Sort 方法进行排序。
*/
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] } // 小根堆
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push 将一个新元素 x 添加到堆中
func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

// Pop 移除并返回堆中的最小元素
func (h *IntHeap) Pop() any {
	// 首先，它保存了堆的当前状态，然后找到最后一个元素（位于切片的末尾），并将切片缩减以移除这个元素。最后，它返回被移除的元素。
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
