package main

import (
	"fmt"
	"reflect"
)

/*
如果需要处理三个 chan，你就可以再添加一个 case clause，用它来处理第三个 chan。
可是，如果要处理 100 个 chan 呢？一万个 chan 呢？或者是，chan 的数量在编译的时候是不定的，
在运行的时候需要处理一个 slice of chan，这个时候，也没有办法在编译前写成字面意义的 select。那该怎么办？

通过 reflect.Select 函数，你可以将一组运行时的 case clause 传入，当作参数执行。
Go 的 select 是伪随机的，它可以在执行的 case 中随机选择一个 case，并把选择的这个 case 的索引（chosen）返回，
如果没有可用的 case 返回，会返回一个 bool 类型的返回值，这个返回值用来表示是否有 case 成功被选择。
如果是 recv case，还会返回接收的元素。Select 的方法签名如下：

func Select(cases []SelectCase) (chosen int, recv Value, recvOK bool)

反射存在的意义：将难以编码实现的代码通过反射交给运行时生成
*/

func main() {
	var ch1 = make(chan int, 10)
	var ch2 = make(chan int, 10)

	// 创建SelectCase
	var cases = createCases(ch1, ch2)

	// 执行10次select
	for i := 0; i < 10; i++ {
		chosen, recv, ok := reflect.Select(cases)
		if recv.IsValid() { // recv case
			fmt.Println("recv:", cases[chosen].Dir, recv, ok)
		} else { // send case
			fmt.Println("send:", cases[chosen].Dir, cases[chosen].Send, ok)
		}
	}
}

func createCases(chs ...chan int) []reflect.SelectCase {
	var cases []reflect.SelectCase

	// 创建recv case
	for _, ch := range chs {
		cases = append(cases, reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(ch),
		})
	}

	// 创建send case
	for i, ch := range chs {
		v := reflect.ValueOf(i)
		cases = append(cases, reflect.SelectCase{
			Dir:  reflect.SelectSend,
			Chan: reflect.ValueOf(ch),
			Send: v,
		})
	}

	return cases
}
