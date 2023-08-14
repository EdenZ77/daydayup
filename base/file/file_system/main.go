package main

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
)

/*
参考文档：https://blog.csdn.net/andrewgithub/article/details/120497838
*/

//go:embed test.txt
var s string

//go:embed test.txt
var b []byte

//go:embed test.txt hello.txt
//go:embed file/file.txt
var f embed.FS

//go:embed file
var d embed.FS

//go:embed file/*.txt
var pre embed.FS

func main() {
	// 直接嵌入
	fmt.Println("========直接嵌入=========")
	fmt.Println(s)
	fmt.Println(b)

	// 嵌入为文件系统
	fmt.Println("========嵌入为文件系统：嵌入多个文件=========")
	data, _ := f.ReadFile("test.txt")
	fmt.Println(string(data))
	data, _ = f.ReadFile("hello.txt")
	fmt.Println(string(data))
	// 嵌入的时候文件是啥，这里要对应指定为相同的文件路径
	data, _ = f.ReadFile("file/file.txt")
	fmt.Println(string(data))

	fmt.Println("========嵌入为文件系统：嵌入文件夹=========")
	data, _ = d.ReadFile("file/file.txt")
	fmt.Println(string(data))

	fmt.Println("========嵌入为文件系统：嵌入指定后缀文件=========")
	//data, _ = pre.ReadFile("file/name.txt")
	//fmt.Println(string(data))

	sub, err := fs.Sub(pre, "file")
	if err != nil {
		fmt.Printf("fs.Sub err:%#v\n", err)
		return
	}
	file, err := sub.Open("name.txt")
	if err != nil {
		fmt.Printf("sub.Open err:%#v\n", err)
		return
	}
	b := make([]byte, 100)
	file.Read(b)
	fmt.Println(string(b))
}
