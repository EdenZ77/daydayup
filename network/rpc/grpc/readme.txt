https://www.liwenzhou.com/posts/Go/gRPC/

gRPC是一种现代化开源的高性能RPC框架。最初由谷歌进行开发。它使用HTTP/2作为传输协议。
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2


像许多 RPC 系统一样，gRPC 基于定义服务的思想，指定可以通过参数和返回类型远程调用的方法。
默认情况下，gRPC 使用 protocol buffers作为接口定义语言(IDL)来描述服务接口和有效负载消息的结构。可以根据需要使用其他的IDL代替

gRPC 服务类型：
在gRPC中你可以定义四种类型的服务方法。
    普通 rpc，客户端向服务器发送一个请求，然后得到一个响应，就像普通的函数调用一样。
          rpc SayHello(HelloRequest) returns (HelloResponse);
    服务器流式 rpc，其中客户端向服务器发送请求，并获得一个流来读取一系列消息。客户端从返回的流中读取，直到没有更多的消息。gRPC 保证在单个 RPC 调用中的消息是有序的。
        rpc LotsOfReplies(HelloRequest) returns (stream HelloResponse);
    客户端流式 rpc，其中客户端写入一系列消息并将其发送到服务器，同样使用提供的流。一旦客户端完成了消息的写入，它就等待服务器读取消息并返回响应。同样，gRPC 保证在单个 RPC 调用中对消息进行排序。
        rpc LotsOfGreetings(stream HelloRequest) returns (HelloResponse);
    双向流式 rpc，其中双方使用读写流发送一系列消息。这两个流独立运行，因此客户端和服务器可以按照自己喜欢的顺序读写: 例如，服务器可以等待接收所有客户端消息后再写响应，或者可以交替读取消息然后写入消息，或者其他读写组合。每个流中的消息是有序的。


gRPC metadata：
元数据（metadata）是指在处理RPC请求和响应过程中需要但又不属于具体业务（例如身份验证详细信息）的信息，采用键值对列表的形式，其中键是string类型，值通常是[]string类型，但也可以是二进制数据。
gRPC中的 metadata 类似于我们在 HTTP headers中的键值对，元数据可以包含认证token、请求标识和监控标签等。
    测试：在服务端发送、接收metadata数据；在客户端发送、接收metadata数据


gRPC 错误处理：
查考示例代码即可



gRPC 加密或认证：
先回顾https和证书相关内容




gRPC 拦截器：
gRPC 为在每个 ClientConn/Server 基础上实现和安装拦截器提供了一些简单的 API。
客户端端拦截器——普通拦截器(一元拦截器)
    一元拦截器的实现通常可以分为三个部分: 调用 RPC 方法之前（预处理）、调用 RPC 方法（RPC调用）和调用 RPC 方法之后（调用后）。

    预处理：用户可以通过检查传入的参数(如 RPC 上下文、方法字符串、要发送的请求和 CallOptions 配置)来获得有关当前 RPC 调用的信息。
    RPC调用：预处理完成后，可以通过执行invoker执行 RPC 调用。
    调用后：一旦调用者返回应答和错误，用户就可以对 RPC 调用进行后处理。通常，它是关于处理返回的响应和错误的。 若要在 ClientConn 上安装一元拦截器，请使用DialOptionWithUnaryInterceptor的DialOption配置 Dial 。

server端拦截器——普通拦截器(一元拦截器)
    服务器端拦截器与客户端类似，但提供的信息略有不同。


gRPC 自定义身份校验

gRPC 负载均衡：







