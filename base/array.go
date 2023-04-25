package main

import (
	"fmt"
	"strings"
)

func main() {
	//fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
	//fmt.Printf("%q\n", strings.Split(" xyz ", ""))
	//fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))

	fmt.Println(strings.Trim("  !!! Achtung! Achtung! !!!  ", " ")) // 去掉前后端两个空格
	fmt.Println(strings.Trim(" !!! Achtung! Achtung! !!! ", "!"))   // 原样返回
	// 首先匹配到空格，串首尾空格字符被删除，然后匹配到 “!”，继续删除首尾的各三个 “!”，于是得到该结果串。
	fmt.Println(strings.Trim(" !!! Achtung! Achtung! !!! ", "! "))
}
