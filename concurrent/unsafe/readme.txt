参考资料：
极客兔兔：https://geektutu.com/post/hpg-struct-alignment.html
https://www.cnblogs.com/-wenli/p/12682477.html
https://mp.weixin.qq.com/s/dulgHWM-mjrYIdD9nHZyYg

unsafe.Pointer称为通用指针，官方文档对该类型有四个重要描述：
（1）任何类型的指针都可以被转化为Pointer
（2）Pointer可以被转化为任何类型的指针
（3）uintptr可以被转化为Pointer
（4）Pointer可以被转化为uintptr
unsafe.Pointer是特别定义的一种指针类型（译注：类似C语言中的void类型的指针），在golang中是用于各种指针相互转换的桥梁，它可以包含任意类型变量的地址。
当然，我们不可以直接通过*p来获取unsafe.Pointer指针指向的真实变量的值，因为我们并不知道变量的具体类型。


Golang指针
*类型: 普通指针类型，用于传递对象地址，不能进行指针运算。
unsafe.Pointer: 通用指针类型，用于转换不同类型的指针，不能进行指针运算，不能读取内存存储的值（必须转换到某一类型的普通指针）。
uintptr: 用于指针运算，GC 不把 uintptr 当指针，uintptr 无法持有对象。uintptr 类型的目标会被回收。

unsafe.Pointer 是桥梁，可以让任意类型的指针实现相互转换，也可以将任意类型的指针转换为 uintptr 进行指针运算。
unsafe.Pointer 不能参与指针运算，比如你要在某个指针地址上加上一个偏移量，Pointer是不能做这个运算的，那么谁可以呢?
就是uintptr类型了，只要将Pointer类型转换成uintptr类型，做完加减法后，转换成Pointer，通过*操作，取值，修改值，随意。

总结：unsafe.Pointer 可以让你的变量在不同的普通指针类型转来转去，也就是表示为任意可寻址的指针类型。而 uintptr 常用于与 unsafe.Pointer 打配合，用于做指针运算。

## 新的理解，参考资料：https://blog.betacat.io/post/golang-atomic-value-exploration/
unsafe.Pointer的特别之处在于，它可以绕过Go语言类型系统的检查，与任意的指针类型互相转换。也就是说，如果两个类型具有相同的内存结构(layout),
我们可以将unsafe.Pointer当作桥梁，让这两种类型的指针相互转换，从而实现同一份内存拥有两种不同的解读方式。

比如说，[]byte和string其实内部的存储结构都是一样的，但 Go 语言的类型系统禁止他俩互换。如果借助unsafe.Pointer，我们就可以实现在零拷贝的情况下，将[]byte数组直接转换成string类型。
bytes := []byte{104, 101, 108, 108, 111}

p := unsafe.Pointer(&bytes) //强制转换成unsafe.Pointer，编译器不会报错
str := *(*string)(p) //然后强制转换成string类型的指针，再将这个指针的值当做string类型取出来
fmt.Println(str) //输出 "hello"

