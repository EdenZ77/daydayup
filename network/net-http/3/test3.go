package main

import (
	"fmt"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, "hello go 大师！")
	if err != nil {
		return
	}
}

func main() {

	//http.HandleFunc("/", sayHello)
	//err := http.ListenAndServe(":9090", nil)
	//if err != nil {
	//	fmt.Printf("http server failed, err:%v\n", err)
	//	return
	//}

	//http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", http.HandlerFunc(sayHello))
	if err != nil {
		return
	}

}
