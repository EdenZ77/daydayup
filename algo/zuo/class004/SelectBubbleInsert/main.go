package main

import "fmt"

// 交换切片中i和j位置的元素
func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

// 选择排序
func selectionSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	for i := 0; i < len(arr)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		swap(arr, i, minIndex)
	}
}

// 冒泡排序
func bubbleSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	for end := len(arr) - 1; end > 0; end-- {
		for i := 0; i < end; i++ {
			if arr[i] > arr[i+1] {
				swap(arr, i, i+1)
			}
		}
	}
}

// 插入排序
func insertionSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	for i := 1; i < len(arr); i++ {
		for j := i - 1; j >= 0 && arr[j] > arr[j+1]; j-- {
			swap(arr, j, j+1)
		}
	}
}

func main() {
	arr := []int{5, 2, 9, 1, 5, 6}
	//selectionSort(arr)
	//bubbleSort(arr)
	insertionSort(arr)
	fmt.Println(arr) // 输出排序后结果
}
