package main

import (
	"fmt"
	"net/http"
)

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
		err := validateAuth(r.Header.Get("auth"))
		if err != nil {
			http.Error(w, "bad auth param", http.StatusUnauthorized)
			return
		}
		h.ServeHTTP(w, r)
	})
}

func main() {
	http.ListenAndServe(":8080", logHandler2(logHandler1(authHandler(http.HandlerFunc(greetings)))))
}
