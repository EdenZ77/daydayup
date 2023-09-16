## 参考资料
https://mp.weixin.qq.com/s/QFr_Pgt9GzOG3zlgcXEVwQ

## 介绍一
常量：不可改变的量，相对于变量来说，也可以说成不可改变的变量；常量只能在初始化时赋值，其他情况下不允许赋值，
    定义常量优点：a.防止在后续程序中被修改 b.可以重复使用
    如：const double P=3.14;

枚举：我们可以定义变量中的类型，然后对变量值进行一定的限制，当用户对我们这种类型进行赋值时，只能赋咱们限定好的值。
我们可以重新定义一种类型，并且在定义这种类型时，我们要指定这个类型的所有值
    enum 自己起的类型名称{类型值1，类型值2，......,类型值n}
    优点： a.限制用户不能随意赋值,只能赋在定义枚举时列举的值中选择 b.不需要死记每个值是什么，只需要选择相应的值就可以了

c语言枚举示例：
```c
enum Weekday {
    SUNDAY,
    MONDAY,
    TUESDAY,
    WEDNESDAY,
    THURSDAY,
    FRIDAY,
    SATURDAY
};
int main() {
    enum Weekday d = SATURDAY;
    printf("%d\n", d); // 6
}
```
你运行上面的 C 语言代码就会发现，其实 C 语言针对枚举类型提供了很多语法上的便利特性。
比如说，如果你没有显式给枚举常量赋初始值，那么枚举类型的第一个常量的值就为 0，后续常量的值再依次加 1。

但 Go 并没有直接继承这一特性，而是将 C 语言枚举类型的这种基于前一个枚举值加 1 的特性，
分解成了 Go 中的两个特性：自动重复上一行，以及引入 const 块中的行偏移量指示器 iota，这样它们就可以分别独立使用了。

## 介绍二

枚举类型的值本质上是常量，因此我们可以使用 Go 语言中的常量来实现类似枚举类型的功能，例如：
```go
const (
   Sunday    = 1
   Tuesday   = 2
   Wednesday = 3
   Thursday  = 4
   Friday    = 5
   Saturday  = 6
   Monday    = 7
)
```
在这个例子中，我们使用了 const 关键字定义了一组常量，其中每个常量的名称代表着一个枚举，其值为对应的整数。

虽然这个例子能实现类似的枚举类型，但它不具备枚举类型的所有特征，例如缺少安全性和约束性，为了解决这两个问题，我们可以使用自定义类型进行改进：
```go
type WeekDay int

const (
   Sunday    WeekDay = 1
   Tuesday   WeekDay = 2
   Wednesday WeekDay = 3
   Thursday  WeekDay = 4
   Friday    WeekDay = 5
   Saturday  WeekDay = 6
   Monday    WeekDay = 7
)
```
在改进后的例子中，我们定义了一个自定义类型 Weekday，用于表示星期几。使用 const 关键字定义一个常量组，其中每个常量都被赋予了一个具体的值，同时使用 Weekday 类型进行类型约束和类型检查。这样，我们就可以通过枚举值的名称来表示某个特定的星期几，并且由于使用了自定义类型，编译器可以进行类型检查，从而提高了类型安全性。

### 使用iota优雅实现枚举
通过前面的例子不难发现，当我们需要定义多个枚举值时，手动指定每个枚举常量的值会变得十分麻烦。为了解决这个问题，我们可以使用 iota 常量生成器，它可以帮助我们生成连续的整数值。

例如，使用 iota 定义一个星期几的枚举类型：
```go
type WeekDay int

const (
	// 这里const块中的行偏移量iota能赋值给WeekDay类型，主要依赖无类型常量能够隐式转换
   Sunday WeekDay = iota
   Tuesday
   Wednesday
   Thursday
   Friday
   Saturday
   Monday
)
```





