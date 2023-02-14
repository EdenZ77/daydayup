package main

import (
	"fmt"
	"strings"
)

//func main() {
//	createStrStream()
//}

func createStrStream() {
	// 使用strings包下的NewReader创建字符串流
	r := strings.NewReader("Hello World! 你好,世界!")
	b := make([]byte, 4) // 创建字节切片,存放流中数据,根据流数据大小创建切片大小;
	n, err := r.Read(b)
	if err != nil {
		fmt.Println("流数据读取失败!", err)
		return
	}
	fmt.Println("读取的数据长度是: ", n)     // 读取的数据长度是:  27
	fmt.Println("数据内容: ", string(b)) // 以字符串形式显示切片中的数据  // 数据内容:  Hello World! 你好,世界!
}
