
策略一：透明错误处理策略=========================================
简单来说，Go 语言中的错误处理，就是根据函数 / 方法返回的 error 类型变量中携带的错误值信息做决策，并选择后续代码执行路径的过程。
这样，最简单的错误策略莫过于完全不关心返回错误值携带的具体上下文信息，只要发生错误就进入唯一的错误处理执行路径，比如下面这段代码：
err := doSomething()
if err != nil {
    // 不关心err变量底层错误值所携带的具体上下文信息
    // 执行简单错误处理逻辑并返回
    ... ...
    return err
}
这也是 Go 语言中最常见的错误处理策略，80% 以上的 Go 错误处理情形都可以归类到这种策略下。
在这种策略下，由于错误处理方并不关心错误值的上下文，所以错误值的构造方（如上面的函数doSomething）可以直接使用 Go 标准库提供的两个基本错误值构造方法errors.New和fmt.Errorf来构造错误值，就像下面这样：
func doSomething(...) error {
    ... ...
    return errors.New("some error occurred")
}
这样构造出的错误值代表的上下文信息，对错误处理方是透明的，因此这种策略称为“透明错误处理策略”。在错误处理方不关心错误值上下文的前提下，透明错误处理策略能最大程度地减少错误处理方与错误值构造方之间的耦合关系。


策略二：“哨兵”错误处理策略=========================================
当错误处理方不能只根据“透明的错误值”就做出错误处理路径选取的情况下，错误处理方会尝试对返回的错误值进行检视，于是就有可能出现下面代码中的反模式：
data, err := b.Peek(1)
if err != nil {
    switch err.Error() {
    case "bufio: negative count":
        // ... ...
        return
    case "bufio: buffer full":
        // ... ...
        return
    case "bufio: invalid use of UnreadByte":
        // ... ...
        return
    default:
        // ... ...
        return
    }
}
简单来说，反模式就是，错误处理方以透明错误值所能提供的唯一上下文信息（描述错误的字符串），作为错误处理路径选择的依据。但这种“反模式”会造成严重的隐式耦合。
这也就意味着，错误值构造方不经意间的一次错误描述字符串的改动，都会造成错误处理方处理行为的变化，并且这种通过字符串比较的方式，对错误值进行检视的性能也很差。

那这有什么办法吗？Go 标准库采用了定义导出的（Exported）“哨兵”错误值的方式，来辅助错误处理方检视（inspect）错误值并做出错误处理分支的决策，比如下面的 bufio 包中定义的“哨兵错误”：
// $GOROOT/src/bufio/bufio.go
var (
    ErrInvalidUnreadByte = errors.New("bufio: invalid use of UnreadByte")
    ErrInvalidUnreadRune = errors.New("bufio: invalid use of UnreadRune")
    ErrBufferFull        = errors.New("bufio: buffer full")
    ErrNegativeCount     = errors.New("bufio: negative count")
)
下面的代码片段利用了上面的哨兵错误，进行错误处理分支的决策：
data, err := b.Peek(1)
if err != nil {
    switch err {
    case bufio.ErrNegativeCount:
        // ... ...
        return
    case bufio.ErrBufferFull:
        // ... ...
        return
    case bufio.ErrInvalidUnreadByte:
        // ... ...
        return
    default:
        // ... ...
        return
    }
}
你可以看到，一般“哨兵”错误值变量以 ErrXXX 格式命名。和透明错误策略相比，“哨兵”策略让错误处理方在有检视错误值的需求时候，可以“有的放矢”。
不过，对于 API 的开发者而言，暴露“哨兵”错误值也意味着这些错误值和包的公共函数 / 方法一起成为了 API 的一部分。一旦发布出去，开发者就要对它进行很好的维护。而“哨兵”错误值也让使用这些值的错误处理方对它产生了依赖。
从 Go 1.13 版本开始，标准库 errors 包提供了 Is 函数用于错误处理方对错误值的检视。Is 函数类似于把一个 error 类型变量与“哨兵”错误值进行比较，比如下面代码：
// 类似 if err == ErrOutOfBounds{ … }
if errors.Is(err, ErrOutOfBounds) {
    // 越界的错误处理
}
不同的是，如果 error 类型变量的底层错误值是一个包装错误（Wrapped Error），errors.Is 方法会沿着该包装错误所在错误链（Error Chain)，与链上所有被包装的错误（Wrapped Error）进行比较，直至找到一个匹配的错误为止。下面是 Is 函数应用的一个例子：
var ErrSentinel = errors.New("the underlying sentinel error")
func main() {
  err1 := fmt.Errorf("wrap sentinel: %w", ErrSentinel)
  err2 := fmt.Errorf("wrap err1: %w", err1)
    println(err2 == ErrSentinel) //false
  if errors.Is(err2, ErrSentinel) {
    println("err2 is ErrSentinel")
    return
  }

  println("err2 is not ErrSentinel")
}
在这个例子中，我们通过 fmt.Errorf 函数，并且使用 %w 创建包装错误变量 err1 和 err2，其中 err1 实现了对 ErrSentinel 这个“哨兵错误值”的包装，而 err2 又对 err1 进行了包装，这样就形成了一条错误链。位于错误链最上层的是 err2，位于最底层的是 ErrSentinel。之后，我们再分别通过值比较和 errors.Is 这两种方法，判断 err2 与 ErrSentinel 的关系。运行上述代码，我们会看到如下结果：
false
err2 is ErrSentinel
我们看到，通过比较操作符对 err2 与 ErrSentinel 进行比较后，我们发现这二者并不相同。而 errors.Is 函数则会沿着 err2 所在错误链，向下找到被包装到最底层的“哨兵”错误值ErrSentinel。

