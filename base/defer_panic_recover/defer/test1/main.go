package main

import (
	"fmt"
	"os"
)

func main() {
	//fmt.Println(f3())
	testLocation()
}

/*
defer后边会接一个函数，但该函数不会立刻被执行，而是等到包含它的程序返回时(包含它的函数执行了return语句、运行到函数结尾自动返回、对应的goroutine panic）defer函数才会被执行。通常用于资源释放、打印日志、异常捕获等
*/
func testLocation() error {
	filename := "test.txt"
	_, err := os.Open(filename)
	if err != nil {
		// 由于defer是在后面定义的，所以这里如果return并不会调用后面的defer，千万切记
		// 针对panic也是同样的道理
		return err
	}
	/**
	 * 这里defer要写在err判断的后边而不是os.Open后边
	 * 如果资源没有获取成功，就没有必要对资源执行释放操作
	 * 如果err不为nil而执行资源执行释放操作，有可能导致panic
	 */
	defer func() {
		fmt.Println("tttttt")
	}()
	return nil
}

/*
如果包含defer函数的外层函数有返回值，而defer函数中可能会修改该返回值，最终导致外层函数实际的返回值可能与你想象的不一致，这里很容易踩坑
*/
func f1() (result int) { // 6
	defer func() {
		result++
	}()
	return 5
}

func f2() (r int) { // 5
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f3() (r int) { // 1
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

// 最重要的一点就是要明白，return xxx这一条语句并不是一条原子指令:
// 含有defer函数的外层函数，返回的过程是这样的：先给返回值赋值，然后调用defer函数，
// 最后才是返回到更上一级调用函数中，可以用一个简单的转换规则将return xxx改写成
/*
返回值 = xxx
调用defer函数(这里可能会有修改返回值的操作)
return 返回值
*/

/*
defer函数的参数值，是在申明defer时确定下来的

在defer函数申明时，对外部变量的引用是有两种方式：作为函数参数和作为闭包引用

作为函数参数，在defer申明时就把值传递给defer，并将值缓存起来，调用defer的时候使用缓存的值进行计算（如上边的例3）
而作为闭包引用，在defer函数执行时根据整个上下文确定当前的值
*/
func m1() {
	i := 0
	defer fmt.Println("a:", i)
	//闭包调用，将外部i传到闭包中进行计算，不会改变i的值，如上边的例3
	defer func(i int) {
		fmt.Println("b:", i)
	}(i)
	//闭包调用，捕获同作用域下的i进行计算
	defer func() {
		fmt.Println("c:", i)
	}()
	i++
}
