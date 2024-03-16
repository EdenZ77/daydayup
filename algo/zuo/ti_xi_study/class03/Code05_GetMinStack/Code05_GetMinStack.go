package main

import (
	"errors"
	"fmt"
)

type MyStack1 struct {
	// 数据栈
	stackData []int
	// 最小值栈
	stackMin []int
}

func (s *MyStack1) Push(newNum int) {
	// 如果最小值栈为空或者新元素小于等于最小值栈的栈顶元素，就将新元素压入最小值栈
	if len(s.stackMin) == 0 || newNum <= s.GetMin() {
		s.stackMin = append(s.stackMin, newNum)
	}
	s.stackData = append(s.stackData, newNum)
}

func (s *MyStack1) Pop() (int, error) {
	if len(s.stackData) == 0 {
		return 0, errors.New("your stack is empty")
	}
	value := s.stackData[len(s.stackData)-1]
	s.stackData = s.stackData[:len(s.stackData)-1]
	if value == s.GetMin() {
		s.stackMin = s.stackMin[:len(s.stackMin)-1]
	}
	return value, nil
}

func (s *MyStack1) GetMin() int {
	if len(s.stackMin) == 0 {
		panic("your stack is empty")
	}
	return s.stackMin[len(s.stackMin)-1]
}

type MyStack2 struct {
	stackData []int
	stackMin  []int
}

func (s *MyStack2) Push(newNum int) {
	min := s.GetMin()
	// 如果最小值栈为空或者新元素小于最小值栈的栈顶元素，就将新元素压入最小值栈
	if len(s.stackMin) == 0 || newNum < min {
		s.stackMin = append(s.stackMin, newNum)
	} else {
		// 否则将最小值栈的栈顶元素再次压入最小值栈
		s.stackMin = append(s.stackMin, min)
	}
	s.stackData = append(s.stackData, newNum)
}

func (s *MyStack2) Pop() (int, error) {
	if len(s.stackData) == 0 {
		return 0, errors.New("your stack is empty")
	}
	// 弹出数据栈和最小值栈的栈顶元素
	s.stackMin = s.stackMin[:len(s.stackMin)-1]
	value := s.stackData[len(s.stackData)-1]
	s.stackData = s.stackData[:len(s.stackData)-1]
	return value, nil
}

// GetMin 获取最小值栈的栈顶元素
func (s *MyStack2) GetMin() int {
	if len(s.stackMin) == 0 {
		panic("your stack is empty")
	}
	return s.stackMin[len(s.stackMin)-1]
}

func main() {
	stack1 := &MyStack1{}
	stack1.Push(3)
	fmt.Println(stack1.GetMin())
	stack1.Push(4)
	fmt.Println(stack1.GetMin())
	stack1.Push(1)
	fmt.Println(stack1.GetMin())
	value, _ := stack1.Pop()
	fmt.Println(value)
	fmt.Println(stack1.GetMin())

	fmt.Println("=============")

	stack2 := &MyStack2{}
	stack2.Push(3)
	fmt.Println(stack2.GetMin())
	stack2.Push(4)
	fmt.Println(stack2.GetMin())
	stack2.Push(1)
	fmt.Println(stack2.GetMin())
	value, _ = stack2.Pop()
	fmt.Println(value)
	fmt.Println(stack2.GetMin())
}
