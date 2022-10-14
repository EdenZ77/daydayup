参考资料
https://segmentfault.com/a/1190000040917752
https://zhuanlan.zhihu.com/p/68792989

context 主要用来在 goroutine 之间传递上下文信息，包括：取消信号、超时时间、截止时间、k-v 等。

context包主要提供了两种方式创建context:
context.Backgroud()
context.TODO()
这两个函数其实只是互为别名，没有差别，官方给的定义是：
context.Background 是上下文的默认值，所有其他的上下文都应该从它衍生（Derived）出来。
context.TODO 应该只在不确定应该使用哪种上下文时使用；
所以在大多数情况下，我们都使用context.Background作为起始的上下文向下传递。

上面的两种方式是创建根context，不具备任何功能，具体实践还是要依靠context包提供的With系列函数来进行派生：
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
func WithValue(parent Context, key, val interface{}) Context

这四个函数都要基于父Context衍生，通过这些函数，就创建了一颗Context树，树的每个节点都可以有任意多个子节点，节点层级可以有任意多个，画个图表示一下：
WithValue携带数据







