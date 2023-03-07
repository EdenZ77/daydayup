package main

import "fmt"

func main() {

	s := []int{1, 1, 1}
	f(s)
	fmt.Println(s)
}

func f(s []int) {
	for i := range s {
		s[i]++
	}
	s = append(s, 7)
}
