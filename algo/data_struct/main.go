package main

import "fmt"

func main() {
	//a := make([]int, 2, 10)
	//fmt.Println(a)      // [0 0]
	//fmt.Println(len(a)) // 2
	//fmt.Println(cap(a)) // 10

	//a := make([]int, 0, 10)
	//fmt.Println(a)      // []
	//fmt.Println(len(a)) // 0
	//fmt.Println(cap(a)) // 10

	a := make([]int, 10)
	fmt.Println(a)      // [0 0 0 0 0 0 0 0 0 0]
	fmt.Println(len(a)) // 10
	fmt.Println(cap(a)) // 10
}
