package main

import (
	"fmt"
	"sort"
)

func main() {
	slice1 := []string{"apple", "banana", "cherry"}
	slice2 := []string{"cherry", "banana", "apple"}

	equal := SlicesEqual(slice1, slice2)
	fmt.Println("Are the slices equal?", equal)
}

func SlicesEqual(slice1, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	// 对切片进行排序
	sort.Strings(slice1)
	sort.Strings(slice2)

	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}

	return true
}
