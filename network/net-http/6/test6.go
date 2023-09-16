package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

/**
ServeMux的一个缺陷是无法使用变量实现URL模式匹配。
而HttpRouter可以，HttpRouter是一个高性能的第三方HTTP路由包，弥补了net/http包中的路由不足问题。

如何使用？
 go get  github.com/julienschmidt/httprouter

httprouter的使用首先得使用New()函数，生成一个*Router路由对象，然后使用GET()，方法去注册匹配的函数，
最后再将这个参数传入http.ListenAndServe函数就可以监听服务。
*/

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)

	log.Fatal(http.ListenAndServe(":8080", router))
}
