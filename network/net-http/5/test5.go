package main

import (
	"fmt"
	"net/http"
)

// 虽然默认的多路复用器很好用，但仍然不推荐使用，因为它是一个全局变量，所有的代码都可以修改它。
// 有些第三方库中可能与默认复用器产生冲突。所以推荐的做法是自定义。

func newservemux(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "NewServeMux")
}

func newservemuxhappy(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Newservemuxhappy")
}

func newservemuxbad(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "NewServeMuxbad")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", newservemux)
	mux.HandleFunc("/happy", newservemuxhappy)
	mux.HandleFunc("/bad", newservemuxbad)
	s := &http.Server{
		Addr:    ":8081",
		Handler: mux,
	}

	s.ListenAndServe()
}
