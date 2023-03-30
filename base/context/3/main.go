package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx := context.Background()
	go func() {

		time.Sleep(5 * time.Second)
		fmt.Println(ctx)
	}()

}
