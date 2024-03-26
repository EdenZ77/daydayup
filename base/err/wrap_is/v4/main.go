package main

import (
	"errors"
	"fmt"
)

// WrappedError 直接嵌套错误和上下文
type WrappedError struct {
	context string
	cause   error
}

func (w *WrappedError) Error() string {
	return w.context + ": " + w.cause.Error()
}

func (w *WrappedError) Unwrap() error {
	return w.cause
}

func main() {
	originalErr := errors.New("original error")
	wrappedErr := &WrappedError{"failed to execute", originalErr}

	// 检查错误是否包含原始错误
	if errors.Is(wrappedErr, originalErr) {
		fmt.Println("The wrapped error contains the original error")
	}
}
