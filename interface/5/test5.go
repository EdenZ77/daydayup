package main

import (
	"fmt"
	"net/http"
)

func greetings(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
}

type HandlerAdd interface {
	Add(int, int)
}

type HandlerFunc func(int, int)

func (receiver HandlerFunc) name() {

}

func testAdd(a int, b int) {
	fmt.Printf(" a + b = %d\n", a+b)
}

func main() {
	//http.ListenAndServe(":8080", http.HandlerFunc(greetings))

	//HandlerFunc(testAdd(1, 2))

}
