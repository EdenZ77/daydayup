package main

import (
	"fmt"
	"strings"
)

func main() {
	// 会将cutset这个字符串拆分成一个个的字符，然后去s字符串的首尾匹配，只要匹配任意一个字符就去掉，直到遇到没有匹配的，剩下的便是返回的结果
	fmt.Printf("%#v\n", strings.Trim("!!! Achtung! Achtung! !!!", "! "))       // "Achtung! Achtung"
	fmt.Printf("%#v\n", strings.Trim(" !!! Achtung! Achtung! !!! ", "! !"))    // "Achtung! Achtung"
	fmt.Printf("%#v\n", strings.Trim("Hello world hello world", "world"))      // "Hello world hello "
	fmt.Printf("%#v\n", strings.Trim("Hello world hello xwwwwworld", "world")) // "Hello world hello x"
	fmt.Println("========TrimLeft  TrimRight=======")
	fmt.Printf("%#v\n", strings.TrimLeft("aabbccdd", "abcd"))   // "" 空字符串，因为s字符串全部被cutset的单个字符匹配了，所以返回值就是空字符串
	fmt.Printf("%#v\n", strings.TrimLeft("aabbccdde", "abcd"))  // "e"
	fmt.Printf("%#v\n", strings.TrimLeft("aabbedcba", "abcd"))  // "edcba"
	fmt.Printf("%#v\n", strings.TrimRight("aabbccdd", "abcd"))  // "" 空字符串
	fmt.Printf("%#v\n", strings.TrimRight("aabbccdde", "abcd")) // "aabbccdde"
	fmt.Printf("%#v\n", strings.TrimRight("aabbedcba", "abcd")) // "aabbe"

}
