package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func A() error {
	fmt.Println("a --> b")
	err := B()
	if err != nil {
		return errors.WithStack(err)
	}
	return nil

}
