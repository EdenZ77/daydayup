package main

import (
	"errors"
	"fmt"
)

var ErrSentinel = errors.New("the underlying sentinel error")

func main() {

	//fmt.Printf("%T\n", ErrSentinel) // *errors.errorString

	err1 := fmt.Errorf("wrap sentinel:%w", ErrSentinel)
	err2 := fmt.Errorf("wrap err1:%w", err1)
	println(err2 == ErrSentinel)

	if errors.Is(err2, ErrSentinel) {
		println("err2 is ErrSentinel")
		fmt.Println(err2)
		fmt.Printf("格式化输出:%v", err2)
		return
	}

}
