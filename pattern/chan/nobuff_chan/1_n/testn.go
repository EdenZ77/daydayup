package main

import "time"

func main() {
	ch1 := make(chan int, 2)
	ch1 <- 5
	ch1 <- 6
	go func() {
		i, ok := <-ch1
		println(i, ok)

		i, ok = <-ch1
		println(i, ok)

		i, ok = <-ch1
		println(i, ok)

		i, ok = <-ch1
		println(i, ok)
	}()
	time.Sleep(5 * time.Second)
}
