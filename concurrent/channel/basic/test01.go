package main

import "fmt"

func main() {
	ch := make(chan int)
	//close(ch)
	select {
	case name := <-ch:
		fmt.Println("case <- ch ", name)
	default:
		fmt.Println("取不出值")
	}
	//ch <- 10
	fmt.Println("发送成功")
}
