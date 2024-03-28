package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type MyMaxHeap struct {
	heap []int
	// 堆的大小
	limit int
	// 堆中有多少个数
	size int
}

func NewMyMaxHeap(limit int) *MyMaxHeap {
	return &MyMaxHeap{
		heap:  make([]int, limit),
		limit: limit,
		size:  0,
	}
}

// IsEmpty 返回堆是否为空
func (h *MyMaxHeap) IsEmpty() bool {
	return h.size == 0
}

// IsFull 返回堆是否已满
func (h *MyMaxHeap) IsFull() bool {
	return h.size == h.limit
}

// Push 把值 value 加入到大根堆中
func (h *MyMaxHeap) Push(value int) error {
	if h.IsFull() {
		return errors.New("heap is full")
	}
	h.heap[h.size] = value
	h.heapInsert(h.size)
	h.size++
	return nil
}

// Pop 用户此时，让你返回最大值，并且在大根堆中，把最大值删掉
// 剩下的数，依然保持大根堆组织
func (h *MyMaxHeap) Pop() (int, error) {
	if h.IsEmpty() {
		return 0, errors.New("heap is empty")
	}
	ans := h.heap[0]
	h.size--
	h.heap[0] = h.heap[h.size]
	h.heapify(0)
	return ans, nil
}

// 新加进来的数，现在停在了index位置，请依次往上移动，
// 移动到0位置，或者干不掉自己的父亲了，停！
func (h *MyMaxHeap) heapInsert(index int) {
	// 这个判断处理了两种情况：
	// 1. index == 0，h.heap[0] !> h.heap[(0-1)/2] = h.heap[0]，移动到0位置，停！
	// 2. 或者干不掉自己的父亲了，停！
	for h.heap[index] > h.heap[(index-1)/2] {
		// 交换index和index的父节点的值
		h.swap(index, (index-1)/2)
		// index更新为index的父节点
		index = (index - 1) / 2
	}
}

// 从index位置，往下看，不断的下沉
// 停：较大的孩子都不再比index位置的数大，或者已经没孩子了
func (h *MyMaxHeap) heapify(index int) {
	// 左孩子的下标
	left := 2*index + 1
	// 如果有左孩子，有没有右孩子，可能有可能没有！
	for left < h.size {
		// 把较大孩子的下标，给largest
		largest := left
		// 如果有右孩子，且右孩子的值大于左孩子的值
		if left+1 < h.size && h.heap[left+1] > h.heap[left] {
			largest = left + 1
		}
		if h.heap[largest] <= h.heap[index] {
			break
		}
		// index和较大孩子，要互换
		h.swap(largest, index)
		// index更新为较大孩子的下标
		index = largest
		left = 2*index + 1
	}
}

func (h *MyMaxHeap) swap(i, j int) {
	h.heap[i], h.heap[j] = h.heap[j], h.heap[i]
}

func main() {
	// ... Code for testing MyMaxHeap as in Java main method
	// You can use the same logic for testing as in the Java code.
	// Here, we'll just demonstrate basic operations.
	rand.Seed(time.Now().UnixNano())

	mh := NewMyMaxHeap(10)
	for i := 0; i < 5; i++ {
		value := rand.Intn(100)
		fmt.Printf("Pushing %d\n", value)
		mh.Push(value)
	}

	for !mh.IsEmpty() {
		value, err := mh.Pop()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Popped %d\n", value)
		}
	}
}
