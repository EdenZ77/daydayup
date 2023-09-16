package main

import (
	"fmt"
	"reflect"
)

type order struct {
	ordId      int
	customerId int
}

func createQuery(q interface{}) {
	t := reflect.TypeOf(q)
	if t.Kind() != reflect.Struct {
		panic("unsupported argument type!")
	}

	v := reflect.ValueOf(q)
	for i := 0; i < t.NumField(); i++ {
		fmt.Println("FieldName:", t.Field(i).Name, "FiledType:", t.Field(i).Type,
			"FiledValue:", v.Field(i))
	}

}
func main() {
	o := order{
		ordId:      456,
		customerId: 56,
	}
	createQuery(o)

}
