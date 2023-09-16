package main

import (
	"fmt"
	"reflect"
)

/*
现在让我们理解一下 ”在运行时知道变量的类型的必要“。假设我们要编写一个简单的函数，它将一个结构体作为参数，并使用这个参数创建一个SQL插入语句。

*/

type order struct {
	ordId      int
	customerId int
}

type employee struct {
	name    string
	id      int
	address string
	salary  int
	country string
}

func createQuery(q interface{}) {
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		t := reflect.TypeOf(q).Name()
		query := fmt.Sprintf("insert into %s values(", t)
		v := reflect.ValueOf(q)
		for i := 0; i < v.NumField(); i++ {
			// 注意reflect.Value 也实现了NumField,Kind这些方法
			// 这里的v.Field(i).Kind()等价于t.Field(i).Type.Kind()
			switch v.Field(i).Kind() {
			case reflect.Int:
				if i == 0 {
					query = fmt.Sprintf("%s%d", query, v.Field(i).Int())
				} else {
					query = fmt.Sprintf("%s, %d", query, v.Field(i).Int())
				}
			case reflect.String:
				if i == 0 {
					query = fmt.Sprintf("%s%s", query, v.Field(i).String())
				} else {
					query = fmt.Sprintf("%s, %s", query, v.Field(i).String())
				}
			default:
				fmt.Println("Unsupported type")
				return
			}
		}
		query = fmt.Sprintf("%s)", query)
		fmt.Println(query)
	} else {
		fmt.Println("Unsupported type==")
	}
}

func main() {
	o := order{
		ordId:      456,
		customerId: 56,
	}
	createQuery(o)

	e := employee{
		name:    "Naveen",
		id:      565,
		address: "Coimbatore",
		salary:  90000,
		country: "India",
	}
	createQuery(e)
	i := 90
	createQuery(i)

}
