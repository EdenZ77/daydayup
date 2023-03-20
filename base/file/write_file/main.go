package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fileWrite_case3()
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

/*
func WriteString(w Writer, s string) (n int, err error)

参数	描述
w	Writer 对象。
s	要写入的文件内容。

返回值	描述
n	写入的字节数。
err	写入失败，则返回错误信息。

说明
使用 WriteString 方法写文件，接受的第一个 参数 是一个 Writer 接口，第二个参数是一个 stirng 类型的要写入的字符串。
如果写入成功，返回成功写入的字节数，如果写入失败，返回 error 信息，在写入文件之前，我们需要判断文件是否存在，如果文件不存在，则需要创建文件。
*/
func ioWriteString_case1() {
	var (
		fileName = "D:\\workspace\\go_project\\study\\daydayup\\base\\file\\haicoder.txt"
		content  = "Hello HaiCoder"
		file     *os.File
		err      error
	)
	if Exists(fileName) {
		//使用追加模式打开文件
		file, err = os.OpenFile(fileName, os.O_APPEND, 0666)
		if err != nil {
			fmt.Println("Open file err =", err)
			return
		}
	} else {
		file, err = os.Create(fileName)
		if err != nil {
			fmt.Println("file create fail")
			return
		}
	}
	defer file.Close()
	//写入文件
	n, err := io.WriteString(file, content)
	if err != nil {
		fmt.Println("Write file err =", err)
		return
	}
	fmt.Println("Write file success, n =", n)
	//读取文件
	fileContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Read file err =", err)
		return
	}
	fmt.Println("Read file success =", string(fileContent))
}

/*
func WriteFile(filename string, data []byte, perm os.FileMode) error

参数	描述
filename	文件路径。
data	要写入的文件内容。
perm	文件权限。

返回值	描述
err	写入失败，则返回错误信息。

说明
使用 WriteFile 方法写文件，接受的第一个 参数 是一个 string 类型 的文件名，第二个参数是一个要写入的文件内容的 byte 数组，最后一个参数是文件的权限。
如果写入成功，返回空的 error 信息，如果写入失败，返回 error 信息，使用 ioutil.WriteFile写文件，在写入文件之前，我们不需要判断文件是否存在，如果文件不存在，会自动创建文件，如果文件存在，则会覆盖原来的内容。
*/
func ioutilWriteFile_case2() {
	var (
		fileName = "D:\\workspace\\go_project\\study\\daydayup\\base\\file\\haicoder.txt"
		content  = "Hello HaiCoder"
		err      error
	)
	// 这个ioutil.WriteFile在go1.16的时候废弃了，因为这个函数仅仅是调用了一下os.WriteFile，所以我们直接使用os.WriteFile即可
	if err = ioutil.WriteFile(fileName, []byte(content), 0666); err != nil {
		fmt.Println("Writefile Error =", err)
		return
	}
	//读取文件
	fileContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Read file err =", err)
		return
	}
	fmt.Println("Read file success =", string(fileContent))
}

/*
file.Write写文件
func (f *File) Write(b []byte) (n int, err error)
参数	描述
f	文件对象。
b	要写入的文件内容。
返回值	描述
n	成功写入的字节数。
err	写入失败，则返回错误信息。
说明:
使用 file.Write 方法写文件，接受的 参数 是一个要写入的文件内容的 字节 数组。如果写入成功，返回成功写入的字节数，如果写入失败，返回 error 信息。
使用此方法在写入文件之前，我们需要判断文件是否存在，如果文件不存在，则需要创建文件。

file.WriteString写文件
func (f *File) WriteString(s string) (n int, err error)
参数	描述
f	文件对象。
s	要写入的文件内容。
返回值	描述
n	成功写入的字节数。
err	写入失败，则返回错误信息。
说明
使用 file.WriteString 方法写文件，接受的参数是一个要写入的文件内容的 字符串。如果写入成功，返回成功写入的字节数，如果写入失败，返回 error 信息。
使用此方法在写入文件之前，我们需要判断文件是否存在，如果文件不存在，则需要创建文件。
*/
func fileWrite_case3() {
	var (
		fileName = "D:\\workspace\\go_project\\study\\daydayup\\base\\file\\haicoder.txt"
		content  = "Hello HaiCoder"
		file     *os.File
		err      error
	)
	//使用追加模式打开文件
	file, err = os.OpenFile(fileName, os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Open file err =", err)
		return
	}
	defer file.Close()
	//写入文件
	//n, err := file.Write([]byte(content))
	n, err := file.Write([]byte(content)[0:4])
	if err != nil {
		fmt.Println("Write file err =", err)
		return
	}
	fmt.Println("Write file success, n =", n)
	//读取文件
	fileContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Read file err =", err)
		return
	}
	fmt.Println("Read file success =", string(fileContent))
}

/*
bufio.Write写文件
func (b *Writer) Write(p []byte) (n int, err error)
参数	描述
b	文件对象。
p	要写入的文件内容。
返回值	描述
n	成功写入的字节数。
err	写入失败，则返回错误信息。
说明
使用 bufio.Write 方法写文件，接受的 参数 是一个要写入的文件内容的 字节 数组。如果写入成功，返回成功写入的字节数，如果写入失败，返回 error 信息。
使用此方法在写入文件之前，我们需要判断文件是否存在，如果文件不存在，则需要创建文件。

bufio.WriteString写文件
func (b *Writer) WriteString(s string) (int, error)
*/
func bufioWrite_case4() {
	var (
		fileName = "D:\\workspace\\go_project\\study\\daydayup\\base\\file\\haicoder.txt"
		content  = "==Hello HaiCoder"
		file     *os.File
		err      error
	)
	//使用追加模式打开文件
	file, err = os.OpenFile(fileName, os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Open file err =", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	//写入文件
	n, err := writer.Write([]byte(content))
	if err != nil {
		fmt.Println("Write file err =", err)
		return
	}
	fmt.Println("Write file success, n =", n)
	writer.Flush()
	//读取文件
	fileContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Read file err =", err)
		return
	}
	fmt.Println("Read file success =", string(fileContent))
}
