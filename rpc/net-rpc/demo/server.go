package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Args struct {
	X, Y int
}

type ServiceA struct{}

// Add 为ServiceA类型增加一个可导出的Add方法
func (s *ServiceA) Add(args *Args, reply *int) error {
	*reply = args.X + args.Y
	return nil
}

func main() {
	service := new(ServiceA)
	err := rpc.Register(service) // 注册RPC服务
	if err != nil {
		return
	}
	rpc.HandleHTTP() // 基于HTTP协议
	l, e := net.Listen("tcp", ":9091")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	http.Serve(l, nil)
}
