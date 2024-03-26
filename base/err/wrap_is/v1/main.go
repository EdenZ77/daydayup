package main

import (
	"errors"
	"fmt"
)

// MyError 自定义错误类型
type MyError struct {
	Msg string
	Err error
}

func (e *MyError) Error() string {
	return e.Msg + ": " + e.Err.Error()
}

// Unwrap Is方法会调用Unwrap方法来获取错误链中的下一个错误
func (e *MyError) Unwrap() error {
	return e.Err
}

func main() {
	originalErr := errors.New("original error")
	wrappedErr := &MyError{"something failed", originalErr}

	// 检查错误是否包含原始错误
	if errors.Is(wrappedErr, originalErr) {
		fmt.Println("The wrapped error contains the original error")
	}
}
