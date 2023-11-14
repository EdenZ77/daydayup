package main

import (
	"fmt"
	"github.com/jinzhu/copier"
)

/*
将结构体赋值到切片:
这个不难，实际上就是根据源对象生成一个和目标切片类型相符合的对象，然后append到目标切片中
*/

type User struct {
	Name string
	Age  int
}

type Employee struct {
	Name string
	Age  int
	Role string
}

/*
下面代码中，Copy先通过user生成一个Employee对象，然后append到切片employees中
*/
func main() {
	user := User{Name: "dj", Age: 18}
	var employees []Employee

	copier.Copy(&employees, &user)
	fmt.Printf("%#v\n", employees)
}
