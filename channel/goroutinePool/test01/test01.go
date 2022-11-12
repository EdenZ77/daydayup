package main

import (
	"fmt"
	"hello/channel/goroutinePool/test01/pool"
	"time"
)

func main() {
	p := pool.New(5, pool.WithPreAllocWorkers(false), pool.WithBlock(false))

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
