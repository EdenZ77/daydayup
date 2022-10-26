package main

import (
	"fmt"
	"sync"
	"time"
)

var threeOnce struct {
	sync.Once
	v float32
}

func three() float32 {
	threeOnce.Do(func() {
		time.Sleep(5 * time.Second)
		threeOnce.v = float32(3.0)
	})
	return threeOnce.v
}

func main() {
	// 等到func执行完成后才打印
	fmt.Println(three())
}
