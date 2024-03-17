package main

import (
	"errors"
	"fmt"
)

type TwoStacksQueue struct {
	stackPush []int
	stackPop  []int
}

func NewTwoStacksQueue() *TwoStacksQueue {
	return &TwoStacksQueue{
		stackPush: []int{},
		stackPop:  []int{},
	}
}

// 将 stackPush 中的数据倒入 stackPop
func (queue *TwoStacksQueue) pushToPop() {
	// Pop栈为空时才能倒入
	if len(queue.stackPop) == 0 {
		// 将Push栈中的数据倒入Pop栈
		for len(queue.stackPush) > 0 {
			value := queue.stackPush[len(queue.stackPush)-1]
			queue.stackPush = queue.stackPush[:len(queue.stackPush)-1]
			queue.stackPop = append(queue.stackPop, value)
		}
	}
}

// Add 添加元素
func (queue *TwoStacksQueue) Add(value int) {
	queue.stackPush = append(queue.stackPush, value)
	// 添加元素后，判断是否需要将 stackPush 中的数据倒入 stackPop
	queue.pushToPop()
}

// Poll 弹出元素
func (queue *TwoStacksQueue) Poll() (int, error) {
	if len(queue.stackPop) == 0 && len(queue.stackPush) == 0 {
		return 0, errors.New("queue is empty")
	}
	// 只要两个栈任意一个不为空，就判断是否需要将 stackPush 中的数据倒入 stackPop
	queue.pushToPop()
	// 从 stackPop 中弹出元素
	value := queue.stackPop[len(queue.stackPop)-1]
	queue.stackPop = queue.stackPop[:len(queue.stackPop)-1]
	return value, nil
}

// Peek 查看栈顶元素
func (queue *TwoStacksQueue) Peek() (int, error) {
	if len(queue.stackPop) == 0 && len(queue.stackPush) == 0 {
		return 0, errors.New("queue is empty")
	}
	queue.pushToPop()
	return queue.stackPop[len(queue.stackPop)-1], nil
}

func main() {
	queue := NewTwoStacksQueue()
	queue.Add(1)
	queue.Add(2)
	queue.Add(3)

	peek, _ := queue.Peek()
	fmt.Println("peek:", peek)

	poll, _ := queue.Poll()
	fmt.Println("poll:", poll)

	peek, _ = queue.Peek()
	fmt.Println("peek:", peek)

	poll, _ = queue.Poll()
	fmt.Println("poll:", poll)

	peek, _ = queue.Peek()
	fmt.Println("peek:", peek)

	poll, _ = queue.Poll()
	fmt.Println("poll:", poll)
}
