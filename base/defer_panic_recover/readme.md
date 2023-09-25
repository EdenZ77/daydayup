参考资料: https://www.jianshu.com/p/63e3d57f285f

## 1. defer
defer后边会接一个函数，但该函数不会立刻被执行，而是等到包含它的程序返回时(包含它的函数执行了return语句、运行到函数结尾自动返回、对应的goroutine panic）defer函数才会被执行。通常用于资源释放、打印日志、异常捕获等

如果有多个defer函数，调用顺序类似于栈，越后面的defer函数越先被执行(后进先出)
```go
func main() {
    defer fmt.Println("1")
    defer fmt.Println("2")
    defer fmt.Println("3")
}
// 结果
3
2
1
```



## 2. panic
```go
func panic(v interface{})
```
panic内置函数停止当前goroutine的正常执行，当函数F调用panic时，函数F的正常执行被立即停止，然后运行所有在F函数中的defer函数，然后F返回到调用他的函数对于调用者G，F函数的行为就像panic一样，终止G的执行并运行G中所defer函数，此过程会一直继续执行到goroutine所有的函数。panic可以通过内置的recover来捕获。

## 3. recover
```go
func recover() interface{}
```
recover内置函数用来管理含有panic行为的goroutine，recover运行在defer函数中，获取panic抛出的错误值，并将程序恢复成正常执行的状态。如果在defer函数之外调用recover，那么recover不会停止并且捕获panic错误如果goroutine中没有panic或者捕获的panic的值为nil，recover的返回值也是nil。由此可见，recover的返回值表示当前goroutine是否有panic行为






