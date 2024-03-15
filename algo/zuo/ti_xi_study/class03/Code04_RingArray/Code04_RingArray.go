package main

import (
	"errors"
	"fmt"
)

// MyQueue 是一个使用固定大小数组实现的环形队列
type MyQueue struct {
	arr   []int
	pushi int // 指向队列末尾, end
	polli int // 指向队列开头, begin
	size  int
	limit int
}

// NewMyQueue 创建一个新的环形队列
func NewMyQueue(limit int) *MyQueue {
	return &MyQueue{
		arr:   make([]int, limit),
		pushi: 0,
		polli: 0,
		size:  0,
		limit: limit,
	}
}

// Push 向队列中加入一个元素
func (q *MyQueue) Push(value int) error {
	if q.size == q.limit {
		return errors.New("队列满了，不能再加了")
	}
	q.size++
	q.arr[q.pushi] = value
	q.pushi = q.nextIndex(q.pushi)
	return nil
}

// Pop 从队列中取出一个元素
func (q *MyQueue) Pop() (int, error) {
	if q.size == 0 {
		return 0, errors.New("队列空了，不能再拿了")
	}
	q.size--
	ans := q.arr[q.polli]
	q.polli = q.nextIndex(q.polli)
	return ans, nil
}

// IsEmpty 检查队列是否为空
func (q *MyQueue) IsEmpty() bool {
	return q.size == 0
}

// nextIndex 计算环形队列的下一个索引位置
func (q *MyQueue) nextIndex(i int) int {
	if i < q.limit-1 {
		return i + 1
	}
	return 0
}

func main() {
	queue := NewMyQueue(5) // 创建一个大小为5的队列
	err := queue.Push(1)
	if err != nil {
		fmt.Println(err)
	}
	err = queue.Push(2)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(queue.Pop()) // 应该返回 1
	fmt.Println(queue.Pop()) // 应该返回 2

	if queue.IsEmpty() {
		fmt.Println("队列现在是空的")
	}
}
