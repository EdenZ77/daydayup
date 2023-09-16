package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"hello/network/rpc/grpc/hello/authentication"
	"hello/network/rpc/grpc/hello/pb"
	"io/ioutil"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc/metadata"

	"google.golang.org/grpc"
)

// hello server

type server struct {
	xx.UnimplementedGreeterServer

	mu    sync.Mutex     // count的并发锁
	count map[string]int // 记录每个name的请求次数
}

func (s *server) SayHello(ctx context.Context, in *xx.HelloRequest) (*xx.HelloResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// 一般通过defer来设置trailer. 创建和发送 trailer
	defer func() {
		trailer := metadata.Pairs("trailer-key", "trailer-val")
		err := grpc.SetTrailer(ctx, trailer)
		if err != nil {
			return
		}
	}()

	//time.Sleep(2 * time.Second)
	// 接收metadata
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		// 输出：map[:authority:[127.0.0.1:8972] content-type:[application/grpc] k1:[v1 v2] k2:[v3] user-agent:[grpc-go/1.46.0 ]]
		log.Println(md)
	}

	// 创建和发送 header
	header := metadata.Pairs("header-key", "header-val")
	err := grpc.SendHeader(ctx, header)
	if err != nil {
		return nil, err
	}
	// 创建和发送 trailer 一般放在最后
	//trailer := metadata.Pairs("trailer-key", "trailer-val")
	//err = grpc.SetTrailer(ctx, trailer)
	//if err != nil {
	//	return nil, err
	//}

	s.count[in.Name]++ // 记录用户的请求次数
	// 超过1次就返回错误
	if s.count[in.Name] > 1 {
		st := status.New(codes.ResourceExhausted, "Request limit exceeded.")
		ds, err := st.WithDetails(
			&errdetails.QuotaFailure{
				Violations: []*errdetails.QuotaFailure_Violation{{
					Subject:     fmt.Sprintf("name:%s", in.Name),
					Description: "限制每个name调用一次",
				}},
			},
		)
		// 在构造更多细节的过程中发生了错误，则直接返回最初的错误信息即可
		if err != nil {
			return nil, st.Err()
		}
		// 构造错误细节成功之后，返回构造后的错误
		return nil, ds.Err()
	}
	// 正常返回响应
	reply := "hello " + in.GetName()
	return &xx.HelloResponse{Reply: reply}, nil
}

func main() {
	// 监听本地的8972端口
	lis, err := net.Listen("tcp", ":8972")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}

	// 指定使用服务端证书创建一个 TLS credentials。 这是演示服务端TLS=====================
	//creds, err := credentials.NewServerTLSFromFile("rpc/grpc/hello/cert/server.crt", "rpc/grpc/hello/cert/server.key")
	//if err != nil {
	//	log.Fatalf("failed to create credentials: %v", err)
	//}
	//============================

	// 从证书相关文件中读取和解析信息，得到证书公钥、密钥对     这是演示双向认证=================
	certificate, err := tls.LoadX509KeyPair("rpc/grpc/hello/cert/server.crt", "rpc/grpc/hello/cert/server.key")
	if err != nil {
		log.Fatal(err)
	}
	// 创建CertPool，后续就用池里的证书来校验客户端证书有效性
	// 所以如果有多个客户端 可以给每个客户端使用不同的 CA 证书，来实现分别校验的目的
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("rpc/grpc/hello/cert/ca.crt")
	if err != nil {
		log.Fatal(err)
	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatal("failed to append certs")
	}
	// 构建基于 TLS 的 TransportCredentials
	creds := credentials.NewTLS(&tls.Config{
		// 设置证书链，允许包含一个或多个
		Certificates: []tls.Certificate{certificate},
		// 要求必须校验客户端的证书 可以根据实际情况选用其他参数
		ClientAuth: tls.RequireAndVerifyClientCert, // NOTE: this is optional!
		// 设置根证书的集合，校验方式使用 ClientAuth 中设定的模式
		ClientCAs: certPool,
	})
	//=================

	// 创建gRPC服务器===============================
	//s := grpc.NewServer()
	s := grpc.NewServer(grpc.ChainUnaryInterceptor(LoggingInterceptor1, LoggingInterceptor2, myEnsureValidToken), grpc.Creds(creds))
	// 注册服务，注意初始化count
	xx.RegisterGreeterServer(s, &server{count: make(map[string]int)})
	// 启动服务
	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
}

// LoggingInterceptor1 多拦截器处理
// LoggingInterceptor1  拦截器 - 打印日志
func LoggingInterceptor1(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	fmt.Printf("1 before server gRPC method: %s, %v\n", info.FullMethod, req)
	resp, err := handler(ctx, req)
	fmt.Printf("1 after server gRPC method: %s, %v\n", info.FullMethod, resp)
	return resp, err
}

func LoggingInterceptor2(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	fmt.Printf("2 before server gRPC method: %s, %v\n", info.FullMethod, req)
	resp, err := handler(ctx, req)
	fmt.Printf("2 after server gRPC method: %s, %v\n", info.FullMethod, resp)
	return resp, err
}

// myEnsureValidToken 自定义 token 校验
func myEnsureValidToken(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// 如果返回err不为nil则说明token验证未通过
	err := authentication.IsValidAuth(ctx)
	if err != nil {
		return nil, err
	}
	return handler(ctx, req)
}
