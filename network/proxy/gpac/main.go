package main

import (
	"fmt"
	"github.com/darren/gpac"
)

var scripts = `
  function FindProxyForURL(url, host) {
    if (isPlainHostName(host)) return DIRECT;
    else return "PROXY 127.0.0.1:8080; PROXY 127.0.0.1:8081; DIRECT";
  }
`

func main() {
	pac, _ := gpac.New(scripts)

	r, _ := pac.FindProxyForURL("http://www.example.com/")
	fmt.Println(r) // returns PROXY 127.0.0.1:8080; PROXY 127.0.0.1:8081; DIRECT

	// 想要使用代理服务器，需要在本地启动代理服务器，监听8080端口，也就是proxy_server.go中的代码
	resp, _ := pac.Get("https://www.baidu.com/")
	fmt.Println(resp.Status)
}
