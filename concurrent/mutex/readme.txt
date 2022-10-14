学到这里，你可能要问了，虽然标准库 Mutex 不是可重入锁，但是如果我就是想要实现一个可重入锁，
可以吗？可以，那我们就自己实现一个。
这里的关键就是，实现的锁要能记住当前是哪个 goroutine 持有这个锁。我来提供两个方案。

11.1方案一：通过 hacker 的方式获取到 goroutine id，记录下获取锁的 goroutine id，它可以实现 Locker 接口。
11.2方案二：调用 Lock/Unlock 方法时，由 goroutine 提供一个 token，用来标识它自己，而不是我们通过 hacker 的方式获取到 goroutine id，但是，这样一来，就不满足 Locker 接口了。

11.3 扩展Mutex 实现TryLock
11.4 扩展Mutex 获取等待者的数量等指标
11.5 使用Mutex实现一个线程安全的队列



Mutex 易错的4大场景：

