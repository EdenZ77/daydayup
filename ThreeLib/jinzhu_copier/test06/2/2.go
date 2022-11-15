package main

import "fmt"

type Per struct {
	Name     string
	Age      int
	HouseIds [2]int
}

/*
结构体的浅拷贝

使用指针进行浅拷贝，浅拷贝中，可以看到p1和p2的内存地址是相同的，修改其中一个对象的属性时，另一个也会产生变化
*/
func main() {
	p1 := Per{
		Name:     "ssgeek",
		Age:      24,
		HouseIds: [2]int{22, 33},
	}
	p2 := &p1
	fmt.Printf("%v %p \n", p1, &p1) // {ssgeek 24 [22 33]} 0xc000076180
	fmt.Printf("%v %p \n", p2, p2)  // &{ssgeek 24 [22 33]} 0xc000076180
	p2.Age = 19
	p2.Name = "likui"
	p2.HouseIds[1] = 44
	fmt.Printf("%v %p \n", p1, &p1) // {likui 19 [22 44]} 0xc000076180
	fmt.Printf("%v %p \n", p2, p2)  // &{likui 19 [22 44]} 0xc000076180
}
