package main

import (
	"errors"
	"fmt"
)

func main() {
	// 创建一个被包装的错误
	err := errors.New("original error")
	wrappedErr := fmt.Errorf("an error occurred: %w", err)

	// 解包错误
	unwrappedErr := errors.Unwrap(wrappedErr)

	// 比较解包后的错误和原始错误
	fmt.Printf("Unwrapped error: %v\n", unwrappedErr)
	fmt.Printf("Is unwrapped error equal to the original error? %v\n", unwrappedErr == err)
}
