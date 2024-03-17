package main

import (
	"fmt"
	"math/rand"
	"time"
)

type TwoQueueStack struct {
	queue []int
	help  []int
}

func NewTwoQueueStack() *TwoQueueStack {
	return &TwoQueueStack{
		queue: make([]int, 0),
		help:  make([]int, 0),
	}
}

// Push 添加元素
func (s *TwoQueueStack) Push(value int) {
	s.queue = append(s.queue, value)
}

// Poll 弹出元素
func (s *TwoQueueStack) Poll() int {
	if len(s.queue) == 0 {
		panic("stack is empty")
	}
	// 将 queue 中的数据倒入 help, 直到 queue 中只剩一个元素
	for len(s.queue) > 1 {
		val := s.queue[0]
		s.queue = s.queue[1:]
		s.help = append(s.help, val)
	}
	// queue 中只剩一个元素，即为要弹出的元素
	ans := s.queue[0]
	// 重新赋值 queue 和 help
	s.queue = s.help
	s.help = []int{}
	return ans
}

func (s *TwoQueueStack) Peek() int {
	if len(s.queue) == 0 {
		panic("stack is empty")
	}
	for len(s.queue) > 1 {
		val := s.queue[0]
		s.queue = s.queue[1:]
		s.help = append(s.help, val)
	}
	ans := s.queue[0]
	// 因为这里只是查看栈顶元素，所以需要将弹出的元素重新放回 queue 中
	s.help = append(s.help, ans)
	s.queue = s.help
	s.help = []int{}
	return ans
}

func (s *TwoQueueStack) IsEmpty() bool {
	return len(s.queue) == 0
}

func main() {
	fmt.Println("test begin")
	myStack := NewTwoQueueStack()
	var testStack []int // 使用切片模拟真实的栈行为
	rand.Seed(time.Now().UnixNano())
	testTime := 1000000
	max := 1000000
	for i := 0; i < testTime; i++ {
		if myStack.IsEmpty() {
			if len(testStack) != 0 {
				fmt.Println("Oops! Test stack should be empty.")
			}
			num := rand.Intn(max)
			myStack.Push(num)
			testStack = append(testStack, num)
		} else {
			decision := rand.Float64()
			if decision < 0.25 {
				num := rand.Intn(max)
				myStack.Push(num)
				testStack = append(testStack, num)
			} else if decision < 0.5 {
				if myStack.Peek() != testStack[len(testStack)-1] {
					fmt.Println("Oops! Peek value mismatch.")
				}
			} else if decision < 0.75 {
				if myStack.Poll() != testStack[len(testStack)-1] {
					fmt.Println("Oops! Poll value mismatch.")
				}
				testStack = testStack[:len(testStack)-1]
			} else {
				if myStack.IsEmpty() != (len(testStack) == 0) {
					fmt.Println("Oops! IsEmpty status mismatch.")
				}
			}
		}
	}

	fmt.Println("test finish!")
}
