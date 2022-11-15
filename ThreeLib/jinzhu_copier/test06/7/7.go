package main

import (
	"fmt"
	"github.com/jinzhu/copier"
)

/*
方法三：使用其他三方库（这里还没深入）
*/

type Per struct {
	Name     string
	Age      int
	HouseIds [2]int // 数组，指定了长度
	CarIds   []int  // 切片，没指定长度
	Labels   map[string]string
}

func main() {
	p1 := Per{
		Name:     "ssgeek",
		Age:      24,
		HouseIds: [2]int{22, 33},  // 数组，指定了长度，深拷贝
		CarIds:   []int{911, 718}, // 切片，引用类型，浅拷贝
		Labels:   map[string]string{"k1": "v1", "k2": "v2"},
		// 前三个都是值类型，深拷贝，这个map是引用类型，浅拷贝
	}
	var p2 Per
	err := copier.CopyWithOption(&p2, &p1, copier.Option{DeepCopy: true})
	if err != nil {
		return
	}
	fmt.Printf("%v %p \n", p1, &p1) // {ssgeek 24 [22 33] [911 718] map[k1:v1 k2:v2]} 0xc00006c050
	fmt.Printf("%v %p \n", p2, &p2) // {ssgeek 24 [22 33] [911 718] map[k1:v1 k2:v2]} 0xc00006c140
	p1.Age = 19
	p1.Name = "likui"
	p1.HouseIds[1] = 44
	p1.CarIds[0] = 119
	p1.Labels["k1"] = "m1"
	fmt.Printf("%v %p \n", p1, &p1) // {likui 19 [22 44] [119 718] map[k1:m1 k2:v2]} 0xc00006c050
	fmt.Printf("%v %p \n", p2, &p2) // {ssgeek 24 [22 33] [911 718] map[k1:v1 k2:v2]} 0xc00006c140
}
