package main

import (
	"fmt"
	"reflect"
)

//type StructField struct {
//	// Name是字段的名字。PkgPath是非导出字段的包路径，对导出字段该字段为""。
//	// 参见http://golang.org/ref/spec#Uniqueness_of_identifiers
//	Name      string
//	PkgPath   string
//	Type      Type      // 字段的类型
//	Tag       StructTag // 字段的标签
//	Offset    uintptr   // 字段在结构体中的字节偏移量
//	Index     []int     // 用于Type.FieldByIndex时的索引切片
//	Anonymous bool      // 是否匿名字段
//}

type student struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

// Study 给student添加两个方法 Study和Sleep(注意首字母大写)
func (s student) Study() string {
	msg := "好好学习，天天向上。"
	fmt.Println(msg)
	return msg
}

func (s student) Sleep() string {
	msg := "好好睡觉，快快长大。"
	fmt.Println(msg)
	return msg
}

func printMethod(x any) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Println(t.NumMethod()) // 2
	for i := 0; i < t.NumMethod(); i++ {
		methodType := t.Method(i).Type
		fmt.Printf("method name:%s\n", t.Method(i).Name)
		fmt.Printf("method:%s\n", methodType)
		// 通过反射调用方法传递的参数必须是 []reflect.Value 类型
		var args []reflect.Value
		v.Method(i).Call(args)
	}
}

func main() {

	stu1 := student{"小王子", 90}
	t := reflect.TypeOf(stu1)
	fmt.Println(t.Name(), t.Kind()) // student struct
	// 通过for循环遍历结构体的所有字段信息
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
	}
	// 通过字段名获取指定结构体字段信息
	if scoreField, ok := t.FieldByName("Score"); ok {
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", scoreField.Name, scoreField.Index, scoreField.Type, scoreField.Tag.Get("json"))
	}

	printMethod(stu1)
}
