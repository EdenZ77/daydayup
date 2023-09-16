package main

import (
	"fmt"
	"reflect"
)

/*
这个文件夹的内容来自：https://mp.weixin.qq.com/s?__biz=MzU5NTAzNjc3Mg==&mid=2247483962&idx=1&sn=e13df5c5e016215302205f5ec8fbb857&scene=21#wechat_redirect


*/

func main() {
	//author := "draven"
	//fmt.Println("TypeOf author:", reflect.TypeOf(author))
	//fmt.Println("ValueOf author:", reflect.ValueOf(author))

	//v := reflect.ValueOf(1)
	//i := v.Interface().(int)
	//fmt.Println(i)

	i := 1
	v := reflect.ValueOf(&i)
	// Elem() Elem返回接口v所包含的值或指针v所指向的值。
	v.Elem().SetInt(10)
	fmt.Println(i)
}
