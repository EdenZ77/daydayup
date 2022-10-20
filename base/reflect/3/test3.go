package main

import (
	"fmt"
	"reflect"
)

func reflectValue(x any) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		fmt.Printf("type is int64, value is %d\n", v.Int())
	case reflect.Float32:
		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
	case reflect.Float64:
		fmt.Printf("type is float64, value is %f\n", v.Float())
	}
}

func main() {
	var a float32 = 3.14
	var b int64 = 100

	reflectValue(a) // type is float32, value is 3.140000
	reflectValue(b) // type is int64, value is 100

	c := reflect.ValueOf(10)
	fmt.Printf("type c :%T\n", c) // type c :reflect.Value
}
