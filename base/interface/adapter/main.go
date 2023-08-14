package main

import (
	"fmt"
	"net/http"
)

func greetings(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
}

func main() {
	http.ListenAndServe(":8080", http.HandlerFunc(greetings))
}

/*
我们可以看到，这个例子通过 http.HandlerFunc 这个适配器函数类型，将普通函数 greetings 快速转化为满足 http.Handler 接口的类型。而 http.HandleFunc 这个适配器函数类型的定义是这样的：

// $GOROOT/src/net/http/server.go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}

type HandlerFunc func(ResponseWriter, *Request)

func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
    f(w, r)
}
经过 HandlerFunc 的适配转化后，我们就可以将它的实例用作实参，传递给接收 http.Handler 接口的 http.ListenAndServe 函数，从而实现基于接口的组合。
*/
