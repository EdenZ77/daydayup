package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func main() {
	go func() {
		fmt.Println("I`m children goroutine")
		fmt.Println(GoID())
	}()
	time.Sleep(2 * time.Second)
	fmt.Println("main over")
}

func GoID() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	// 得到id字符串
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}
