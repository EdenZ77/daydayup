package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var bStop = false

func makeStop() {
	time.Sleep(time.Second * 4)
	bStop = true
}

func producer(threadId int, wg *sync.WaitGroup, ch chan string) {
	count := 0

	for !bStop {
		time.Sleep(time.Second * 2)
		count++
		data := strconv.Itoa(threadId) + "+++++" + strconv.Itoa(count)
		fmt.Println("producer:", data)
		ch <- data
	}
	wg.Done()
}

func main() {

}
