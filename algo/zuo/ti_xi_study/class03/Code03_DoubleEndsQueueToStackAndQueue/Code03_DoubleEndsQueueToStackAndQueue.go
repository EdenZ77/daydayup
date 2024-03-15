package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Node 这是一个双向链表的节点
type Node struct {
	value int
	last  *Node
	next  *Node
}

// DoubleEndsQueue 这是一个双端队列
type DoubleEndsQueue struct {
	head *Node
	tail *Node
}

// AddFromHead 从头部加入一个节点
func (dq *DoubleEndsQueue) AddFromHead(value int) {
	cur := &Node{value: value}
	if dq.head == nil {
		dq.head = cur
		dq.tail = cur
	} else {
		cur.next = dq.head
		// 把原来的头部的last指向新的头部
		dq.head.last = cur
		// 更新头部
		dq.head = cur
	}
}

// AddFromBottom 从尾部加入一个节点
func (dq *DoubleEndsQueue) AddFromBottom(value int) {
	cur := &Node{value: value}
	if dq.head == nil {
		dq.head = cur
		dq.tail = cur
	} else {
		cur.last = dq.tail
		// 把原来的尾部的next指向新的尾部
		dq.tail.next = cur
		// 更新尾部
		dq.tail = cur
	}
}

// PopFromHead 从头部弹出一个节点
func (dq *DoubleEndsQueue) PopFromHead() (value int, ok bool) {
	// 如果头部为空，说明队列为空
	if dq.head == nil {
		return 0, false
	}
	cur := dq.head
	// 如果头部和尾部是同一个节点
	if dq.head == dq.tail {
		dq.head = nil
		dq.tail = nil
	} else {
		// 更新头部
		dq.head = dq.head.next
		// 把原来的头部的next指向nil，这步其实可以不做，对于弹出的节点来说，next的指向是没有意义的
		cur.next = nil
		// 更新新的头部的last指向nil
		dq.head.last = nil
	}
	return cur.value, true
}

// PopFromBottom 从尾部弹出一个节点
func (dq *DoubleEndsQueue) PopFromBottom() (value int, ok bool) {
	if dq.head == nil {
		return 0, false
	}
	cur := dq.tail
	if dq.head == dq.tail {
		dq.head = nil
		dq.tail = nil
	} else {
		// 更新尾部
		dq.tail = dq.tail.last
		dq.tail.next = nil
		// 更新弹出节点的last指向其实是没有意义的
		cur.last = nil
	}
	return cur.value, true
}

func (dq *DoubleEndsQueue) IsEmpty() bool {
	return dq.head == nil
}

type MyStack struct {
	queue DoubleEndsQueue
}

// Push 向栈中加入一个元素
func (stack *MyStack) Push(value int) {
	stack.queue.AddFromHead(value)
}

// Pop 从栈中取出一个元素
func (stack *MyStack) Pop() (value int, ok bool) {
	return stack.queue.PopFromHead()
}

func (stack *MyStack) IsEmpty() bool {
	return stack.queue.IsEmpty()
}

type MyQueue struct {
	queue DoubleEndsQueue
}

// Push 向队列中加入一个元素
func (queue *MyQueue) Push(value int) {
	queue.queue.AddFromBottom(value)
}

// Poll 从队列中取出一个元素
func (queue *MyQueue) Poll() (value int, ok bool) {
	return queue.queue.PopFromHead()
}

func (queue *MyQueue) IsEmpty() bool {
	return queue.queue.IsEmpty()
}

func isEqual(o1, o2 int, o1Ok, o2Ok bool) bool {
	if !o1Ok && !o2Ok {
		return true
	}
	if o1Ok != o2Ok {
		return false
	}
	return o1 == o2
}

func randomValue(value int) int {
	return rand.Intn(value)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	oneTestDataNum := 100
	value := 10000
	testTimes := 100000
	for i := 0; i < testTimes; i++ {
		myStack := MyStack{}
		myQueue := MyQueue{}

		for j := 0; j < oneTestDataNum; j++ {
			nums := randomValue(value)
			if myStack.IsEmpty() {
				myStack.Push(nums)
			} else {
				if rand.Float64() < 0.5 {
					myStack.Push(nums)
				} else {
					stackPop, stackOk := myStack.Pop()
					if !isEqual(nums, stackPop, true, stackOk) {
						fmt.Println("Oops! Stack mismatch found")
						return
					}
				}
			}

			numq := randomValue(value)
			if myQueue.IsEmpty() {
				myQueue.Push(numq)
			} else {
				if rand.Float64() < 0.5 {
					myQueue.Push(numq)
				} else {
					queuePoll, queueOk := myQueue.Poll()
					if !isEqual(numq, queuePoll, true, queueOk) {
						fmt.Println("Oops! Queue mismatch found")
						return
					}
				}
			}
		}
	}
	fmt.Println("finish!")
}
