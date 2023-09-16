package main

import (
	"fmt"
	pool2 "hello/concurrent/channel/goroutinePool/test01/pool"
	"time"
)

func main() {
	p := pool2.New(5, pool2.WithPreAllocWorkers(false), pool2.WithBlock(false))

	time.Sleep(time.Second * 2)
	for i := 0; i < 10; i++ {
		err := p.Schedule(func() {
			time.Sleep(time.Second * 1)
			fmt.Printf("task end\n")
		})
		if err != nil {
			fmt.Printf("task[%d]: error: %s\n", i, err.Error())
		}
	}

	time.Sleep(10 * time.Second)
	p.Free()
}
