package main

import (
	"fmt"
	"github.com/jinzhu/copier"
)

/*
方法赋值：
我们给User添加一个DoubleAge方法。Employee结构有字段DoubleAge，User中没有，但是User有一个同名的方法，
这时Copy调用user的DoubleAge方法为employee的DoubleAge赋值，得到 36。
*/

type User struct {
	Name string
	Age  int
}

func (u *User) DoubleAge() int {
	return 2 * u.Age
}

type Employee struct {
	Name      string
	DoubleAge int
	Role      string
}

func main() {
	user := User{"dj", 18}
	employee := Employee{}

	copier.Copy(&employee, &user)
	fmt.Printf("%#v\n", employee)
}
