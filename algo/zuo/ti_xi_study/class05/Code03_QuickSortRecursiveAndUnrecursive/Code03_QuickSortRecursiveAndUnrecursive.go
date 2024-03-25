package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 荷兰国旗问题
func netherlandsFlag(arr []int, L int, R int) (int, int) {
	if L > R {
		return -1, -1
	}
	if L == R {
		return L, R
	}
	less := L - 1
	more := R
	index := L
	for index < more {
		if arr[index] < arr[R] {
			less++
			arr[index], arr[less] = arr[less], arr[index]
			index++
		} else if arr[index] > arr[R] {
			more--
			arr[index], arr[more] = arr[more], arr[index]
		} else {
			index++
		}
	}
	arr[more], arr[R] = arr[R], arr[more]
	return less + 1, more
}

// 快排递归版本
func quickSort1(arr []int, L int, R int) {
	if L >= R {
		return
	}
	rand.Seed(time.Now().UnixNano())
	pivot := L + rand.Intn(R-L+1)
	arr[pivot], arr[R] = arr[R], arr[pivot]
	el, er := netherlandsFlag(arr, L, R)
	quickSort1(arr, L, el-1)
	quickSort1(arr, er+1, R)
}

type op struct {
	l int
	r int
}

// 快排非递归版本，使用栈
func quickSort2(arr []int) {
	if len(arr) < 2 {
		return
	}
	stack := make([]op, 0)
	stack = append(stack, op{0, len(arr) - 1})
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1] // Pop
		if cur.l < cur.r {
			rand.Seed(time.Now().UnixNano())
			pivot := cur.l + rand.Intn(cur.r-cur.l+1)
			arr[pivot], arr[cur.r] = arr[cur.r], arr[pivot]
			el, er := netherlandsFlag(arr, cur.l, cur.r)
			stack = append(stack, op{cur.l, el - 1})
			stack = append(stack, op{er + 1, cur.r})
		}
	}
}

// 快排非递归版本，使用队列
func quickSort3(arr []int) {
	if len(arr) < 2 {
		return
	}
	queue := make([]op, 0)
	queue = append(queue, op{0, len(arr) - 1})
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:] // Dequeue
		if cur.l < cur.r {
			rand.Seed(time.Now().UnixNano())
			pivot := cur.l + rand.Intn(cur.r-cur.l+1)
			arr[pivot], arr[cur.r] = arr[cur.r], arr[pivot]
			el, er := netherlandsFlag(arr, cur.l, cur.r)
			queue = append(queue, op{cur.l, el - 1})
			queue = append(queue, op{er + 1, cur.r})
		}
	}
}

// 测试
func main() {
	arr1 := []int{3, 1, 5, 7, 2, 4, 6, 8}
	arr2 := make([]int, len(arr1))
	copy(arr2, arr1)
	arr3 := make([]int, len(arr1))
	copy(arr3, arr1)

	quickSort1(arr1, 0, len(arr1)-1)
	fmt.Println("Sorted with quickSort1:", arr1)

	quickSort2(arr2)
	fmt.Println("Sorted with quickSort2:", arr2)

	quickSort3(arr3)
	fmt.Println("Sorted with quickSort3:", arr3)
}
