package main

import (
	"fmt"
	"net/http"
)

/*
在 Go Web 编程中，“中间件”常常指的是一个实现了 http.Handler 接口的 http.HandlerFunc 类型实例。实质上，这里的中间件就是包装模式和适配器模式结合的产物。
*/

func validateAuth(s string) error {
	if s != "123456" {
		return fmt.Errorf("%s", "bad auth token")
	}
	return nil
}

// 相比之前的包装模式，这里仅有一个函数，之前的包装模式拥有自己的类型和相应的属性；对于只有一个函数，则需要通过显示的类型转换，
// 将函数转换为 http.HandlerFunc 类型，从而满足 http.Handler 接口。
func greetings(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
}

// logHandler2执行过程中调用logHandler1的返回值，这个返回值被显示转换满足了接口实现，执行过程中又调用了logHandler1传递的参数
// 同理了。
func logHandler1(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//t := time.Now()
		//log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t)
		fmt.Println("log_1 =========before")
		h.ServeHTTP(w, r)
		fmt.Println("log_1 =========after")
	})
}

// 这个放到最外层，那么http.ListenAndServe执行接口的实现时，就会调用logHandler2的返回值，这个返回值被显示转换，满足了接口要求
// 执行时，又去调用logHandler2传递的参数接口的实现，这个实现就是logHandler1的返回值。
func logHandler2(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//t := time.Now()
		//log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t)
		fmt.Println("log_2 =========before")
		h.ServeHTTP(w, r)
		fmt.Println("log_2 =========after")
	})
}

func authHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("authHandler before=======")
		err := validateAuth(r.Header.Get("auth"))
		if err != nil {
			http.Error(w, "bad auth param", http.StatusUnauthorized)
			return
		}
		h.ServeHTTP(w, r) // 最终这里就是greetings函数
		fmt.Println("authHandler after=========")
	})
}

/*
log_2 =========before
log_1 =========before
authHandler before=======
authHandler after=========
log_1 =========after
log_2 =========after
*/
func main() {
	http.ListenAndServe(":8080", logHandler2(logHandler1(authHandler(http.HandlerFunc(greetings)))))
}

// http.ListenAndServe中参数是Handler接口，那ListenAndServe里面肯定会调用接口中的ServeHTTP函数
// 我们需要传递的是接口的实现，

//type Handler interface {
//	ServeHTTP(ResponseWriter, *Request)
//}
//
//type HandlerFunc func(ResponseWriter, *Request)
//
//func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
//	f(w, r)
//}
