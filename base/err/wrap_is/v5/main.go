package main

import (
	"errors"
	"fmt"
	"io/fs"
)

type MyFS struct {
	Msg string
}

func (m MyFS) Error() string {
	return m.Msg
}

// Is 方法用于判断错误链中是否包含特定错误
// 使用 Is 方法为自定义错误类型提供额外的比较逻辑，以确定一个错误是否应该被视为等同于另一个特定的错误类型或值。
/*
这是如何工作的：
1. 当你调用 errors.Is(err, target) 时，如果 err 本身实现了 Is(target error) bool 方法，errors.Is 会调用 err.Is(target)。
2. 如果 err.Is(target) 返回 true，它表示 err 想要告诉调用者它与 target 表示相同的错误情况，即使它们可能是不同的具体错误类型或实例。
3. 如果 err.Is(target) 返回 true，errors.Is(err, target) 也会返回 true，表示 err 与 target 匹配。
这种机制允许你自定义错误的等价性规则。例如，假设你有一个代表文件系统错误的自定义错误类型 MyFS，并且你希望建立 MyFS 与标准库中 fs.ErrNotExist 错误的等价性：
*/
func (m MyFS) Is(target error) bool {
	return target == fs.ErrNotExist
}

func someFunction() error {
	return MyFS{Msg: "file does not exist"}
}

func main() {
	err := someFunction()

	if errors.Is(err, fs.ErrNotExist) {
		fmt.Println("The file does not exist.")
	} else {
		fmt.Println("The error is not about a missing file.")
	}
}
