## 参考资料
github的研磨设计模式：https://github.com/senghoo/golang-design-pattern
国外设计模式翻译：https://refactoringguru.cn/design-patterns/creational-patterns
mohuishou：https://lailin.xyz/post/go-design-pattern.html
元润子：https://www.yrunz.com/tags/%E5%AE%9E%E8%B7%B5gof%E7%9A%8423%E7%A7%8D%E8%AE%BE%E8%AE%A1%E6%A8%A1%E5%BC%8F/

掘金go实现23种设计模式：https://juejin.cn/post/7095581880200167432#heading-1
k8s中使用的设计模式：https://juejin.cn/post/7087779598913519653#heading-16
极客时间：设计模式之美

### 包装器模式(装饰器模式)
当返回值的类型与参数类型相同时，我们能得到下面形式的函数原型：
```go
func YourWrapperFunc(param YourInterfaceType) YourInterfaceType
```
通过这个函数，我们可以实现对输入参数的类型的包装，并在不改变被包装类型（输入参数类型）的定义的情况下，返回具备新功能特性的、实现相同接口类型的新类型。这种接口应用模式我们叫它包装器模式，也叫装饰器模式。包装器多用于对输入数据的过滤、变换等操作。



