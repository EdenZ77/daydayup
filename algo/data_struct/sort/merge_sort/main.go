package main

import "fmt"

func mergeSort(nums []int) []int {
	// 就是拆分到单个元素
	if len(nums) < 2 {
		// 分治，两两拆分，一直拆到基础元素才向上递归。
		return nums
	}
	i := len(nums) / 2
	left := mergeSort(nums[0:i])
	// 左侧数据递归拆分
	right := mergeSort(nums[i:])
	// 右侧数据递归拆分
	result := merge(left, right)
	// 排序 & 合并
	return result
}

func merge(left, right []int) []int {
	result := make([]int, 0)
	i, j := 0, 0
	l, r := len(left), len(right)
	for i < l && j < r {
		// 这里保证了归并排序是稳定的排序算法，当两个值相同时，不会改变两个元素的位置
		if left[i] > right[j] {
			result = append(result, right[j])
			j++
		} else {
			result = append(result, left[i])
			i++
		}
	}
	// 判断那个子数组中有剩余的数据、
	// 将剩余的数据拷贝到result数组中
	if i < l {
		result = append(result, left[i:]...)
	} else {
		result = append(result, right[j:]...)
	}
	return result
}

func main() {
	//arr := []int{8, 9, 5, 7, 1, 2, 5, 7, 6, 3, 5, 4, 8, 1, 8, 5, 3, 5, 8, 4}
	arr := []int{5, 3, 6, 2, 7, 1, 9}
	result := mergeSort(arr)
	fmt.Println(result)
}
