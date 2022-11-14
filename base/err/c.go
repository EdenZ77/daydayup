package main

import (
	"os"

	"github.com/pkg/errors"
)

//var FSError = errors.New("自定义错误FS")

func C() error {
	_, err := os.Open("abc")
	if err != nil {
		//err = errors.Wrap(err, "我是 c中的Wrap ")
		//err = fmt.Errorf("使用fmt.Errorf封装错误，err:%w", FSError)
		//return err
		return errors.Errorf("cccccccc")
	}
	return nil
}
