package main

import (
	"fmt"
	"log"
	"net"
)

// http分块传输
// 参考资料：https://www.bwangel.me/2018/11/01/http-chunked/

// 短连接
func handleClosedHttpReq(conn net.Conn) {
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("=============")
	fmt.Println(n, string(buffer))

	data := []byte("hello, world!")
	conn.Write([]byte("HTTP/1.1 200 OK\r\n"))
	conn.Write([]byte("\r\n"))
	conn.Write(data)
	// 关闭连接
	conn.Close()
}

// 没有添加 Content-Length 的响应，浏览器会处于 pending 的状态
func handleKeepAliveHttpReqNoContextLength(conn net.Conn) {
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("=============")
	fmt.Println(n, string(buffer))

	data := []byte("hello, keel-alive!")
	conn.Write([]byte("HTTP/1.1 200 OK\r\n"))
	conn.Write([]byte("\r\n"))
	conn.Write(data)
}

// 添加了 Content-Length，浏览器就可以正常处理响应了
func handleKeepAliveHttpReqAddContextLength(conn net.Conn) {
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(n, string(buffer))

	data := []byte("hello, keel-alive!")
	// 状态行
	conn.Write([]byte("HTTP/1.1 200 OK\r\n"))
	// 响应头
	conn.Write([]byte(fmt.Sprintf("Content-Length: %d\r\n", len(data))))
	// 空行
	conn.Write([]byte("\r\n"))
	// 响应体
	conn.Write(data)
}

// 分块传输编码
func handleChunkedHttpResp(conn net.Conn) {
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(n, string(buffer))

	conn.Write([]byte("HTTP/1.1 200 OK\r\n"))
	conn.Write([]byte("Transfer-Encoding: chunked\r\n"))
	conn.Write([]byte("\r\n"))

	conn.Write([]byte("6\r\n"))
	conn.Write([]byte("hello,\r\n"))

	conn.Write([]byte("8\r\n"))
	conn.Write([]byte("chunked!\r\n"))

	conn.Write([]byte("0\r\n"))
	conn.Write([]byte("\r\n"))
}

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleChunkedHttpResp(conn)
	}
}
