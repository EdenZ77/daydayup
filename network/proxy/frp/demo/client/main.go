package main

import (
	"bufio"
	"hello/network/proxy/frp/demo/network"
	"io"
	"log"
	"net"
)

/*
我们先来实现相对简单的客户端，客户端主要做的事情是 3 件：

1、连接服务端的控制通道
2、等待服务端从控制通道中发来建立连接的消息
3、收到建立连接的消息时，将本地服务和远端隧道建立连接（这里就要用到我们的工具方法了）
*/
var (
	// 本地需要暴露的服务端口
	localServerAddr = "127.0.0.1:32768"

	remoteIP = "111.111.111.111"
	// 远端的服务控制通道，用来传递控制信息，如出现新连接和心跳
	remoteControlAddr = remoteIP + ":8009"
	// 远端服务端口，用来建立隧道
	remoteServerAddr = remoteIP + ":8008"
)

func main() {
	// 服务器端的控制通道需要有对应的服务端程序来发送新连接的指令。同时，要确保网络安全性，控制通道应该仅对信任的客户端开放。
	tcpConn, err := network.CreateTCPConn(remoteControlAddr)
	if err != nil {
		log.Println("[连接失败]" + remoteControlAddr + err.Error())
		return
	}
	log.Println("[已连接]" + remoteControlAddr)

	reader := bufio.NewReader(tcpConn)
	for {
		s, err := reader.ReadString('\n')
		if err != nil || err == io.EOF {
			break
		}

		// 每次接收到 network.NewConnection 消息时，connectLocalAndRemote 都在新的 goroutine 中调用，因此每个请求都会产生各自的本地和远程连接。
		// 这意味着如果服务端连续发送多个 network.NewConnection 请求，客户端将为每个请求创建一个新的隧道连接。这些隧道是独立的，允许多个用户的请求通过客户端转发到本地服务。
		// 每个隧道的一端是客户端所在机器上的本地服务地址 127.0.0.1:32768，另一端是服务端的远端服务地址 111.111.111.111:8008。
		if s == network.NewConnection+"\n" {
			go connectLocalAndRemote()
		}
	}

	log.Println("[已断开]" + remoteControlAddr)
}

func connectLocalAndRemote() {
	local := connectLocal()
	remote := connectRemote()

	if local != nil && remote != nil {
		network.Join2Conn(local, remote)
	} else {
		if local != nil {
			_ = local.Close()
		}
		if remote != nil {
			_ = remote.Close()
		}
	}
}

func connectLocal() *net.TCPConn {
	conn, err := network.CreateTCPConn(localServerAddr)
	if err != nil {
		log.Println("[连接本地服务失败]" + err.Error())
	}
	return conn
}

// 此方法将连接8008端口的远端服务
func connectRemote() *net.TCPConn {
	conn, err := network.CreateTCPConn(remoteServerAddr)
	if err != nil {
		log.Println("[连接远端服务失败]" + err.Error())
	}
	return conn
}
