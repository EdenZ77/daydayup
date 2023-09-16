package p

import (
	"fmt"
	"unsafe"
)

// W 1+3+4+8
type W struct {
	b byte
	i int32
	j int64
}

func init() {
	var w *W = new(W)
	fmt.Printf("size=%d\n", unsafe.Sizeof(*w))
}
