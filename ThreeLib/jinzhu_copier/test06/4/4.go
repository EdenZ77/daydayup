package main

import "fmt"

type Per struct {
	Name     string
	Age      int
	HouseIds [2]int // 数组，指定了长度
	CarIds   []int  // 切片，没指定长度
	Labels   map[string]string
}

/*
结构体默认是深拷贝，但如果结构体中包含map、slice等这些引用类型，默认也还是浅拷贝

map是引用类型，引用类型浅拷贝是默认的情况
*/
func main() {
	p1 := Per{
		Name:     "ssgeek",
		Age:      24,
		HouseIds: [2]int{22, 33},  // 数组，指定了长度，深拷贝
		CarIds:   []int{911, 718}, // 切片，引用类型，浅拷贝
		Labels:   map[string]string{"k1": "v1", "k2": "v2"},
		// 上述三个都是值类型，深拷贝，这个map是引用类型，浅拷贝
	}
	p2 := p1
	fmt.Printf("%v %p \n", p1, &p1) // {ssgeek 24 [22 33] [911 718] map[k1:v1 k2:v2]} 0xc000076180
	fmt.Printf("%v %p \n", p2, &p2) // {ssgeek 24 [22 33] [911 718] map[k1:v1 k2:v2]} 0xc0000761e0
	p2.Age = 19
	p2.Name = "likui"
	p2.HouseIds[1] = 44
	p2.CarIds[0] = 119              // 引用类型，同时改变
	p2.Labels["k1"] = "m1"          // 引用类型，同时改变
	fmt.Printf("%v %p \n", p1, &p1) // {ssgeek 24 [22 33] [119 718] map[k1:m1 k2:v2]} 0xc000076180
	fmt.Printf("%v %p \n", p2, &p2) // {likui 19 [22 44] [119 718] map[k1:m1 k2:v2]} 0xc0000761e0
}
