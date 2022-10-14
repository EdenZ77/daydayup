package main

import (
	"errors"
	"fmt"
)

type MyError struct {
	error
}

var ErrBad = MyError{
	errors.New("bad things happened"),
}

func bad() bool {
	return false
}
func returnsError() error {
	var p *MyError = nil
	if bad() {
		p = &ErrBad
	}
	return p
}

func main() {
	err := returnsError()
	if err != nil {
		fmt.Printf("error occor: %+v\n", err)
	}
	fmt.Println("ok")
}
