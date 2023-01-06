package main

import "fmt"

func main() {
	//fmt.Println(5 / 2)
	arr := []int{5, 3, 6, 2, 7, 1, 9}
	fmt.Println(cap(arr[0:2]))
	fmt.Println(len(arr[0:2]))
	fmt.Println(cap(arr[2:]))
	fmt.Println(len(arr[2:]))
}

//var result = make([]int, 50)

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n-1)%1000000007 + fib(n-2)%1000000007
}
