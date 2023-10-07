package main

import (
	"reflect"
	"unsafe"
)

// 参考资料：https://mp.weixin.qq.com/s/dulgHWM-mjrYIdD9nHZyYg
/*
经典应用：string与[]byte的相互转换
实现string与byte的转换，正常情况下，我们可能会写出这样的标准转换：

// string to []byte
str1 := "Golang梦工厂"
by := []byte(s1)

// []byte to string
str2 := string(by)

使用这种方式进行转换都会涉及底层数值的拷贝，所以想要实现零拷贝，我们可以使用unsafe.Pointer来实现，
通过强转换直接完成指针的指向，从而使string和[]byte指向同一个底层数据。

在reflect包中有·string和slice对应的结构体，他们的分别是：
type StringHeader struct {
 Data uintptr
 Len  int
}

type SliceHeader struct {
 Data uintptr
 Len  int
 Cap  int
}

StringHeader代表的是string运行时的表现形式(SliceHeader同理)，通过对比string和slice运行时的表达可以看出，
他们只有一个Cap字段不同，所以他们的内存布局是对齐的，所以可以通过unsafe.Pointer进行转换，因为可以写出如下代码：

*/

func stringToByte(s string) []byte {
	header := (*reflect.StringHeader)(unsafe.Pointer(&s))

	newHeader := reflect.SliceHeader{
		Data: header.Data,
		Len:  header.Len,
		Cap:  header.Len,
	}

	return *(*[]byte)(unsafe.Pointer(&newHeader))
}

func bytesToString(b []byte) string {
	header := (*reflect.SliceHeader)(unsafe.Pointer(&b))

	newHeader := reflect.StringHeader{
		Data: header.Data,
		Len:  header.Len,
	}

	return *(*string)(unsafe.Pointer(&newHeader))
}

/*
上面的代码我们通过重新构造slice header和string header完成了类型转换，
其实[]byte转换成string可以省略掉自己构造StringHeader的方式，直接使用强转就可以，因为string的底层也是[]byte，强转会自动构造，省略后的代码如下：
*/
func bytesToString2(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

/*
虽然这种方式更高效率，但是不推荐大家使用，前面也提到了，这要是不安全的，使用不当会出现极大的隐患，一些严重的情况recover也不能捕获。
*/
func main() {

}
