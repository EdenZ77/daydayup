package network

import (
	"io"
	"log"
	"net"
)

const (
	KeepAlive     = "KEEP_ALIVE"
	NewConnection = "NEW_CONNECTION"
)

func CreateTCPListener(addr string) (*net.TCPListener, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return nil, err
	}
	tcpListener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		return nil, err
	}
	return tcpListener, nil
}

func CreateTCPConn(addr string) (*net.TCPConn, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return nil, err
	}
	tcpListener, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return nil, err
	}
	return tcpListener, nil
}

// Join2Conn
/*
这段代码基本上是在两个 TCP 连接之间建立了一个简单的数据桥接。当你想要让两个 TCP 端点之间互相通信时，
比如在构建代理服务时，这段代码可以实现数据的双向传输。数据从 local 连接读取，然后写入 remote 连接，同时也从 remote 连接读取数据写入 local 连接。
这样可以实现数据的双向流动，从而连接 local 和 remote 两端的通信。

这种模式也常用于内网穿透中，允许位于内网的机器与外界建立通信。
比如，你可以将 local 端点连接到内网的一个服务，而 remote 端点连接到外部服务器，从而允许通过外部服务器中转来访问内网服务。
*/
func Join2Conn(local *net.TCPConn, remote *net.TCPConn) {
	// 如果你只启动一个任务或者在一个单一的线程/协程中串行处理这两种传输，那么在一个方向上的数据传输必须等待另一个方向上的数据传输完成，
	// 这会导致延迟和阻塞，使得两个连接中的一方面可能会等待很长时间来接收或发送数据。
	// 每一个 joinConn 调用都在它自己的 goroutine 中运行，这意味着 local 到 remote 的数据流和 remote 到 local 的数据流可以同时发生，各自不受对方影响。
	go joinConn(local, remote)
	go joinConn(remote, local)
}

// 别看这短短几行代码，这就是核心了
func joinConn(local *net.TCPConn, remote *net.TCPConn) {
	defer local.Close()
	defer remote.Close()
	_, err := io.Copy(local, remote)
	if err != nil {
		log.Println("copy failed ", err.Error())
		return
	}
}
