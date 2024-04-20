package main

import (
	"fmt"
)

type HeapGreater struct {
	heap     []int
	indexMap map[int]int
	comp     func(a, b int) bool
}

func NewHeapGreater(comp func(a, b int) bool) *HeapGreater {
	return &HeapGreater{
		heap:     make([]int, 0),
		indexMap: make(map[int]int),
		comp:     comp,
	}
}

func (h *HeapGreater) isEmpty() bool {
	return len(h.heap) == 0
}

func (h *HeapGreater) size() int {
	return len(h.heap)
}

func (h *HeapGreater) contains(index int) bool {
	_, exists := h.indexMap[index]
	return exists
}

func (h *HeapGreater) peek() int {
	return h.heap[0]
}

func (h *HeapGreater) push(index int) {
	h.heap = append(h.heap, index)
	h.indexMap[index] = len(h.heap) - 1
	h.heapInsert(len(h.heap) - 1)
}

func (h *HeapGreater) pop() int {
	ansIndex := h.heap[0]
	h.swap(0, len(h.heap)-1)
	delete(h.indexMap, ansIndex)
	h.heap = h.heap[:len(h.heap)-1]
	h.heapify(0)
	return ansIndex
}

func (h *HeapGreater) remove(index int) {
	replaceIndex := h.heap[len(h.heap)-1]
	if _, exists := h.indexMap[index]; !exists {
		return
	}
	heapIndex := h.indexMap[index]
	h.swap(heapIndex, len(h.heap)-1)
	delete(h.indexMap, index)
	h.heap = h.heap[:len(h.heap)-1]
	if index != replaceIndex {
		h.resign(replaceIndex)
	}
}

func (h *HeapGreater) resign(index int) {
	h.heapInsert(h.indexMap[index])
	h.heapify(h.indexMap[index])
}

func (h *HeapGreater) getAllElements() []int {
	return h.heap
}

func (h *HeapGreater) heapInsert(index int) {
	for index > 0 && h.comp(h.heap[index], h.heap[(index-1)/2]) {
		h.swap(index, (index-1)/2)
		index = (index - 1) / 2
	}
}

func (h *HeapGreater) heapify(index int) {
	left := 2*index + 1
	for left < len(h.heap) {
		smallest := left
		if right := left + 1; right < len(h.heap) && h.comp(h.heap[right], h.heap[left]) {
			smallest = right
		}
		if h.comp(h.heap[index], h.heap[smallest]) {
			break
		}
		h.swap(index, smallest)
		index = smallest
		left = 2*index + 1
	}
}

func (h *HeapGreater) swap(i, j int) {
	h.indexMap[h.heap[i]], h.indexMap[h.heap[j]] = j, i
	h.heap[i], h.heap[j] = h.heap[j], h.heap[i]
}

func main() {
	// 测试代码，使用整数代替Student类型，我们将这些整数看作是数组下标
	arr := []int{17, 10, 29, 33, 54, 93}
	heap := NewHeapGreater(func(i, j int) bool {
		return arr[i] < arr[j] // 最小堆
	})

	// 将所有下标放入堆中
	for i := range arr {
		heap.push(i)
	}

	// 修改一个值并重新调整堆
	arr[4] = 4
	heap.resign(4) // 传入下标4，而不是值

	// 弹出所有元素
	for !heap.isEmpty() {
		index := heap.pop()
		fmt.Printf("下标: %d, 数字: %d\n", index, arr[index])
	}
}
