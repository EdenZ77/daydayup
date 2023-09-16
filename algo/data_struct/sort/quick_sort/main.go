package main

import "fmt"

func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	i, j := left, right
	// i从左边开始遍历，j从右边开始遍历。
	p := nums[i]
	// 设置基准数
	for i < j {
		for i < j && nums[i] < p {
			// 寻找一个比基准数大的数
			i++
		}
		for i < j && nums[j] > p {
			// 寻找一个比基准数小的数
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
		// 替换这两个数
	}
	nums[j] = p
	// 如果i和j相遇，就将基准数和j下标的值进行替换

	quickSort(nums, left, i-1)
	quickSort(nums, i+1, right)
}

func main() {
	a := []int{6, 5, 3, 1, 8, 7, 2, 4}
	quickSort(a, 0, len(a)-1)
	fmt.Println(a)
}
