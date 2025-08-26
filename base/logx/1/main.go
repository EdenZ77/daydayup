package main

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
)

func main() {
	flag.Parse()
	defer glog.Flush()
	if err := loadConfig(); err != nil {
		glog.Error(err)
	}
}

// 正例：直接返回错误
func loadConfig() error {
	return decodeConfig() // 直接返回
}

// 正例：如果需要基于函数返回的错误，封装更多的信息，可以封装返回的 err。否则，建议直接返回 err
func decodeConfig() error {
	if err := readConfig(); err != nil {
		// 添加必要的信息，用户名称
		return fmt.Errorf("could not decode configuration data for user %s: %v", "colin", err)
	}
	return nil
}

func readConfig() error {
	glog.Errorf("read: end of input.")
	return fmt.Errorf("read: end of input")
}
