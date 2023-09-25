package main

import "fmt"

func main() {
	//defer func() {
	//	if err := recover(); err != nil {
	//		fmt.Println("捕获异常:", err.Error()) // 编译错误：err.Error undefined (type interface {} is interface with no methods)
	//	}
	//}()
	//panic("a")

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("捕获异常:", fmt.Errorf("%v", err).Error())
		}
	}()
	panic("a")
}
