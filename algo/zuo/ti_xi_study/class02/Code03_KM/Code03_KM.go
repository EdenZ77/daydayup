package main

import (
	"fmt"
	"math/rand"
	"time"
)

func test(arr []int, k int, m int) int {
	countMap := make(map[int]int)
	for _, num := range arr {
		countMap[num]++
	}
	for num, count := range countMap {
		if count == k {
			return num
		}
	}
	return 0
}

func onlyKTimes(arr []int, k int, m int) int {
	helpMap := make(map[int]int)
	value := 1
	for i := 0; i < 32; i++ {
		helpMap[value] = i
		value <<= 1
	}

	t := make([]int, 32)
	for _, num := range arr {
		for num != 0 {
			rightOne := num & (-num)
			t[helpMap[rightOne]]++
			num ^= rightOne
		}
	}

	ans := 0
	for i := 0; i < 32; i++ {
		if t[i]%m != 0 {
			ans |= (1 << i)
		}
	}

	return ans
}

func km(arr []int, k int, m int) int {
	help := make([]int, 32)
	// 统计数组中所有数据在0~31每个位置上面1的个数
	for _, num := range arr {
		for i := 0; i < 32; i++ {
			help[i] += (num >> i) & 1
		}
	}

	ans := 0
	// 遍历刚才获得数组0~31的每个位置
	for i := 0; i < 32; i++ {
		// 扩展1：不存在则返回-1
		//if help[i] % m == 0 {
		//	continue
		//}
		//if help[i] % m == k {
		//	ans |= 1 << i
		//} else {
		//	return -1
		//}

		help[i] %= m
		// 如果该位置不能够被m整除，表明出现k次的那个数字一定在该位置为1
		if help[i] != 0 {
			// 就用ans来 或 上所有位置为1，最后ans就是所求的那个出现k次的数字
			ans |= 1 << i
		}
	}
	return ans
}

func randomArray(maxKinds int, rangeNum int, k int, m int) []int {
	rand.Seed(time.Now().UnixNano())
	ktimeNum := randomNumber(rangeNum)
	times := k
	numKinds := rand.Intn(maxKinds) + 2
	arr := make([]int, times+(numKinds-1)*m)
	index := 0
	for ; index < times; index++ {
		arr[index] = ktimeNum
	}
	numKinds--
	set := make(map[int]bool)
	set[ktimeNum] = true
	for numKinds != 0 {
		curNum := 0
		for {
			curNum = randomNumber(rangeNum)
			if !set[curNum] {
				break
			}
		}
		set[curNum] = true
		numKinds--
		for i := 0; i < m; i++ {
			arr[index] = curNum
			index++
		}
	}
	for i := 0; i < len(arr); i++ {
		j := rand.Intn(len(arr))
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func randomNumber(rangeNum int) int {
	return rand.Intn(rangeNum+1) - rand.Intn(rangeNum+1)
}

func main() {
	kinds := 5
	rangeNum := 30
	testTime := 1000
	max1 := 9
	fmt.Println("测试开始")
	for i := 0; i < testTime; i++ {
		a := rand.Intn(max1) + 1
		b := rand.Intn(max1) + 1
		k := min(a, b)
		m := max(a, b)
		if k == m {
			m++
		}
		arr := randomArray(kinds, rangeNum, k, m)
		ans1 := test(arr, k, m)
		ans2 := onlyKTimes(arr, k, m)
		ans3 := km(arr, k, m)
		if ans1 != ans2 || ans1 != ans3 {
			fmt.Printf("出错了！ans1: %d, ans2: %d, ans3: %d\n", ans1, ans2, ans3)
		}
	}
	fmt.Println("测试结束")
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
