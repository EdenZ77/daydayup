package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 递归方法实现归并排序
func mergeSort1(arr []int) {
	if len(arr) < 2 {
		return
	}
	process(arr, 0, len(arr)-1)
}

// process函数将arr[L..R]排有序
func process(arr []int, L int, R int) {
	if L == R {
		return
	}
	mid := L + (R-L)/2 // 避免溢出
	process(arr, L, mid)
	process(arr, mid+1, R)
	// merge函数合并两个有序数组
	merge(arr, L, mid, R)
}

// merge函数合并两个有序数组
func merge(arr []int, L int, M int, R int) {
	// 辅助数组
	help := make([]int, R-L+1)
	i, p1, p2 := 0, L, M+1
	for p1 <= M && p2 <= R {
		// 将arr[p1]和arr[p2]中较小的数放入help中
		if arr[p1] <= arr[p2] {
			help[i] = arr[p1]
			p1++
		} else {
			help[i] = arr[p2]
			p2++
		}
		// help数组下标右移
		i++
	}
	// 以下两个循环只有一个会执行
	// 如果p1 <= M，将arr[p1..M]放入help中
	for p1 <= M {
		help[i] = arr[p1]
		i++
		p1++
	}
	// 如果p2 <= R，将arr[p2..R]放入help中
	for p2 <= R {
		help[i] = arr[p2]
		i++
		p2++
	}
	// 将help数组拷贝回arr数组
	copy(arr[L:], help)
}

// 非递归方法实现归并排序
func mergeSort2(arr []int) {
	if len(arr) < 2 {
		return
	}
	N := len(arr)
	// 设置mergeSize为1，mergeSize表示每次合并的数组长度
	mergeSize := 1
	for mergeSize < N {
		// 当前左组的，第一个位置。每次步长改变后，都要从0开始
		L := 0
		// 每次步长改变后，重新开始计算L可能的位置，但是L肯定不能超过N
		for L < N {
			// 当前左组的，最后一个位置
			M := L + mergeSize - 1
			if M >= N {
				break
			}
			// 当前右组的最后一个位置
			R := min(M+mergeSize, N-1)
			merge(arr, L, M, R)
			// 下一次左组的第一个位置
			L = R + 1
		}
		// 例如：对于有9个数的数组，步长为1，2，4，此时4 !> 9/2，继续循环，步长变为8，然后整个数组有序。此时8 > 9/2，提前结束循环，避免多一次循环
		// 例如：对于有8个数的数组，步长为1，2，4，此时4 !> 8/2，继续循环，步长变为8，然后整个数组有序。此时8 > 8/2，结束循环。在这种情况下，还是会多判断一次for循环，但是mergeSize不会溢出
		// 还有一个目的是防止溢出，假如N无限接近int32的最大值，那么merge就可能接近最大值，此时mergeSize <<= 1就会溢出
		if mergeSize > N/2 {
			break
		}
		mergeSize <<= 1 // 等价于mergeSize *= 2
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 生成随机数组用于测试
func generateRandomArray(maxSize int, maxValue int) []int {
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, rand.Intn(maxSize+1))
	for i := range arr {
		arr[i] = rand.Intn(2*maxValue+1) - maxValue
	}
	return arr
}

// 复制数组用于测试
func copyArray(arr []int) []int {
	res := make([]int, len(arr))
	copy(res, arr)
	return res
}

// 测试数组是否相等
func isEqual(arr1, arr2 []int) bool {
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

// 打印数组
func printArray(arr []int) {
	for _, num := range arr {
		fmt.Printf("%d ", num)
	}
	fmt.Println()
}

// 测试
func main() {
	testTime := 500000
	maxSize := 100
	maxValue := 100
	fmt.Println("测试开始")
	for i := 0; i < testTime; i++ {
		arr1 := generateRandomArray(maxSize, maxValue)
		arr2 := copyArray(arr1)
		mergeSort1(arr1)
		mergeSort2(arr2)
		if !isEqual(arr1, arr2) {
			fmt.Println("出错了！")
			printArray(arr1)
			printArray(arr2)
			break
		}
	}
	fmt.Println("测试结束")
}
