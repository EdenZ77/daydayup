package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
)

func main() {
	// 监听本地端口
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	fmt.Println("代理服务器已启动，监听端口 8080")

	for {
		// 接受客户端连接
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		// 处理客户端请求
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	// 读取客户端请求
	req, err := http.ReadRequest(bufio.NewReader(conn))
	if err != nil {
		fmt.Println("读取请求失败：", err)
		return
	}

	// 处理 CONNECT 请求
	if req.Method == "CONNECT" {
		handleConnect(conn, req)
		return
	}

	// 处理普通的 HTTP 请求
	handleHTTP(conn, req)
}

func handleConnect(conn net.Conn, req *http.Request) {
	// 解析目标服务器的主机名和端口号
	targetAddr := req.URL.Host
	if strings.Index(targetAddr, ":") == -1 {
		targetAddr += ":443"
	}

	// 建立到目标服务器的 TCP 连接
	targetConn, err := net.Dial("tcp", targetAddr)
	if err != nil {
		fmt.Println("连接目标服务器失败：", err)
		return
	}
	defer targetConn.Close()

	// 返回客户端 200 OK 响应
	_, err = fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n\r\n")
	if err != nil {
		fmt.Println("发送响应失败：", err)
		return
	}

	// 转发客户端和目标服务器之间的数据
	go func() {
		_, err = io.Copy(targetConn, conn)
		if err != nil {
			fmt.Println("转发数据失败：", err)
			return
		}
	}()

	go func() {
		_, err = io.Copy(conn, targetConn)
		if err != nil {
			fmt.Println("转发数据失败：", err)
			return
		}
	}()
}

func handleHTTP(conn net.Conn, req *http.Request) {
	// 建立到目标服务器的 TCP 连接
	targetConn, err := net.Dial("tcp", req.Host)
	if err != nil {
		fmt.Println("连接目标服务器失败：", err)
		return
	}
	defer targetConn.Close()

	// 发送客户端请求到目标服务器
	err = req.Write(targetConn)
	if err != nil {
		fmt.Println("发送请求失败：", err)
		return
	}

	// 转发目标服务器的响应到客户端
	resp, err := http.ReadResponse(bufio.NewReader(targetConn), req)
	if err != nil {
		fmt.Println("读取响应失败：", err)
		return
	}
	defer resp.Body.Close()

	err = resp.Write(conn)
	if err != nil {
		fmt.Println("发送响应失败：", err)
		return
	}
}
