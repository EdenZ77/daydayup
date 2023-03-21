package main

import (
	"bufio"
	"fmt"
	"strings"
)

//参考资料：https://books.studygolang.com/The-Golang-Standard-Library-by-Example/chapter01/01.4.html
// https://blog.csdn.net/ElzatAhmed/article/details/125748468      bufio.Reader源码详解
/*
ReadSlice、ReadBytes、ReadString 和 ReadLine 方法
之所以将这几个方法放在一起，是因为他们有着类似的行为。事实上，后三个方法最终都是调用ReadSlice来实现的。所以，我们先来看看ReadSlice方法。(感觉这一段直接看源码较好)

*/

/*
bufio.Reader.fill
该方法是 bufio.Reader 非常核心的方法之一，它将暂存有效（还没有被读取）的 buffer 数据左滑倒初始位置，并多次尝试从 rd 中读取新的字节，
只要读取到长度n大于0，错误等于nil的数据，将其append到buffer中。
*/

func main() {
	testReadBytes()
}

/*
ReadSlice方法签名如下：

	func (b *Reader) ReadSlice(delim byte) (line []byte, err error)

ReadSlice 从输入中读取，直到遇到第一个界定符（delim）为止，返回一个指向缓存中字节的 slice，在下次调用读操作（read）时，这些字节会无效。举例说明：
输出：

	the line:123#
	the line:4567
	456789====

注意：返回的结果是包含界定符本身的。
如果 ReadSlice 在找到界定符之前遇到了 error ，它就会返回缓存中所有的数据和错误本身（经常是 io.EOF）。
如果在找到界定符之前缓存已经满了，ReadSlice 会返回 bufio.ErrBufferFull 错误。
当且仅当返回的结果（line）没有以界定符结束的时候，ReadSlice 返回err != nil，
也就是说，如果ReadSlice 返回的结果 line 不是以界定符 delim 结尾，那么返回的 err 一定不等于 nil（可能是bufio.ErrBufferFull或io.EOF）。
*/
func testReadSlice() {
	reader := bufio.NewReaderSize(strings.NewReader("1234567890abcdf67#1234567890abcdf6789===="), 16)
	line, err1 := reader.ReadSlice('#')
	if err1 != nil {
		if err1 == bufio.ErrBufferFull {
			fmt.Printf("%+v\n", err1)
		}
	}
	fmt.Printf("the line-1:%s\n", line)
	// 这里可以换上任意的 bufio 的 Read/Write 操作
	n, err := reader.ReadSlice('#')
	// 如果 ReadSlice 在找到界定符之前遇到了 error ，它就会返回缓存中所有的数据和错误本身（经常是 io.EOF）
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("the line-2:%s\n", line)
	fmt.Println(string(n))
}

/*
func (b *Reader) ReadBytes(delim byte) (line []byte, err error)

该方法的参数和返回值类型与 ReadSlice 都一样。 ReadBytes 从输入中读取直到遇到界定符（delim）为止，返回的 slice 包含了从当前到界定符的内容 （包括界定符）。
在讲解ReadSlice时说到，它返回的 []byte 是指向 Reader 中的 buffer，而不是 copy 一份返回，也正因为如此，通常我们会使用 ReadBytes 或 ReadString。
很显然，ReadBytes 返回的 []byte 不会是指向 Reader 中的 buffer，通过查看源码可以证实这一点。
输出：

	the line-1:123#
	the line-2:123#
	4567===
*/
func testReadBytes() {
	reader := bufio.NewReader(strings.NewReader("123#4567==="))
	line, _ := reader.ReadBytes('#')
	fmt.Printf("the line-1:%s\n", line)
	// 这里可以换上任意的 bufio 的 Read/Write 操作
	n, _ := reader.ReadBytes('#')
	fmt.Printf("the line-2:%s\n", line)
	fmt.Println(string(n))
}

/*
	func (b *Reader) ReadString(delim byte) (line string, err error) {
	    bytes, err := b.ReadBytes(delim)
	    return string(bytes), err
	}

它调用了 ReadBytes 方法，并将结果的 []byte 转为 string 类型。
*/
func testReadString() {
}

/*
func (b *Reader) ReadLine() (line []byte, isPrefix bool, err error)
ReadLine 是一个底层的原始行读取命令。许多调用者或许会使用 ReadBytes('\n') 或者 ReadString('\n') 来代替这个方法。
*/
func testReadLine() {
}
