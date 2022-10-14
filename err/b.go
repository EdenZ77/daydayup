package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func B() error {
	fmt.Println("b --> c")
	err := C()
	if err != nil {
		return errors.WithMessage(err, "我是b中的withMessage")
		//return err
	}
	return nil
}
