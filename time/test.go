package main

import (
	"fmt"
	"io"
	"os"
)

var (
	myFprintf = func(w io.Writer, format string, a ...interface{}) (int, error) {
		return fmt.Fprintf(w, format, a...)
	}
)

func main() {
	fmt.Printf("%T\n", myFprintf)             // func(io.Writer, string, ...interface {}) (int, error)
	myFprintf(os.Stdout, "%s\n", "Hello, Go") // 输出Hello，Go
}
