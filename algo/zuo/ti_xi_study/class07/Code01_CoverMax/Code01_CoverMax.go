package main

import (
	"container/heap"
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Line struct {
	start, end int
}

// ByStart 实现sort.Interface来排序线段的起始点
type ByStart []Line

func (a ByStart) Len() int           { return len(a) }
func (a ByStart) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByStart) Less(i, j int) bool { return a[i].start < a[j].start } // 按照start升序排序

// MinHeap 实现heap.Interface来构建小根堆，根据线段的结束点排序
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] } // 按照end升序排序, 小根堆
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// maxCover1 暴力解法
func maxCover1(lines [][]int) int {
	min, max := int(^uint(0)>>1), -int(^uint(0)>>1) // max int and min int
	// find the min and max value of the lines
	for _, line := range lines {
		if line[0] < min {
			min = line[0]
		}
		if line[1] > max {
			max = line[1]
		}
	}

	cover := 0
	// 以0.5为步长遍历(min, max)区间，统计每个点的覆盖线段数
	for p := float64(min) + 0.5; p < float64(max); p++ {
		cur := 0
		for _, line := range lines {
			// 如果当前点在某个线段的范围内，cur++
			if float64(line[0]) < p && float64(line[1]) > p {
				cur++
			}
		}
		// 更新最大覆盖线段数
		if cur > cover {
			cover = cur
		}
	}
	return cover
}

// maxCover2 使用小根堆来解决问题
func maxCover2(lines [][]int) int {
	lineObjects := make(ByStart, len(lines))
	// 构建Line对象数组
	for i, line := range lines {
		lineObjects[i] = Line{start: line[0], end: line[1]}
	}

	// Sort lines by start
	sort.Sort(lineObjects)

	// Initialize a priority queue (min-heap) for the end points of the lines
	minHeap := &MinHeap{}
	heap.Init(minHeap)
	maxCover := 0

	for _, line := range lineObjects {
		// 弹出所有小于等于当前线段起始点的线段结束点
		for minHeap.Len() > 0 && (*minHeap)[0] <= line.start {
			heap.Pop(minHeap)
		}
		// 将当前线段的结束点加入小根堆
		heap.Push(minHeap, line.end)
		// 更新最大覆盖线段数
		if minHeap.Len() > maxCover {
			maxCover = minHeap.Len()
		}
	}
	return maxCover
}

func generateLines(N, L, R int) [][]int {
	rand.Seed(time.Now().UnixNano())
	size := rand.Intn(N) + 1
	lines := make([][]int, size)
	for i := range lines {
		a := L + rand.Intn(R-L+1)
		b := L + rand.Intn(R-L+1)
		if a == b {
			b = a + 1
		}
		lines[i] = []int{min(a, b), max(a, b)}
	}
	return lines
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// 测试代码
	testTimes := 200000
	N, L, R := 100, 0, 200
	for i := 0; i < testTimes; i++ {
		lines := generateLines(N, L, R)
		ans1 := maxCover1(lines)
		ans2 := maxCover2(lines)
		if ans1 != ans2 {
			fmt.Println("Oops!")
		}
	}
	fmt.Println("test end")
}
