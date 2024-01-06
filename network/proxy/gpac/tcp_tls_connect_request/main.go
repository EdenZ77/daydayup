package main

import (
	"crypto/tls"
	"fmt"
)

const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkb21haW5OYW1lIjoieHh4IiwiZXhwaXJlVGltZSI6MTcxMDgzOTEwNSwidGVuYW50SWQiOjEsInVzZXJJZCI6MSwidXNlcm5hbWUiOiJhZG1pbiJ9.a29rkFkSGmKiv5vA7aeKMN1HCL5EiQcXBxBa0PJSpGQ"

// 模拟发送CONNECT请求
func main() {
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true, // 跳过证书验证（仅用于示例，请勿在生产环境使用）
	}
	proxy_conn, err := tls.Dial("tcp", "172.30.5.95:8089", tlsConfig)
	if err != nil {
		fmt.Println("连接到代理服务器失败:", err)
		return
	}
	fmt.Println(proxy_conn)
	//connectReq := &http.Request{
	//	Method: http.MethodConnect,
	//	URL:    &url.URL{Opaque: "www.baidu.com:443"},
	//	Host:   "www.baidu.com:443",
	//	Header: make(http.Header),
	//}
	//connectReq.Header.Set("Authorization", token)
	//connectReq.Write(proxy_conn)
	//br := bufio.NewReader(proxy_conn)
	//resp, err := http.ReadResponse(br, connectReq)
	//if err != nil {
	//	proxy_conn.Close()
	//}
	//fmt.Println(resp)
}
