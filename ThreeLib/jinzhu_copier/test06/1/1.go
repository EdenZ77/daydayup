package main

import "fmt"

type Per struct {
	Name     string
	Age      int
	HouseIds [2]int
}

/*
展示struct的深拷贝

结构体类型中的字段是值类型，拷贝时都是深拷贝
*/
func main() {
	p1 := Per{
		Name:     "eden",
		Age:      24,
		HouseIds: [2]int{22, 33},
	}
	p2 := p1
	fmt.Printf("%v %p \n", p1, &p1) // {ssgeek 24 [22 33]} 0xc000180030
	fmt.Printf("%v %p \n", p2, &p2) // {ssgeek 24 [22 33]} 0xc000180060
	p2.Age = 19
	p2.Name = "likui"
	p2.HouseIds[1] = 44
	fmt.Printf("%v %p \n", p1, &p1) // {ssgeek 24 [22 33]} 0xc000098180
	fmt.Printf("%v %p \n", p2, &p2) // {likui 19 [22 44]} 0xc0000981b0
}
