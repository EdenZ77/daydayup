state1，一个具有复合意义的字段，包含 WaitGroup 的计数、阻塞在检查点的 waiter 数和信号量

梳理大致的实现原理：

Add 方法主要操作的是 state 的计数部分。你可以为计数值增加一个 delta 值，内部通过原子操作把这个值加到计数值上。
需要注意的是，这个 delta 也可以是个负数，相当于为计数值减去一个值，Done 方法内部其实就是通过 Add(-1) 实现的。

Wait 方法的实现逻辑是：不断检查 state 的值。如果其中的计数值变为了 0，那么说明所有的任务已完成，调用者不必再等待，直接返回。
如果计数值大于 0，说明此时还有任务没完成，那么调用者就变成了等待者，需要加入 waiter 队列，并且阻塞住自己。

如果计数值v为0并且waiter的数量w不为0，那么state的值就是waiter的数量
将waiter的数量设置为0，因为计数值v也是0,所以它们俩的组合*statep直接设置为0即可。此时需要并唤醒所有的waiter

常见错误：



noCopy：辅助 vet 检查
其实，它就是指示 vet 工具在做检查的时候，这个数据结构不能做值复制使用。
更严谨地说，是不能在第一次使用之后复制使用 ( must not be copied after first use)。


vet 会对实现 Locker 接口的数据类型做静态检查，一旦代码中有复制使用这种数据类型的情况，就会发出警告。
但是，WaitGroup 同步原语不就是 Add、Done 和 Wait 方法吗？vet 能检查出来吗？
其实是可以的。通过给 WaitGroup 添加一个 noCopy 字段，我们就可以为 WaitGroup 实现 Locker 接口，这样 vet 工具就可以做复制检查了。
而且因为 noCopy 字段是未输出类型，所以 WaitGroup 不会暴露 Lock/Unlock 方法。
