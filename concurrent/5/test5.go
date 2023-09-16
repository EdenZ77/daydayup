package main

import (
	"fmt"
	"sync"
	"time"
)

var active = make(chan struct{}, 3)
var jobs = make(chan int, 10)

func main() {
	go func() {
		for i := 0; i < 8; i++ {
			jobs <- i + 1
		}
		close(jobs)
	}()

	var wg sync.WaitGroup
	for job := range jobs {
		go func(j int) {
			wg.Add(1)
			active <- struct{}{}
			fmt.Printf("handle job: %d\n", j)
			time.Sleep(2 * time.Second)
			<-active
			wg.Done()
		}(job)
	}
	wg.Wait()
}
