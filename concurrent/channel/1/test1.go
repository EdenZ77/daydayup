package main

import "time"

/*
你一定要始终牢记：channel 是用于 Goroutine 间通信的，所以绝大多数对 channel 的读写都被分别放在了不同的 Goroutine 中。

由于无缓冲 channel 的运行时层实现不带有缓冲区，所以 Goroutine 对无缓冲 channel 的接收和发送操作是同步的。也就是说，
对同一个无缓冲 channel，只有对它进行接收操作的 Goroutine 和对它进行发送操作的 Goroutine 都存在的情况下，通信才能得以进行，否则单方面的操作会让对应的 Goroutine 陷入挂起状态

我们可以得出结论：对无缓冲 channel 类型的发送与接收操作，一定要放在两个不同的 Goroutine 中进行，否则会导致 deadlock。


*/

func main() {
	ch1 := make(chan int)
	go func() {
		time.Sleep(time.Second * 2)
		ch1 <- 13
	}()
	n := <-ch1
	println(n)
}
