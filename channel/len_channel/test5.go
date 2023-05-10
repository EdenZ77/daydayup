package main

import (
	"fmt"
	"sync"
	"time"
)

/*

var ch chan T = make(chan T, capacity)

// 判空
if len(ch) == 0 {
    // 此时channel ch空了?
}

// 判有
if len(ch) > 0 {
    // 此时channel ch中有数据?
}

// 判满
if len(ch) == cap(ch) {
    // 此时channel ch满了?
}

你可以看到，我在上面代码注释的“空了”、“有数据”和“满了”的后面都打上了问号。这是为什么呢？
这是因为，channel 原语用于多个 Goroutine 间的通信，一旦多个 Goroutine 共同对 channel 进行收发操作，len(channel) 就会在多个 Goroutine 间形成“竞态”。单纯地依靠 len(channel) 来判断 channel 中元素状态，是不能保证在后续对 channel 的收发时 channel 状态是不变的。

channel 原语用于多个 Goroutine 间的通信，一旦多个 Goroutine 共同对 channel 进行收发操作，len(channel) 就会在多个 Goroutine 间形成“竞态”。
单纯地依靠 len(channel) 来判断 channel 中元素状态，是不能保证在后续对 channel 的收发时 channel 状态是不变的。

为了不阻塞在 channel 上，常见的方法是将“判空与读取”放在一个“事务”中，将“判满与写入”放在一个“事务”中，而这类“事务”我们可以通过 select 实现
*/

func producer(c chan<- int) {
	var i int = 1
	// 这个生产者通过for循环不停的向channel发送数据，一般来说，当channel满了就会阻塞，我们这里实现了非阻塞，
	for {
		time.Sleep(2 * time.Second)
		// 非阻塞的发送数据
		ok := trySend(c, i)
		if ok {
			fmt.Printf("[producer]: send [%d] to channel\n", i)
			i++
			continue
		}
		fmt.Printf("[producer]: try send [%d], but channel is full\n", i)
	}
}

// 实现“事务”操作
func trySend(c chan<- int, i int) bool {
	select {
	case c <- i:
		return true
	default:
		return false
	}
}

// 非阻塞的接收数据
func tryRecv(c <-chan int) (int, bool) {
	select {
	case i := <-c:
		return i, true
	default:
		return 0, false
	}
}

func consumer(c <-chan int) {
	for {
		i, ok := tryRecv(c)
		if !ok {
			fmt.Println("[consumer]: try to recv from channel, but the channel is empty")
			time.Sleep(1 * time.Second)
			continue
		}
		fmt.Printf("[consumer]: recv [%d] from channel\n", i)
		if i >= 3 {
			fmt.Println("[consumer]: exit")
			return
		}
	}
}

func main() {
	var wg sync.WaitGroup
	c := make(chan int, 3)
	wg.Add(2)
	go func() {
		defer wg.Done()
		producer(c)
	}()

	go func() {
		defer wg.Done()
		consumer(c)
	}()

	wg.Wait()
}
