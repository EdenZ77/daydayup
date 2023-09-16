package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	a := []int{6, 5, 3, 1, 8, 7, 2, 4}
	array := sortArray(a)
	fmt.Println(array)
}

func sortArray(nums []int) []int {
	// 桶排序，基于哈希思想的外排稳定算法，空间换时间，时间O(n+k)
	// 相当于计数排序的改进版，服从均匀分布，先将数据分到有限数量的桶中，
	// 每个桶分别排序，最后将非空桶的数据拼接起来
	var bucket func(nums []int, bucketSize int) []int
	bucket = func(nums []int, bucketSize int) []int {
		if len(nums) < 2 {
			return nums
		}
		// 获取最大最小值
		minAndMax := func(nums []int) (min, max int) {
			minNum := math.MaxInt32
			maxNum := math.MinInt32
			for i := 0; i < len(nums); i++ {
				if nums[i] < minNum {
					minNum = nums[i]
				}
				if nums[i] > maxNum {
					maxNum = nums[i]
				}
			}
			return minNum, maxNum
		}
		min_, max_ := minAndMax(nums)
		// 定义桶
		// 构建计数桶
		bucketCount := (max_-min_)/bucketSize + 1
		buckets := make([][]int, bucketCount)
		for i := 0; i < bucketCount; i++ {
			buckets[i] = make([]int, 0)
		}
		// 装桶-排序过程
		for i := 0; i < len(nums); i++ {
			// 桶序号
			bucketNum := (nums[i] - min_) / bucketSize
			buckets[bucketNum] = append(buckets[bucketNum], nums[i])
		}
		// 桶中排序
		// 上述装桶完成，出桶填入元素组
		index := 0
		for _, bucket := range buckets {
			sort.Slice(bucket, func(i, j int) bool {
				return bucket[i] < bucket[j]
			})
			for _, num := range bucket {
				nums[index] = num
				index++
			}
		}
		return nums
	}
	// 定义桶中的数量
	var bucketSize int = 2
	return bucket(nums, bucketSize)
}
