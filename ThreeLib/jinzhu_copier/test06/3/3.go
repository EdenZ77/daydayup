package main

import "fmt"

type Per struct {
	Name     string
	Age      int
	HouseIds [2]int
}

/*
使用new函数实现值类型的浅拷贝

值类型的默认是深拷贝，想要实现值类型的浅拷贝，一般是两种方法:

	使用指针
	使用new函数（new函数返回的是指针）
*/
func main() {
	p1 := new(Per)
	p1.HouseIds = [2]int{22, 33}
	p1.Name = "songjiang"
	p1.Age = 20

	p2 := p1
	fmt.Printf("%v %p \n", p1, p1) // &{songjiang 20 [22 33]} 0xc000076180
	fmt.Printf("%v %p \n", p2, p2) // &{songjiang 20 [22 33]} 0xc000076180
	p2.Age = 19
	p2.Name = "likui"
	p2.HouseIds[1] = 44
	fmt.Printf("%v %p \n", p1, p1) // &{likui 19 [22 44]} 0xc000076180
	fmt.Printf("%v %p \n", p2, p2) // &{likui 19 [22 44]} 0xc000076180
}
