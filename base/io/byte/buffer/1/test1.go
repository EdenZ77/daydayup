package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	testInit()
	//testWrite()
	//testWriteString()
	//testRead()
	//testReadFrom()
}

func testInit() {
	//第一种定义方式
	var b bytes.Buffer        //直接定义一个 Buffer 变量，而不用初始化
	b.Write([]byte("Hello ")) // 可以直接使用
	fmt.Println(b.String())
	//第二种定义方式
	c := new(bytes.Buffer)
	c.WriteString("World")
	fmt.Println(c)
	//第三种定义方式
	d := bytes.NewBuffer(nil)
	d.WriteString("这是第三种定义方式")
	fmt.Println(d.String())
	//第四张定义方式
	f := bytes.NewBufferString("这是第四种定义方式")
	fmt.Println(f.String())
}

func testWrite() {
	newBytes := []byte(" go")
	//创建一个内容Learning的缓冲器
	buf := bytes.NewBuffer([]byte("Learning"))
	//将newBytes这个slice写到buf的尾部
	buf.Write(newBytes)
	fmt.Println(buf.String())
}

func testWriteString() {
	newString := " go"
	//创建一个string内容Learning的缓冲器
	buf := bytes.NewBufferString("Learning")
	//将newString这个string写到buf的尾部
	buf.WriteString(newString)
	fmt.Println(buf.String())
}

func testRead() {
	bufs := bytes.NewBufferString("Learning swift.")
	fmt.Println("缓冲器：" + bufs.String())
	l := make([]byte, 5)
	//把bufs的内容读入到l内,因为l容量为5,所以只读了5个过来
	bufs.Read(l)
	fmt.Println("读取到的内容：" + string(l))
	fmt.Println("缓冲器：" + bufs.String())
}

func testReadFrom() {
	file, _ := os.Open("D:\\workspace\\go_project\\daydayup\\study01\\base\\byte\\buffer\\1\\a.txt")
	buf := bytes.NewBufferString("Learning swift.")
	buf.ReadFrom(file) //将text.txt内容追加到缓冲器的尾部
	fmt.Println(buf.String())
}

func testReset() {
	bufs := bytes.NewBufferString("现在开始 Learning go.")
	bufs.Reset()
	fmt.Println(bufs.String())
}
