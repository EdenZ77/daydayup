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
结构体中含有引用类型的字段，那么这个字段就是浅拷贝，但是往往希望的是深拷贝，解决方案如下

方法一：挨个把可导致浅拷贝的引用类型字段自行赋值
赋值后，修改值就相互不影响了
*/
func main() {
	p1 := Per{
		Name:     "ssgeek",
		Age:      24,
		HouseIds: [2]int{22, 33},  // 数组，指定了长度，深拷贝
		CarIds:   []int{911, 718}, // 切片，引用类型，浅拷贝
		Labels:   map[string]string{"k1": "v1", "k2": "v2"},
		// 前三个都是值类型，深拷贝，这个map是引用类型，浅拷贝
	}
	p2 := p1
	// 切片赋值到新的切片
	//tmpCarIds := make([]int, 0)
	//for _, c := range p1.CarIds {
	//	tmpCarIds = append(tmpCarIds, c)
	//}
	tmpCarIds := make([]int, len(p1.CarIds))
	copy(tmpCarIds, p1.CarIds)
	// map赋值到新的map
	tmpLabels := make(map[string]string)
	for k, v := range p1.Labels {
		tmpLabels[k] = v
	}
	p2.CarIds = tmpCarIds
	p2.Labels = tmpLabels
	fmt.Printf("%v %p \n", p1, &p1) // {ssgeek 24 [22 33] [911 718] map[k1:v1 k2:v2]} 0xc00006c050
	fmt.Printf("%v %p \n", p2, &p2) // {ssgeek 24 [22 33] [911 718] map[k1:v1 k2:v2]} 0xc00006c0a0
	p1.Age = 19
	p1.Name = "likui"
	p1.HouseIds[1] = 44
	p1.CarIds[0] = 119
	p1.Labels["k1"] = "m1"
	fmt.Printf("%v %p \n", p1, &p1) // {likui 19 [22 44] [119 718] map[k1:m1 k2:v2]} 0xc00006c050
	fmt.Printf("%v %p \n", p2, &p2) // {ssgeek 24 [22 33] [911 718] map[k1:v1 k2:v2]} 0xc00006c0a0
}
