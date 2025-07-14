package main

import (
	"errors"
	"fmt"
)

type MyError struct {
	Code    int
	Message string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("error %d: %s", e.Code, e.Message)
}

func main() {
	err := &MyError{Code: 404, Message: "Not Found"}

	var target *MyError
	if errors.As(err, &target) {
		fmt.Printf("匹配成功: Code=%d, Message=%s\n", target.Code, target.Message)
	} else {
		fmt.Println("未匹配到MyError类型")
	}
}
