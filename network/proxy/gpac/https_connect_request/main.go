package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
)

const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkb21haW5OYW1lIjoieHh4IiwiZXhwaXJlVGltZSI6MTcxMDgzOTEwNSwidGVuYW50SWQiOjEsInVzZXJJZCI6MSwidXNlcm5hbWUiOiJhZG1pbiJ9.a29rkFkSGmKiv5vA7aeKMN1HCL5EiQcXBxBa0PJSpGQ"

func main() {
	// 服务器的地址和端口
	serverAddr := "https://172.30.5.95:8089"

	// 创建一个不检查服务器证书的HTTP客户端（仅用于测试）
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	// 创建一个新的HTTP请求对象
	req, err := http.NewRequest("GET", serverAddr, nil)
	if err != nil {
		fmt.Println("创建请求失败:", err)
		return
	}
	req.Header.Set("Authorization", token)

	// 发起请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("发起HTTPS请求失败:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
