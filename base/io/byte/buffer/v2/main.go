package main

import (
	"bytes"
	"fmt"
)

func main() {
	// 使用 Write 方法实现 io.Writer 接口
	b := []byte("Hello")
	buf := bytes.NewBuffer(b) // 用现有的字节切片初始化缓冲区
	fmt.Fprintf(buf, "%s %d", "age", 25)
	fmt.Println(buf.String())
}
