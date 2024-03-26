package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func main() {
	originalErr := errors.New("original error")
	wrappedErr := errors.Wrap(originalErr, "additional context")

	// 检查错误是否包含原始错误
	if errors.Is(wrappedErr, originalErr) {
		fmt.Println("The wrapped error contains the original error")
	}
}
