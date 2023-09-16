===这部分主要是channel相关的使用案例

无缓冲 channel 的惯用法
第一种用法：用作信号传递。无缓冲 channel 用作信号传递的时候，有两种情况，分别是 1 对 1 通知信号和 1 对 n 通知信号。



第二种用法：用于替代锁机制


注意：
使用 Channel 最常见的错误是 panic 和 goroutine 泄漏。
首先，我们来总结下会 panic 的情况，总共有 3 种：
1、close 为 nil 的 chan；
2、send 已经 close 的 chan；
3、close 已经 close 的 chan。

goroutine泄露示例
//=======
func process(timeout time.Duration) bool {
    ch := make(chan bool)

    go func() {
        // 模拟处理耗时的业务
        time.Sleep((timeout + time.Second))
        ch <- true // block
        fmt.Println("exit goroutine")
    }()
    select {
    case result := <-ch:
        return result
    case <-time.After(timeout):
        return false
    }
}
如果发生超时，process 函数就返回了，这就会导致 unbuffered 的 chan 从来就没有被读取。
我们知道，unbuffered chan 必须等 reader 和 writer 都准备好了才能交流，否则就会阻塞。
超时导致未读，结果就是子 goroutine 就阻塞在第 7 行永远结束不了，进而导致 goroutine 泄漏。
//=======


Go 的开发者极力推荐使用 Channel，不过，这两年，大家意识到，Channel 并不是处理并发问题的“银弹”，
有时候使用并发原语更简单，而且不容易出错。所以，我给你提供一套选择的方法:


//=======多生产者多消费者模型
