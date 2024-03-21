package main

import "fmt"

func countRangeSum(nums []int, lower int, upper int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	// 计算前缀和
	sum := make([]int64, len(nums))
	sum[0] = int64(nums[0])
	for i := 1; i < len(nums); i++ {
		sum[i] = sum[i-1] + int64(nums[i])
	}
	return process(sum, 0, len(sum)-1, int64(lower), int64(upper))
}

func process(sum []int64, L, R int, lower, upper int64) int {
	// merge无法处理的情况，因为merge是处理右组每个数的左组，而无法处理没有左组的情况，也就是L==R的情况只有一个数
	if L == R {
		if sum[L] >= lower && sum[L] <= upper {
			return 1
		}
		return 0
	}
	M := L + ((R - L) >> 1)
	return process(sum, L, M, lower, upper) + process(sum, M+1, R, lower, upper) + merge(sum, L, M, R, lower, upper)
}

func merge(arr []int64, L, M, R int, lower, upper int64) int {
	ans := 0
	// 定义窗口左右边界，窗口在左组，从左组的最左边开始向右扩
	windowL := L
	windowR := L
	// 遍历右组
	for i := M + 1; i <= R; i++ {
		// 计算右组每个数对应的区间范围
		min := arr[i] - upper
		max := arr[i] - lower
		// 下面这两个for循环就用到了顺序这个条件，之前我自己思考的时候认为不需要顺序，但是这里确实需要顺序才能拥有O(N)的时间复杂度
		// 窗口右边界不能超过max
		for windowR <= M && arr[windowR] <= max {
			windowR++
		}
		// 窗口左边界不能小于min
		for windowL <= M && arr[windowL] < min {
			windowL++
		}
		// 统计窗口内的数有多少个
		ans += windowR - windowL
	}
	help := make([]int64, R-L+1)
	i := 0
	p1 := L
	p2 := M + 1
	for p1 <= M && p2 <= R {
		if arr[p1] <= arr[p2] {
			help[i] = arr[p1]
			p1++
		} else {
			help[i] = arr[p2]
			p2++
		}
		i++
	}
	for p1 <= M {
		help[i] = arr[p1]
		i++
		p1++
	}
	for p2 <= R {
		help[i] = arr[p2]
		i++
		p2++
	}
	copy(arr[L:], help)
	return ans
}

func main() {
	// Test the function with an example
	nums := []int{-2, 5, -1}
	lower := -2
	upper := 2
	fmt.Println(countRangeSum(nums, lower, upper)) // Expected output: 3
}
