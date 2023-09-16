package main

import "fmt"

// 参考资料：https://mp.weixin.qq.com/s?__biz=MjM5MDUwNTQwMQ==&mid=2257483743&idx=1&sn=af5059b90933bef5a7c9d491509d56d9&scene=19#wechat_redirect

func main() {
	test1()
}

func test0() {
	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(len(data)) // 10
	fmt.Println(cap(data)) // 10

	/*
		对 data 使用3个索引值，截取出新的slice。这里 data 可以是数组或者 slice。
		low 是最低索引值，这里是闭区间，也就是说第一个元素是 data 位于 low 索引处的元素；
		而 high 和 max 则是开区间，表示最后一个元素只能是索引 high-1 处的元素，而最大容量则只能是索引 max-1 处的元素。
	*/
	slice := data[2:4:6]    // data[low, high, max]
	fmt.Println(slice)      // [2 3]
	fmt.Println(len(slice)) // 2
	fmt.Println(cap(slice)) // 4
	fmt.Println("=========")
	slice1 := data[2:4]      // data[low, high, max]
	fmt.Println(slice1)      // [2 3]
	fmt.Println(len(slice1)) // 2
	fmt.Println(cap(slice1)) // 8
	fmt.Println("=========")
	slice2 := data[2:2]      // data[low, high, max]
	fmt.Println(slice2)      // []
	fmt.Println(len(slice2)) // 0
	fmt.Println(cap(slice2)) // 8
}

func test1() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := slice[2:5]     // s1 从 slice 索引2（闭区间）到索引5（开区间，元素真正取到索引4），长度为3，容量默认到数组结尾，为8。
	s2 := s1[2:6:7]      //  s2 从 s1 的索引2（闭区间）到索引6（开区间，元素真正取到索引5），容量到索引7（开区间，真正到索引6），为5。
	fmt.Println(s1)      // [2 3 4]
	fmt.Println(s2)      // [4 5 6 7]
	fmt.Println(len(s2)) // 4
	fmt.Println(cap(s2)) // 5

	s2 = append(s2, 100)
	s2 = append(s2, 200)

	s1[2] = 20

	fmt.Println(s1)    // [2 3 20]
	fmt.Println(s2)    // [4 5 6 7 100 200]
	fmt.Println(slice) // [0 1 2 3 20 5 6 7 100 9]
}
