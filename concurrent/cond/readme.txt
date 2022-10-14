
Cond的实现原理
Wait：

runtime_notifyListXXX 是运行时实现的方法，实现了一个等待 / 通知的队列。如果你想深入学习这部分，可以再去看看 runtime/sema.go 代码中。
copyChecker 是一个辅助结构，可以在运行时检查 Cond 是否被复制使用。
Wait 把调用者加入到等待队列时会释放锁，在被唤醒之后还会请求锁。在阻塞休眠期间，调用者是不持有锁的，这样能让其他 goroutine 有机会检查或者更新等待变量。

常见错误：
调用 Wait 的时候没有加锁

Cond 有三点特性是 Channel 无法替代的：
Cond 和一个 Locker 关联，可以利用这个 Locker 对相关的依赖条件更改提供保护。
Cond 可以同时支持 Signal 和 Broadcast 方法，而 Channel 只能同时支持其中一种。
Cond 的 Broadcast 方法可以被重复调用。等待条件再次变成不满足的状态后，我们又可以调用 Broadcast 再次唤醒等待的 goroutine。
    这也是 Channel 不能支持的，Channel 被 close 掉了之后不支持再 open。

在实践中，处理等待 / 通知的场景时，我们常常会使用 Channel 替换 Cond，因为 Channel 类型使用起来更简洁，而且不容易出错。
但是对于需要重复调用 Broadcast 的场景，比如上面 Kubernetes 的例子，每次往队列中成功增加了元素后就需要调用 Broadcast 通知所有的等待者，使用 Cond 就再合适不过了。

使用 Cond 之所以容易出错，就是 Wait 调用需要加锁，以及被唤醒后一定要检查条件是否真的已经满足。你需要牢记这两点。

本质上 WaitGroup 和 Cond 是有区别的：WaitGroup 是主 goroutine 等待确定数量的子 goroutine 完成任务；
而 Cond 是等待某个条件满足，这个条件的修改可以被任意多的 goroutine 更新，而且 Cond 的 Wait 不关心也不知道其他 goroutine 的数量，只关心等待条件。
而且 Cond 还有单个通知的机制，也就是 Signal 方法。

一个 Cond 的 waiter 被唤醒的时候，为什么需要再检查等待条件，而不是唤醒后进行下一步？
唤醒的方式有broadcast，第N个waiter被唤醒后需要检查等待条件，因为不知道前N-1个被唤醒的waiter所作的修改是否使等待条件再次成立。
其实这也是wait方法先unlock，然后挂起，被唤醒后lock的原因，因为被唤醒后这个goroutine可能会修改条件，所以要加锁！

示例展示了k8s在优先级队列中使用了Cond，说实话，很少在项目中看到使用队列这个数据结构，在k8s应该会大量使用到，因为k8s调度会有大量优先级的处理
这个时候使用优先级队列就很合适，测试3就实现一个有限容量的queue


