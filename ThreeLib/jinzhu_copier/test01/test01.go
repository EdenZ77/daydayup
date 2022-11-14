package main

import (
	"fmt"
	"github.com/jinzhu/copier"
)

type User struct {
	Name string
	Age  string
}

type Employee struct {
	Name string
	Age  int32
	Role string
}

func main() {
	user := User{"dj", "18"}
	employee := Employee{}
	// 很好理解，就是将user对象中的字段赋值到employee的同名字段中。如果目标对象中没有同名的字段，则该字段被忽略。
	// int8也能够转到赋值到int32中；但是如果源对象中age为string，目标对象age是int32，这个时候不会报错，目标对象被赋值为默认值
	copier.Copy(&employee, &user)
	fmt.Printf("%#v\n", employee)
}
