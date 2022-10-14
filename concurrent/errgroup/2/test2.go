package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	fetchUrlDemo()
}

func fetchUrlDemo() {
	wg := sync.WaitGroup{}
	var urls = []string{
		"http://pkg.go.dev",
		"http://www.liwenzhou.com",
		"http://www.yixieqitawangzhi.com",
	}
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			resp, err := http.Get(url)
			if err == nil {
				fmt.Printf("获取%s成功\n", url)
				resp.Body.Close()
			}
			return
		}(url)
	}
	wg.Wait()
}
