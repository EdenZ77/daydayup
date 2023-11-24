package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	// 监听本地指定端口
	listener, err := net.Listen("tcp", "127.0.0.1:1080")
	if err != nil {
		log.Fatalf("Failed to set up listener: %v", err)
	}
	defer listener.Close()

	log.Println("SOCKS5 server listening on 127.0.0.1:1080")

	for {
		clientConn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept client connection: %v", err)
			continue
		}
		go handleClient(clientConn)
	}
}

func handleClient(clientConn net.Conn) {
	defer clientConn.Close()

	// 握手阶段
	if err := socks5Handshake(clientConn); err != nil {
		log.Printf("SOCKS5 handshake failed: %v", err)
		return
	}

	// 获取客户端请求
	targetConn, err := socks5GetRequest(clientConn)
	if err != nil {
		log.Printf("SOCKS5 request failed: %v", err)
		return
	}
	defer targetConn.Close()

	// 进行数据转发
	go io.Copy(targetConn, clientConn)
	io.Copy(clientConn, targetConn)
}

func socks5Handshake(conn net.Conn) error {
	buf := make([]byte, 2)

	// 读取版本号和认证方法数量
	if _, err := io.ReadAtLeast(conn, buf, 2); err != nil {
		return err
	}
	version, nmethods := buf[0], buf[1]

	// 只支持SOCKS5
	if version != 0x05 {
		return io.ErrUnexpectedEOF
	}

	// 读取支持的认证方法
	methods := make([]byte, nmethods)
	if _, err := io.ReadAtLeast(conn, methods, int(nmethods)); err != nil {
		return err
	}

	// 检查是否支持无认证方式
	for _, method := range methods {
		if method == 0x00 {
			// 发送无认证响应
			_, err := conn.Write([]byte{0x05, 0x00})
			return err
		}
	}

	// 不支持的认证方法
	return io.ErrUnexpectedEOF
}

func socks5GetRequest(conn net.Conn) (net.Conn, error) {
	buf := make([]byte, 4)

	// 读取版本号、命令和保留字段
	if _, err := io.ReadAtLeast(conn, buf, 4); err != nil {
		return nil, err
	}
	version, cmd, _ := buf[0], buf[1], buf[2]

	// 只支持SOCKS5
	if version != 0x05 {
		return nil, io.ErrUnexpectedEOF
	}

	// 只处理TCP连接请求
	if cmd != 0x01 {
		return nil, io.ErrUnexpectedEOF
	}

	// 获取目标地址
	addrType := make([]byte, 1)
	if _, err := conn.Read(addrType); err != nil {
		return nil, err
	}

	var host string
	switch addrType[0] {
	case 0x01: // IPv4
		ipv4 := make(net.IP, net.IPv4len)
		conn.Read(ipv4)
		host = ipv4.String()
	case 0x03: // 域名
		length := make([]byte, 1)
		conn.Read(length)
		domain := make([]byte, length[0])
		conn.Read(domain)
		host = string(domain)
	case 0x04: // IPv6
		ipv6 := make(net.IP, net.IPv6len)
		conn.Read(ipv6)
		host = ipv6.String()
	default:
		return nil, io.ErrUnexpectedEOF
	}

	// 获取端口号
	port := make([]byte, 2)
	if _, err := io.ReadAtLeast(conn, port, 2); err != nil {
		return nil, err
	}

	// 创建到目标地址的连接
	destAddr := net.JoinHostPort(host, fmt.Sprintf("%d", binary.BigEndian.Uint16(port)))
	destConn, err := net.Dial("tcp", destAddr)
	if err != nil {
		return nil, err
	}

	// 发送成功响应
	resp := []byte{
		0x05,        // 版本号
		0x00,        // 响应：成功
		0x00,        // 保留
		addrType[0], // 地址类型
	}
	resp = append(resp, addrType...)
	resp = append(resp, port...)
	if _, err := conn.Write(resp); err != nil {
		destConn.Close()
		return nil, err
	}

	return destConn, nil
}
