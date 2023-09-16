package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	//testStack()

	//testQueue()

	//testMap()

	//testSort()

	//testMath()

	testDelete()

	//testCastType()

}

func testCastType() {
	s := "12345"           // s[0]类型是byte
	fmt.Println(s[0])      // 49
	num := int(s[0] - '0') // 1
	fmt.Println(num)       // 1
	str := string(s[0])    // "1"
	fmt.Println(str)
	b := byte(num + '0') // '1'
	fmt.Println(b)       // 49
}

func testDelete() {
	intA := make([]int, 0, 10)
	intA = append(intA, 4, 3, 2, 5, 7, 6, 8)
	fmt.Println(len(intA)) // 7
	fmt.Println(cap(intA)) // 10

	copy(intA[2:], intA[3:])
	intA = intA[:len(intA)-1]
	fmt.Println("测试删除a[2]")
	fmt.Println(intA)

	//   intA[2:]就是切片[2,5,7,6,8],intA[4:]就是切片[7,6,8],copy过程就是将intA[4:]覆盖了intA[2:]前三位，后面的6,8没有动
	//copy(intA[2:], intA[4:])
	//fmt.Println(intA) // [4 3 7 6 8 6 8]
	//intA = intA[:len(intA)-2]
	// 相当于删除了6、8。所以这种方式可以用来进行删除a[i]
	//fmt.Println(intA)      // [4 3 7 6 8]
	//fmt.Println(len(intA)) // 5
	//fmt.Println(cap(intA)) // 10

}

func testMath() {
	// int32 最大值，最小值
	fmt.Println(math.MaxInt32) //  1<<31 - 1
	fmt.Println(math.MinInt32) //  -1 << 31
}

func testSort() {
	// int排序
	intA := []int{4, 3, 2, 5, 7, 6}
	sort.Ints(intA)
	fmt.Println(intA)
	// 字符串排序
	stringB := []string{"zz", "aa", "ab", "22", "33", "2a"}
	sort.Strings(stringB)
	fmt.Println(stringB)
	// 自定义排序
	intB := []int{4, 3, 2, 5, 7, 6}
	sort.Slice(intB, func(i, j int) bool {
		return intB[i] > intB[j]
	})
	fmt.Println(intB)
}

func testMap() {
	// 创建
	m := make(map[string]int)
	// 设置kv
	m["hello"] = 1
	m["eden"] = 18
	// 删除k
	delete(m, "hello")
	// 遍历
	for k, v := range m {
		fmt.Println(k, v)
	}
}

func testQueue() {
	// 创建队列
	queue := make([]int, 0, 10)
	// enqueue 入队
	queue = append(queue, 10)
	queue = append(queue, 11)
	queue = append(queue, 12)
	// dequeue 出队
	v := queue[0]
	fmt.Println(v)
	queue = queue[1:]
	fmt.Println(queue)
	fmt.Println(len(queue)) // 2
	fmt.Println(cap(queue)) // 9

	// 判空
	//len(queue) == 0
}

func testStack() {
	// 创建栈 len=0, cap=0
	stack := make([]int, 0, 10)
	fmt.Println(len(stack)) // 0
	fmt.Println(cap(stack)) // 10

	// push压入
	stack = append(stack, 10)
	stack = append(stack, 11)
	stack = append(stack, 12)

	// pop弹出
	v := stack[len(stack)-1]
	fmt.Println(v)
	stack = stack[:len(stack)-1]
	fmt.Println(stack)
	fmt.Println(len(stack)) // 2
	fmt.Println(cap(stack)) // 10

	// 检查栈空
	//len(stack) == 0
}
