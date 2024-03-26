package main

import (
	"fmt"
	"math/rand"
	"time"
)

// netherlandsFlag 解决荷兰国旗问题，返回等于区域的左右下标
func netherlandsFlag(arr []int, L, R int) (int, int) {
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
		if arr[index] == arr[R] {
			index++
		} else if arr[index] < arr[R] {
			swap(arr, index, less+1)
			index++
			less++
		} else {
			swap(arr, index, more-1)
			more--
		}
	}
	swap(arr, more, R)
	return less + 1, more
}

// swap 交换数组中两个元素的位置
func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

// quickSort1 快排递归版本
func quickSort1(arr []int) {
	if len(arr) < 2 {
		return
	}
	process(arr, 0, len(arr)-1)
}

// process 快排的处理过程
func process(arr []int, L, R int) {
	if L >= R {
		return
	}
	swap(arr, L+rand.Intn(R-L+1), R)
	equalL, equalR := netherlandsFlag(arr, L, R)
	process(arr, L, equalL-1)
	process(arr, equalR+1, R)
}

// Op 快排非递归版本需要的辅助类
type Op struct {
	l, r int
}

// quickSort2 快排非递归版本，使用栈来执行
func quickSort2(arr []int) {
	if len(arr) < 2 {
		return
	}
	N := len(arr)
	swap(arr, rand.Intn(N), N-1)
	equalL, equalR := netherlandsFlag(arr, 0, N-1)
	// 用栈这个结构来记录需要处理的区间，用户自己定义栈来代替递归的系统栈
	stack := []Op{{0, equalL - 1}, {equalR + 1, N - 1}}
	for len(stack) > 0 {
		current := stack[len(stack)-1]
		// 弹出栈顶元素，也就是切片的最后一个元素
		stack = stack[:len(stack)-1]
		// 如果当前区间的左边界小于右边界，说明当前区间还没有处理完
		if current.l < current.r {
			swap(arr, current.l+rand.Intn(current.r-current.l+1), current.r)
			// 使用荷兰国旗确定当前区间的等于区域
			equalL, equalR = netherlandsFlag(arr, current.l, current.r)
			// 将当前区间的左右两部分分别压入栈中, 顺序无所谓; 等待后面弹出栈顶元素时处理
			stack = append(stack, Op{current.l, equalL - 1}, Op{equalR + 1, current.r})
		}
		// 如果当前区间的左边界大于等于右边界，说明当前区间已经处理完，继续弹出栈顶元素，直至栈为空
	}
}

// quickSort3 快排非递归版本，使用队列来执行
func quickSort3(arr []int) {
	if len(arr) < 2 {
		return
	}
	N := len(arr)
	swap(arr, rand.Intn(N), N-1)
	equalL, equalR := netherlandsFlag(arr, 0, N-1)
	queue := []Op{{0, equalL - 1}, {equalR + 1, N - 1}}
	for len(queue) > 0 {
		current := queue[0]
		// 弹出队列头部元素，也就是切片的第一个元素
		queue = queue[1:]
		if current.l < current.r {
			swap(arr, current.l+rand.Intn(current.r-current.l+1), current.r)
			// 使用荷兰国旗确定当前区间的等于区域
			equalL, equalR = netherlandsFlag(arr, current.l, current.r)
			// 将当前区间的左右两部分分别压入队列中, 顺序无所谓; 等待后面弹出队列头部元素时处理
			queue = append(queue, Op{current.l, equalL - 1}, Op{equalR + 1, current.r})
		}
		// 如果当前区间的左边界大于等于右边界，说明当前区间已经处理完，继续弹出队列头部元素，直至队列为空
	}
}

// 可以发现，对于非递归版本的快排，使用栈和队列都可以，只是弹出元素的顺序不同，栈是后进先出，队列是先进先出，说明顺序无所谓，主要是记录需要处理的区间即可

// generateRandomArray 生成随机数组（用于测试）
func generateRandomArray(maxSize, maxValue int) []int {
	arr := make([]int, rand.Intn(maxSize+1))
	for i := range arr {
		arr[i] = rand.Intn(maxValue+1) - rand.Intn(maxValue)
	}
	return arr
}

// isEqual 对比两个数组（用于测试）
func isEqual(arr1, arr2 []int) bool {
	if (arr1 == nil && arr2 != nil) || (arr1 != nil && arr2 == nil) {
		return false
	}
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

func main() {
	rand.Seed(time.Now().UnixNano())
	testTime := 500000
	maxSize := 100
	maxValue := 100
	succeed := true
	fmt.Println("test begin")
	for i := 0; i < testTime; i++ {
		arr1 := generateRandomArray(maxSize, maxValue)
		arr2 := make([]int, len(arr1))
		arr3 := make([]int, len(arr1))
		copy(arr2, arr1)
		copy(arr3, arr1)
		quickSort1(arr1)
		quickSort2(arr2)
		quickSort3(arr3)
		if !isEqual(arr1, arr2) || !isEqual(arr1, arr3) {
			succeed = false
			break
		}
	}
	fmt.Println("test end")
	fmt.Println("测试", testTime, "组是否全部通过：", succeed)
}