策略三：错误值类型检视策略=================================================
上面我们看到，基于 Go 标准库提供的错误值构造方法构造的“哨兵”错误值，除了让错误处理方可以“有的放矢”的进行值比较之外，并没有提供其他有效的错误上下文信息。那如果遇到错误处理方需要错误值提供更多的“错误上下文”的情况，上面这些错误处理策略和错误值构造方式都无法满足。

这种情况下，我们需要通过自定义错误类型的构造错误值的方式，来提供更多的“错误上下文”信息。并且，由于错误值都通过 error 接口变量统一呈现，
要得到底层错误类型携带的错误上下文信息，错误处理方需要使用 Go 提供的类型断言机制（Type Assertion）或类型选择机制（Type Switch），这种错误处理方式，我称之为错误值类型检视策略。

这里，一般自定义导出的错误类型以XXXError的形式命名。和“哨兵”错误处理策略一样，错误值类型检视策略，由于暴露了自定义的错误类型给错误处理方，
因此这些错误类型也和包的公共函数 / 方法一起，成为了 API 的一部分。一旦发布出去，开发者就要对它们进行很好的维护。而它们也让使用这些类型进行检视的错误处理方对其产生了依赖。

从 Go 1.13 版本开始，标准库 errors 包提供了As函数给错误处理方检视错误值。As函数类似于通过类型断言判断一个 error 类型变量是否为特定的自定义错误类型，如下面代码所示：
// 类似 if e, ok := err.(*MyError); ok { … }
var e *MyError
if errors.As(err, &e) {
    // 如果err类型为*MyError，变量e将被设置为对应的错误值
}
不同的是，如果 error 类型变量的动态错误值是一个包装错误，errors.As函数会沿着该包装错误所在错误链，与链上所有被包装的错误的类型进行比较，直至找到一个匹配的错误类型，就像 errors.Is 函数那样。下面是As函数应用的一个例子：
type MyError struct {
    e string
}
func (e *MyError) Error() string {
    return e.e
}
func main() {
    var err = &MyError{"MyError error demo"}
    err1 := fmt.Errorf("wrap err: %w", err)
    err2 := fmt.Errorf("wrap err1: %w", err1)
    var e *MyError
    if errors.As(err2, &e) {
        println("MyError is on the chain of err2")
        println(e == err)
        return
    }
    println("MyError is not on the chain of err2")
}
运行上述代码会得到：
MyError is on the chain of err2
true
我们看到，errors.As函数沿着 err2 所在错误链向下找到了被包装到最深处的错误值，并将 err2 与其类型 * MyError成功匹配。
匹配成功后，errors.As 会将匹配到的错误值存储到 As 函数的第二个参数中，这也是为什么println(e == err)输出 true 的原因。

策略四：错误行为特征检视策略========================================================
不知道你注意到没有，在前面我们已经讲过的三种策略中，其实只有第一种策略，也就是“透明错误处理策略”，有效降低了错误的构造方与错误处理方两者之间的耦合。
虽然前面的策略二和策略三，都是我们实际编码中有效的错误处理策略，但其实使用这两种策略的代码，依然在错误的构造方与错误处理方两者之间建立了耦合。

在 Go 标准库中，我们发现了这样一种错误处理方式：将某个包中的错误类型归类，统一提取出一些公共的错误行为特征，并将这些错误行为特征放入一个公开的接口类型中。这种方式也被叫做错误行为特征检视策略。
以标准库中的net包为例，它将包内的所有错误类型的公共行为特征抽象并放入net.Error这个接口中，如下面代码：
// $GOROOT/src/net/net.go
type Error interface {
    error
    Timeout() bool
    Temporary() bool
}
我们看到，net.Error 接口包含两个用于判断错误行为特征的方法：Timeout 用来判断是否是超时（Timeout）错误，Temporary 用于判断是否是临时（Temporary）错误。
而错误处理方只需要依赖这个公共接口，就可以检视具体错误值的错误行为特征信息，并根据这些信息做出后续错误处理分支选择的决策。

这里，我们再看一个 http 包使用错误行为特征检视策略进行错误处理的例子，加深下理解：


pkg/errors使用
















