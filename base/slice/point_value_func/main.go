package main

import "fmt"

/*
参考资料：
https://mp.weixin.qq.com/s?__biz=MjM5MDUwNTQwMQ==&mid=2257483743&idx=1&sn=af5059b90933bef5a7c9d491509d56d9&scene=19#wechat_redirect

# 传 slice 和 slice 指针有什么区别?
前面我们说到，slice 其实是一个结构体，包含了三个成员：len, cap, array。分别表示切片长度，容量，底层数据的地址。
当 slice 作为函数参数时，就是一个普通的结构体。其实很好理解：若直接传 slice，在调用者看来，实参 slice 并不会被函数中的操作改变；若传的是 slice 的指针，在调用者看来，是会被改变原 slice 的。

值的注意的是，不管传的是 slice 还是 slice 指针，如果改变了 slice 底层数组的数据，会反应到实参 slice 的底层数据。为什么能改变底层数组的数据？很好理解：底层数据在 slice 结构体里是一个指针，仅管 slice 结构体自身不会被改变，也就是说底层数据地址不会被改变。 但是通过指向底层数据的指针，可以改变切片的底层数据，没有问题。

通过 slice 的 array 字段就可以拿到数组的地址。在代码里，是直接通过类似 s[i]=10 这种操作改变 slice 底层数组元素值。



*/

//func main() {
//	s := []int{1, 1, 1}
//	f(s)
//	fmt.Println(s)
//}

func f(s []int) {
	// i只是一个副本，不能改变s中元素的值
	/*for _, i := range s {
	      i++
	  }
	*/

	for i := range s {
		s[i] += 1
	}
}

func myAppend(s []int) []int {
	// 这里 s 虽然改变了，但并不会影响外层函数的 s
	s = append(s, 100)
	return s
}

func myAppendPtr(s *[]int) {
	// 会改变外层 s 本身
	*s = append(*s, 100)
	return
}

func main() {
	s := []int{1, 1, 1}
	newS := myAppend(s)

	/*
		myAppend 函数里，虽然改变了 s，但它只是一个值传递，并不会影响外层的 s，因此第一行打印出来的结果仍然是 [1 1 1]。
		而 newS 是一个新的 slice，它是基于 s 得到的。因此它打印的是追加了一个 100 之后的结果： [1 1 1 100]。
	*/
	fmt.Println(s)
	fmt.Println(newS)

	s = newS

	/*
		最后，将 newS 赋值给了 s，s 这时才真正变成了一个新的slice。之后，再给 myAppendPtr 函数传入一个 s 指针，这回它真的被改变了：[1 1 1 100 100]。
	*/
	myAppendPtr(&s)
	fmt.Println(s)
}
