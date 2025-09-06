package main

import (
	"fmt"
	"reflect"
)

func main() {
	methodCallTest()
}

func baseTypeTest() {
	// 整型
	i := 42
	v := reflect.ValueOf(&i).Elem()    // 必须获取可设置的值
	fmt.Println("CanSet:", v.CanSet()) // true
	v.SetInt(100)
	fmt.Println(i) // 100

	// 整型
	a := 42
	va := reflect.ValueOf(a)
	fmt.Println("CanSet:", va.CanSet()) // false
	// panic: reflect: reflect.Value.SetInt using unaddressable value
	//va.SetInt(100)

	// 字符串
	s := "hello"
	vs := reflect.ValueOf(&s).Elem()
	vs.SetString("world")
	fmt.Println(s) // world
}

func settabilityTest() {
	s := []int{1, 2, 3}
	v := reflect.ValueOf(s)
	v.Index(0).SetInt(10)    // 直接修改元素，不需要指针
	fmt.Printf("s: %v\n", s) // s: [10 2 3]

	m := map[string]int{"a": 1}
	vm := reflect.ValueOf(m)
	vm.SetMapIndex(reflect.ValueOf("a"), reflect.ValueOf(2))
	fmt.Printf("m: %v\n", m) // m: map[a:2]

	var i interface{} = 42
	iv := reflect.ValueOf(&i).Elem()
	iv.Set(reflect.ValueOf(100)) // 修改接口持有的值
	fmt.Printf("i: %v\n", i)     // i: 100

	type S struct{ n int } // 小写 n 未导出
	s1 := &S{n: 1}
	sv := reflect.ValueOf(s1).Elem()
	field := sv.Field(0)
	fmt.Println(field.CanSet()) // false
}

func modifyValueTest() {
	var i int = 42
	modifyValue(&i)
	fmt.Println(i) // 100

	var s string = "old"
	modifyValue(&s)
	fmt.Println(s) // "new"
}

func modifyValue(ptr any) {
	v := reflect.ValueOf(ptr).Elem()
	if v.CanSet() {
		// 根据类型设置值
		switch v.Kind() {
		case reflect.Int:
			v.SetInt(100)
		case reflect.String:
			v.SetString("new")
			// 其他类型处理...
		}
	}
}

type Person struct {
	Name string
	Age  int
}

func structTest() {
	p := Person{"Alice", 30}
	v := reflect.ValueOf(&p).Elem()

	// 获取字段值
	nameField := v.FieldByName("Name")
	fmt.Println(nameField.String()) // Alice

	// 设置字段值
	ageField := v.FieldByName("Age")
	if ageField.CanSet() {
		ageField.SetInt(31)
	}
	fmt.Println(p.Age) // 31
}

func sliceTest() {
	slice := []int{1, 2, 3}
	v := reflect.ValueOf(&slice).Elem()

	// 追加元素
	newSlice := reflect.Append(v, reflect.ValueOf(4))
	// 如果使用 v := reflect.ValueOf(slice) 则将导致下面的panic
	// panic: reflect: reflect.Value.Set using unaddressable value
	v.Set(newSlice)
	fmt.Println(slice) // [1 2 3 4]

	// 修改元素
	if v.Len() > 0 {
		v.Index(0).SetInt(10)
	}
	fmt.Println(slice) // [10 2 3 4]
}

func mapTest() {
	m := map[string]int{"a": 1, "b": 2}
	v := reflect.ValueOf(&m).Elem()

	// 添加新键值对
	key := reflect.ValueOf("c")
	value := reflect.ValueOf(3)
	v.SetMapIndex(key, value)
	fmt.Println(m) // map[a:1 b:2 c:3]

	// 删除键
	v.SetMapIndex(key, reflect.Value{})
	fmt.Println(m) // map[a:1 b:2]
}

func Add(a, b int) int {
	return a + b
}

func funcCallTest() {
	// 获取函数值
	funcValue := reflect.ValueOf(Add)

	// 准备参数
	args := []reflect.Value{
		reflect.ValueOf(3),
		reflect.ValueOf(4),
	}

	// 调用函数
	results := funcValue.Call(args)
	fmt.Println(results[0].Int()) // 7
}

type Calculator struct{}

func (c Calculator) Multiply(x, y int) int {
	return x * y
}

func methodCallTest() {
	calc := Calculator{}
	v := reflect.ValueOf(calc)

	// 获取方法
	method := v.MethodByName("Multiply")

	// 调用方法
	result := method.Call([]reflect.Value{
		reflect.ValueOf(5),
		reflect.ValueOf(6),
	})
	fmt.Println(result[0].Int()) // 30
}
