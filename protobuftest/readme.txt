protobuf是一种高效的数据格式，平台无关、语言无关、可扩展，可用于 RPC 系统。目前Protobuf作为接口规范的描述语言，可以作为Go语言RPC接口的基础工具。
protobuf是一个与语言无关的一个数据协议，所以我们需要先编写IDL文件然后借助专用工具生成指定语言的代码，从而实现数据的序列化与反序列化过程。

protobuf协议编译器是用c++编写的，根据自己的操作系统下载对应版本的protoc编译器：https://github.com/protocolbuffers/protobuf/releases，解压后拷贝到GOPATH/bin目录下。
protoc.exe、protoc-gen-go.exe

Protobuf3语言指南：https://colobu.com/2017/03/16/Protobuf3-language-guide/














