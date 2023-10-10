package main

import "fmt"

func main() {
	buf := make([]byte, 0, 5*1024*1024)
	fmt.Println(len(buf))
	fmt.Println(cap(buf))
	//fmt.Println(buf[0]) // panic: runtime error: index out of range [0] with length 0
	fmt.Println(buf[5242880-1]) // panic: runtime error: index out of range [5242879] with length 0
	copyBuf := buf[0:cap(buf)]
	fmt.Println(len(copyBuf))
	fmt.Println(cap(copyBuf))
	fmt.Println(copyBuf[5242880-1])
	fmt.Println(copyBuf[0])
}
