package main

import (
	"fmt"
	"reflect"
)

type order struct {
	ordId      int
	customerId int
}

func createQuery(q any) {
	t := reflect.TypeOf(q)
	fmt.Println("Type ", t)
	if t.Kind() != reflect.Struct {
		panic("unsupported argument type!")
	}

	v := reflect.ValueOf(q)
	fmt.Println("Value ", v)
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
