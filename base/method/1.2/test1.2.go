package main

import (
	"fmt"
	"time"
)

type field struct {
	name string
}

func (p *field) print() {
	fmt.Println(p.name)
}

func main() {
	data1 := []*field{
		{
			name: "one",
		},
		{
			name: "two",
		},
		{
			name: "three",
		},
	}
	for _, v := range data1 {
		//go v.print()
		go (*field).print(v)
	}

	data2 := []field{{"four"}, {"five"}, {"six"}}
	for _, v := range data2 {
		//go v.print()
		go (*field).print(&v)
		//
		time.Sleep(time.Second)
	}

	time.Sleep(5 * time.Second)
}
