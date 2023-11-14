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

	// Get issues request via a list of proxies and returns at the first request that succeeds
	resp, _ := pac.Get("http://www.example.com/")
	fmt.Println(resp.Status)
}
