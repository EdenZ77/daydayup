实现一个简单的内网穿透demo
参考资料：https://cloud.tencent.com/developer/article/2090539

## 1. net包
net.ListenTCP 和 net.DialTCP 是 Go 语言标准库 net 包中的两个函数，分别用于监听 TCP 网络地址和发起 TCP 连接。在实现内网穿透时，你可能会用到这两个函数，尤其是当你需要更细粒度的控制 TCP 连接时。
### net.ListenTCP
net.ListenTCP 函数用于在本地网络地址上监听来自 TCP 网络的连接请求。它返回一个 *net.TCPListener 类型的对象，可用于接受连接。ListenTCP 允许更细粒度的控制，比如设置 socket 选项。

语法:
```go
func ListenTCP(network string, laddr *net.TCPAddr) (*net.TCPListener, error)
```
示例代码：
```go
laddr, err := net.ResolveTCPAddr("tcp", ":8080")
if err != nil {
    log.Fatal(err)
}

listener, err := net.ListenTCP("tcp", laddr)
if err != nil {
    log.Fatal(err)
}

for {
    conn, err := listener.AcceptTCP()
    if err != nil {
        log.Println(err)
        continue
    }
    go handleClient(conn)
}
```

### net.DialTCP
net.DialTCP 函数用于发起到指定 TCP 网络地址的连接。它返回一个 *net.TCPConn 类型的对象，可用于读写数据。与 Dial 函数不同，DialTCP 提供了底层 TCP 连接的额外控制能力。

语法:
```go
func DialTCP(network string, laddr, raddr *net.TCPAddr) (*net.TCPConn, error)

```
示例代码：
```go
raddr, err := net.ResolveTCPAddr("tcp", "server.example.com:8080")
if err != nil {
    log.Fatal(err)
}

conn, err := net.DialTCP("tcp", nil, raddr)
if err != nil {
    log.Fatal(err)
}

// 使用conn进行读写操作

```
### 内网穿透上下文中的使用
在内网穿透的上下文中，net.ListenTCP 可能会用在服务端上，监听对应的端口以便接受来自客户端的持久连接和外部用户的连接请求。而 net.DialTCP 可能会用在客户端上，用于主动发起到服务端的连接。

使用 net.ListenTCP 和 net.DialTCP 的好处在于，你可以更精确地控制 TCP 连接参数，例如设置超时、keep-alive 选项等。这对于长时间运行的连接，例如内网穿透场景中的连接，是非常有用的。

请注意，对于大多数简单的应用场景，net.Listen 和 net.Dial 提供的功能已经足够。只有当你需要额外的控制并理解 TCP 协议的细节时，才有必要使用 ListenTCP 和 DialTCP。