package main

import (
	"fmt"
	"hello/package1"
	package1_sub "hello/package1/sub_package1"
	package1_sub2 "hello/package1/sub_package2"
)

func main() {
	var name package1.Person
	name.Name = "44"
	fmt.Println(package1.Name)
	fmt.Println(package1_sub2.Sub2Name)
	fmt.Println(package1_sub.SubName)

}
