package p

import "fmt"

type V struct {
	i int32
	j int64
}

func (v V) PutI() {
	fmt.Printf("i=%d\n", v.i)
}

func (v V) PutJ() {
	fmt.Printf("j=%d\n", v.j)
}
