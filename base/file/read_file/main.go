package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fileRead_case2()
}

/*
ioutil.ReadFile读取文件
func ReadFile(filename string) ([]byte, error)

参数	描述
filename	文件名

返回值	描述
[]byte	读取到的文件内容。
error	如果读取失败，返回错误信息，否则，返回 nil

说明
ioutil.ReadFile 读取文件只需要传入一个文件名做为 参数，读取成功，会将文件的内容做为一个字节数组返回，如果读取错误，将返回 error 信息。
使用 ReadFile 读取文件，不需要手动 打开与关闭文件，打开与关闭文件的动作，系统自动帮我们完成。同时，使用 ReadFile 读取文件时，只适合读取小文件，不适合读取大文件。
*/
func ioutilReadFile_case1() {
	//fileName := "E:\\workspace\\go\\daydayup\\base\\file\\haicoder.txt"
	// 这两种路径方式都可以
	fileName := "E:/workspace/go/daydayup/base/file/haicoder.txt"
	// 这个ioutil.ReadFile在go1.16的时候废弃了，因为这个函数仅仅是调用了一下os.ReadFile，所以我们直接使用os.ReadFile即可
	// 因为ReadFile读取整个文件，所以它不会将Read中的EOF视为要报告的错误。
	fileData, err := ioutil.ReadFile(fileName)
	//fileData, err := os.ReadFile(fileName)
	if err == nil {
		fmt.Println("file content =", string(fileData))
	} else {
		fmt.Println("read file error =", err)
	}
}

/*
func (f *File) Read(b []byte) (n int, err error)

参数	描述
f	打开的文件句柄。
b	读取文件内容存放的切片（根据切片的len长度来存放的，）

返回值	说明
n	读取到的字节数。
err	如果读取失败，返回错误信息，否则，返回 nil

说明
使用 file.Read 读取文件时，首先，我们需要打开文件，接着， 使用打开的文件返回的文件句柄，来读取文件。
文件读取结束的标志是返回的 n 等于 0，因此，如果我们需要读取整个文件内容，那么我们需要使用 for 循环 不停的读取文件，直到 n 等于 0。
*/
func fileRead_case2() {
	fileName := "E:/workspace/go/daydayup/base/file/haicoder.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Read file err, err =", err)
		return
	}
	defer file.Close()
	var chunk []byte
	// len=1024
	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println("read buf fail", err)
			return
		}
		// 读取到文件最后，返回n=0,io.EOF
		if n == 0 {
			break
		}
		//读取到最终的缓冲区中
		chunk = append(chunk, buf[:n]...)
	}
	fmt.Println("File Content =", string(chunk))
}

/*
语法
r := bufio.NewReader(file)
n, err := r.Read(buf)

参数	描述
file	要读取的文件句柄。
buf	读取的数据存放的缓冲区。

返回值	描述
n	读取到的长度。
err	读取失败，则返回错误信息。

说明
使用 NewReader 读取文件时，首先，我们需要打开文件，接着， 使用打开的文件返回的文件句柄当作 函数参数 传入 NewReader。
最后，我们使用 NewReader 返回的 reader 对象调用 Read 来读取文件。文件读取结束的标志是返回的 n 等于 0，因此，如果我们需要读取整个文件内容，那么我们需要使用 for 循环 不停的读取文件，直到 n 等于 0。
*/
func bufioNewRead_case3() {
	fileName := "E:/workspace/go/daydayup/base/file/haicoder.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Read file err, err =", err)
		return
	}
	defer file.Close()
	r := bufio.NewReader(file)
	var chunk []byte
	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println("read buf fail", err)
			return
		}
		//说明读取结束
		if n == 0 {
			break
		}
		//读取到最终的缓冲区中
		chunk = append(chunk, buf[:n]...)
	}
	fmt.Println("File Content =", string(chunk))
}

/*
func ReadAll(r io.Reader) ([]byte, error)

参数	描述
r	Reader 对象。

返回值	描述
[]byte	读取到的数据。
error	读取失败，则返回错误信息。

说明
使用 ReadAll 读取文件时，首先，我们需要打开文件，接着， 使用打开的文件返回的文件句柄当作 函数参数 传入 ReadAll。
ReadAll 函数将会将整个文件的内容一次性读取出来，如果读取出错，则返回 error 信息。
*/
func ioutilReadAll_case4() {
	fileName := "D:\\workspace\\go_project\\study\\daydayup\\base\\file\\haicoder.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Read file err, err =", err)
		return
	}
	defer file.Close()
	// 这个ioutil.ReadAll在go1.16的时候废弃了，因为这个函数仅仅是调用了一下os.ReadAll，所以我们直接使用os.ReadAll即可
	fileContent, err := ioutil.ReadAll(file)
	//fileContent, err := io.ReadAll(file)
	if err == nil {
		fmt.Println("File Content =", string(fileContent))
	} else {
		fmt.Println("Read file err, err =", err)
	}
}
