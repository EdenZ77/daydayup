unsafe.Pointer称为通用指针，官方文档对该类型有四个重要描述：
（1）任何类型的指针都可以被转化为Pointer
（2）Pointer可以被转化为任何类型的指针
（3）uintptr可以被转化为Pointer
（4）Pointer可以被转化为uintptr
unsafe.Pointer是特别定义的一种指针类型（译注：类似C语言中的void类型的指针），在golang中是用于各种指针相互转换的桥梁，它可以包含任意类型变量的地址。
当然，我们不可以直接通过*p来获取unsafe.Pointer指针指向的真实变量的值，因为我们并不知道变量的具体类型。
