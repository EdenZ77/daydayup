package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"hello/common_lib/grpc-gateway/gen/go/protobuf/admin/v1"
)

func main() {
	// 连接服务器
	conn, err := grpc.Dial(":8972", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("faild to connect: %v", err)
	}
	defer conn.Close()

	c := servicev1.NewGreeterClient(conn)
	// 调用服务端的SayHello
	r, err := c.SayHello(context.Background(), &servicev1.HelloRequest{Name: "q1mi"})
	if err != nil {
		fmt.Printf("could not greet: %v", err)
	}
	fmt.Printf("Greeting: %s !\n", r.Message)
}
