package main

import (
	"fmt"
	"github.com/jinzhu/copier"
)

/*
调用目标方法
有时候源对象中的某个字段没有出现在目标对象中，但是目标对象有一个同名的方法，方法接受一个同类型的参数，
这时Copy会以源对象的这个字段作为参数调用目标对象的该方法

我们给Employee添加了一个Role方法，User的字段Role没有出现在Employee中，但是Employee有一个同名方法。
Copy函数内部会以user对象的Role字段为参数调用employee的Role方法。最终，我们的employee对象的SuperRole值变为SuperAdmin。实际上，这个方法中可以执行任何操作，不一定是赋值。
*/

type User struct {
	Name string
	Age  int
	Role string
}

type Employee struct {
	Name      string
	Age       int
	SuperRole string
}

// Role 注意参数的类型与源对象的Role属性类型一样
func (e *Employee) Role(xx string) {
	e.SuperRole = "Super" + xx
}

// 如果类型不一样，就赋值不成功，最后就是默认初始值
//func (e *Employee) Role(xx int) {
//	e.SuperRole = "Super" + fmt.Sprintf("%d", xx)
//}

func main() {
	user := User{Name: "dj", Age: 18, Role: "Admin"}
	employee := Employee{}

	copier.Copy(&employee, &user)
	fmt.Printf("%#v\n", employee)
}
