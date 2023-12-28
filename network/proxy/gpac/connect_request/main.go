package main

import (
	"bufio"
	"fmt"
	"net"
)

// 模拟发送CONNECT请求
func main() {
	// 建立到代理服务器的连接
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("连接到代理服务器失败:", err)
		return
	}
	defer conn.Close()

	// 发送CONNECT请求
	fmt.Fprintf(conn, "CONNECT www.example.com:443 HTTP/1.1\r\nHost: www.example.com\r\n\r\n")

	// 读取代理服务器的响应
	response, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("读取响应失败:", err)
		return
	}

	fmt.Println("代理服务器响应:", response)

	// 在这里可以继续进行TLS握手和其他操作，以建立一个HTTPS连接
	// 但为了简单起见，本示例只发送CONNECT请求并读取响应
}
