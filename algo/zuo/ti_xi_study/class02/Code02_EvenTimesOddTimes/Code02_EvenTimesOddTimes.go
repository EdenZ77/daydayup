package main

import "fmt"

// printOddTimesNum1 打印出现奇数次的数（假设只有一个数出现奇数次）
func printOddTimesNum1(arr []int) {
	eor := 0
	for i := 0; i < len(arr); i++ {
		eor ^= arr[i]
	}
	fmt.Println(eor)
}

// printOddTimesNum2 打印两个出现奇数次的数
func printOddTimesNum2(arr []int) {
	eor := 0
	for i := 0; i < len(arr); i++ {
		eor ^= arr[i]
	}

	rightOne := eor & (-eor)

	onlyOne := 0
	for i := 0; i < len(arr); i++ {
		if (arr[i] & rightOne) != 0 {
			onlyOne ^= arr[i]
		}
	}
	fmt.Println(onlyOne, eor^onlyOne)
}

// bit1counts 计算一个整数二进制表示中1的个数
func bit1counts(N int) int {
	count := 0
	for N != 0 {
		rightOne := N & (-N)
		count++
		N ^= rightOne
	}
	return count
}

func main() {
	a := 5
	b := 7

	// 交换a和b的值
	a = a ^ b
	b = a ^ b
	a = a ^ b

	fmt.Println(a)
	fmt.Println(b)

	arr1 := []int{3, 3, 2, 3, 1, 1, 1, 3, 1, 1, 1}
	printOddTimesNum1(arr1)

	arr2 := []int{4, 3, 4, 2, 2, 2, 4, 1, 1, 1, 3, 3, 1, 1, 1, 4, 2, 2}
	printOddTimesNum2(arr2)
}
