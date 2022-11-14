package main

import (
	"fmt"
	"reflect"
)

type withNoCopy struct {
	noCopy noCopy
}

type noCopy struct{}

func (*noCopy) Lock()   {}
func (*noCopy) Unlock() {}

func main() {
	t := withNoCopy{}
	dumpMethodSet(t)
	dumpMethodSet(&t)
}

func dumpMethodSet(i interface{}) {
	dynTyp := reflect.TypeOf(i)

	if dynTyp == nil {
		fmt.Printf("there is no dynamic type\n")
		return
	}

	n := dynTyp.NumMethod()
	if n == 0 {
		fmt.Printf("%s's method set is empty!\n", dynTyp)
		return
	}

	fmt.Printf("%s's method set:\n", dynTyp)
	for j := 0; j < n; j++ {
		fmt.Println("-", dynTyp.Method(j).Name)
	}
	fmt.Printf("\n")
}
