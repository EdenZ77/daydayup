package main

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// 读取请求
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading request:", err)
		return
	}
	req := string(buf[:n])
	fmt.Println("=============")
	fmt.Println(req)

	// 构建响应
	resp := "HTTP/1.1 301 Moved Permanently\r\nLocation: http://www.baidu.com\r\n\r\n"

	// 发送响应
	_, err = conn.Write([]byte(resp))
	if err != nil {
		fmt.Println("Error writing response:", err)
		return
	}
}

func main() {
	ln, err := net.Listen("tcp", ":8082")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}
