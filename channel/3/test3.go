package main

import (
	"errors"
	"fmt"
	"time"
)

// 无缓冲 channel用法：用作同步传递goroutine之间的信息

// 这个示例在 main goroutine 与子 goroutine 之间建立了一个元素类型为 error 的 channel，
// 子 goroutine 退出时，会将它执行的函数的错误返回值写入这个 channel，
// main goroutine 可以通过读取 channel 的值来获取子 goroutine 的退出状态。

func spawn(f func() error) <-chan error {
	c := make(chan error)

	go func() {
		c <- f()
	}()

	return c
}

func main() {
	c := spawn(func() error {
		time.Sleep(2 * time.Second)
		return errors.New("timeout")
		//return nil
	})
	fmt.Println(<-c)
}
