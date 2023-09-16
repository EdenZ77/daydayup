package main

import (
	"fmt"
	"time"
)

/*
格式化时间涉及到两个转换函数

func Parse(layout, value string) (Time, error) {}
	Parse 函数用于将时间字符串根据它所能对应的布局转换为 time.Time 对象。
func (t Time) Format(layout string) string {}
	Formate 函数用于将 time.Time 对象根据给定的布局转换为时间字符串。
*/

const (
	layoutISO = "2006-01-02"
	layoutUS  = "January 2, 2006"
)

func main() {
	date := "2012-08-09"
	t, _ := time.Parse(layoutISO, date)
	fmt.Println(t)                  // 2012-08-09 00:00:00 +0000 UTC
	fmt.Println(t.Format(layoutUS)) // August 9, 2012
}
