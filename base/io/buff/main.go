package main

import (
	"bytes"
	"fmt"
)

func main() {
	//// 创建一个空的缓冲区
	//var buf bytes.Buffer
	//// 向缓冲区写入字符串
	//buf.WriteString("Hello, World!")
	//// 从缓冲区读取字符串
	//fmt.Println(buf.String()) // 输出：Hello, World!

	//buf := bytes.NewBufferString("hello world")
	//fmt.Println(buf.String()) // 输出：hello world

	//var buf bytes.Buffer
	//buf.WriteString("hello")
	//buf.WriteString(" ")
	//buf.WriteString("world")
	//fmt.Println(buf.String()) // 输出：hello world

	// 3. 写入数据
	//buf := bytes.NewBuffer(nil)
	//n, err := buf.Write([]byte("hello world"))
	//if err != nil {
	//	fmt.Println("write error:", err)
	//}
	//fmt.Printf("write %d bytes\n", n) // 输出：write 11 bytes
	//fmt.Println(buf.String())         // 输出：hello world

	// 4. 读取数据
	//buf := bytes.NewBufferString("hello world")
	//data := make([]byte, 5)
	//n, err := buf.Read(data)
	//if err != nil {
	//	fmt.Println("read error:", err)
	//}
	//fmt.Printf("read %d bytes\n", n) // 输出：read 5 bytes
	//fmt.Println(string(data))        // 输出：hello

	// 5. 截取缓冲区
	//buf := bytes.NewBufferString("hello world")
	//buf.Truncate(5)
	//fmt.Println(buf.String()) // 输出：hello

	// 6. 扩容缓冲区
	//buf := bytes.NewBufferString("hello")
	//buf.Grow(10)
	//fmt.Printf("len=%d, cap=%d\n", buf.Len(), buf.Cap()) // 输出：len=5, cap=16

	// 7. 重置缓冲区
	//buf := bytes.NewBufferString("hello")
	//fmt.Println(buf.String()) // 输出：hello
	//buf.Reset()
	//fmt.Println(buf.String()) // 输出：

	// 8. 序列化和反序列化
	//type Person struct {
	//	Name string
	//	Age  int
	//}
	//// 将结构体编码为 JSON
	//p := Person{"Alice", 25}
	////buf := bytes.NewBuffer(nil) 两种声明buf的方式都可以
	//var buf bytes.Buffer
	//enc := json.NewEncoder(&buf)
	//enc.Encode(p)
	//fmt.Println(buf.String()) // 输出：{"Name":"Alice","Age":25}
	//
	//// 从 JSON 解码为结构体
	//var p2 Person
	//dec := json.NewDecoder(&buf)
	//dec.Decode(&p2)
	//fmt.Printf("Name: %s, Age: %d\n", p2.Name, p2.Age) // 输出：Name: Alice, Age: 25

	// 9. Buffer的应用场景

	// 9.2 文件操作
	//// 从文件中读取数据
	//file, err := os.Open("D:\\workspace\\go_project\\study\\daydayup\\base\\buff\\example.txt")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer file.Close()
	//var buf bytes.Buffer
	//_, err = io.Copy(&buf, file)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(buf.String())
	//// 将数据写入文件
	//out, err := os.Create("output.txt")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer out.Close()
	//_, err = io.Copy(out, &buf)
	//if err != nil {
	//	log.Fatal(err)
	//}

	// 9.3 二进制数据处理
	//// 读取字节数组
	//data := []byte{0x48, 0x65, 0x6c, 0x6c, 0x6f}
	//var buf bytes.Buffer
	//buf.Write(data)
	//// 转换大小端序
	//var num uint16
	//binary.Read(&buf, binary.BigEndian, &num)
	//fmt.Println(num) // 输出：0x4865
	//// 写入字节数组
	//data2 := []byte{0x57, 0x6f, 0x72, 0x6c, 0x64, 0x21}
	//buf.Write(data2)
	//fmt.Println(buf.Bytes()) // 输出：[72 101 108 108 111 87 111 114 108 100 33]

	// 9.4 字符串拼接
	//s1 := "hello"
	//s2 := "world"
	//s3 := "!"
	//s := concatStrings(s1, s2, s3)
	//fmt.Println(s) // 输出：hello world!

	// 9.5 格式化输出
	var buf bytes.Buffer
	for i := 0; i < 10; i++ {
		fmt.Fprintf(&buf, "%d", i)
	}
	fmt.Println(buf.String())
}

/*
我们使用 Buffer 类型将多个字符串拼接成一个字符串。由于 Buffer 类型会动态扩容，因此可以避免产生大量的中间变量，提高程序的效率。
*/
func concatStrings(strs ...string) string {
	var buf bytes.Buffer
	for _, s := range strs {
		buf.WriteString(s)
	}
	return buf.String()
}
