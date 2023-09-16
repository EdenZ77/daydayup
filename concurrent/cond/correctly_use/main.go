package main

import (
	"sync"
	"time"
)

func main() {

	m := sync.Mutex{}
	c := sync.NewCond(&m)

	go func() {
		time.Sleep(1 * time.Second)
		c.Broadcast()
	}()

	m.Lock()
	time.Sleep(2 * time.Second)
	// Broadcast 先执行完毕，main goroutine将无法被唤醒
	c.Wait()

}
