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

func greetings(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
}

func logHandler1(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//t := time.Now()
		//log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t)
		fmt.Println("log_1 =========before")
		h.ServeHTTP(w, r)
		fmt.Println("log_1 =========after")
	})
}

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
		h.ServeHTTP(w, r)
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
