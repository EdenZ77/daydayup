package main

import (
	"fmt"
	"reflect"
)

func main() {
	m1 := map[int]interface{}{1: []int{1, 2, 3}, 2: 3, 3: "a"}
	m2 := map[int]interface{}{1: []int{1, 2, 3}, 2: 3, 3: "a"}
	if reflect.DeepEqual(m1, m2) {
		fmt.Println("相等")
	}
}
