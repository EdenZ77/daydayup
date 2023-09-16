package main

import "fmt"

type MySlice[T int | float32] []T

func (s MySlice[T]) Sum() T {
	var sum T
	for _, value := range s {
		sum += value
	}
	return sum
}

func main() {
	var s MySlice[int] = []int{1, 2, 3, 4}
	fmt.Println(s.Sum())
}
