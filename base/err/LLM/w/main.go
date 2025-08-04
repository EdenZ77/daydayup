package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

// %w格式动词的核心作用是：创建包含原始错误的新错误对象，同时保留完整的错误链。

func main() {
	_, err := os.Open("missing.txt")
	// 原始错误：*fs.PathError

	wrappedErr := fmt.Errorf("doSomething failed: %w", err)
	// 即使经过多层包装，仍能识别原始错误
	if errors.Is(wrappedErr, os.ErrNotExist) {
		fmt.Println("原始错误是文件不存在")
	}

	// 还能提取错误类型
	var pathErr *fs.PathError
	if errors.As(wrappedErr, &pathErr) {
		fmt.Println("Path:", pathErr.Path)
	}
}
