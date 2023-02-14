package main

import (
	"fmt"
	"net/http"
)

type greeting string

func (g greeting) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, g)
	if err != nil {
		return
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	// Fprintln的第一个参数是接口io.Writer，而w参数的类型是http.ResponseWriter,但是这两个接口都有相同的Write方法，所以w参数能传递过来
	// 变量能否赋值给接口类型就看变量的方法集是否包含接口所有方法
	_, err := fmt.Fprintln(w, "Hello World")
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/greeting", greeting("Welcome, dj"))
	http.ListenAndServe(":8080", nil)
}
