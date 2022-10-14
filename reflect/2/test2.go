package main

import (
	"fmt"
	"reflect"
)

type myInt int64

func reflectType(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Printf("type:%v kind:%v\n", t.Name(), t.Kind())
}

func main() {
	var a *float32 // 指针
	var b myInt    // 自定义类型
	var c rune     // 类型别名

	reflectType(a)
	reflectType(b)
	reflectType(c)

	type person struct {
		name string
		age  int
	}
	type book struct {
		title string
	}
	d := person{"eden", 18}
	e := &book{"<<go master>>"}
	reflectType(d)
	reflectType(e)
}
