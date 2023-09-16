package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"hello/network/rpc/grpc/hello/authentication"
	"hello/network/rpc/grpc/hello/pb"
	"io/ioutil"
	"log"
	"time"

	"google.golang.org/grpc/credentials"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc/metadata"

	"google.golang.org/grpc"
)

// hello_client

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "127.0.0.1:8972", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

const (
	myScheme      = "17x"
	myServiceName = "resolver.17x.lixueduan.com"
)

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	// 指定客户端interceptor
	opts = append(opts, grpc.WithChainUnaryInterceptor(interceptor2, interceptor1))
	// 客户端通过ca证书来验证服务的提供的证书 这是演示服务端TLS=====================
	//creds, err := credentials.NewClientTLSFromFile("rpc/grpc/hello/cert/ca.crt", "www.lixueduan.com")
	//if err != nil {
	//	log.Fatalf("failed to load credentials: %v", err)
	//}
	//====================================

	// 构建一个自定义的PerRPCCredentials。
	myAuth := authentication.NewMyAuth()

	// 加载客户端证书 =================
	certificate, err := tls.LoadX509KeyPair("rpc/grpc/hello/cert/client.crt", "rpc/grpc/hello/cert/client.key")
	if err != nil {
		log.Fatal(err)
	}
	// 构建CertPool以校验服务端证书有效性
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("rpc/grpc/hello/cert/ca.crt")
	if err != nil {
		log.Fatal(err)
	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatal("failed to append ca certs")
	}

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{certificate},
		ServerName:   "www.lixueduan.com", // NOTE: this is required!
		RootCAs:      certPool,
	})
	// ==============

	opts = append(opts, grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(myAuth))
	// 连接到server端
	conn, err := grpc.Dial(*addr, opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := xx.NewGreeterClient(conn)

	// 创建带有metadata的context，带有超时控制
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	md := metadata.Pairs("k1", "v1", "k1", "v2", "k2", "v3")
	ctxWithMD := metadata.NewOutgoingContext(ctx, md)

	var header, trailer metadata.MD // 声明存储header和trailer的变量

	r, err := c.SayHello(ctxWithMD, &xx.HelloRequest{Name: *name}, grpc.Header(&header), grpc.Trailer(&trailer))
	// 当服务端返回错误时，尝试从错误中获取detail信息
	if err != nil {
		s := status.Convert(err)        // 将err转为status
		for _, d := range s.Details() { // 获取details
			// 错误类型断言
			switch info := d.(type) {
			case *errdetails.QuotaFailure:
				fmt.Printf("Quota failure: %s\n", info)
			default:
				fmt.Printf("Unexpected type: %s\n", info)
			}
		}
		fmt.Printf("c.SayHello failed, err:%v\n", err)
		return
	}
	log.Printf("Greeting: %s", r.GetReply()) // Greeting: Hello world
	log.Println(header, trailer)             // map[content-type:[application/grpc] header-key:[header-val]] map[trailer-key:[trailer-val]]
}

// interceptor 客户端拦截器
func interceptor1(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	fmt.Printf("1 before client method=%s req=%v rep=%v \n", method, req, reply)
	start := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)
	fmt.Printf("1 after client method=%s req=%v rep=%v duration=%s error=%v\n", method, req, reply, time.Since(start), err)
	return err
}

func interceptor2(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	fmt.Printf("2 before client method=%s req=%v rep=%v \n", method, req, reply)
	start := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)
	fmt.Printf("2 after client method=%s req=%v rep=%v duration=%s error=%v\n", method, req, reply, time.Since(start), err)
	return err
}
