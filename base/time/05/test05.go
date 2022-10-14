package main

import (
	"fmt"
	"time"
)

func main() {
	doc := time.Now().Format("2006-01-02T15:04:05.000000")
	fmt.Println(doc)

	doc2 := time.Now().Format(time.RFC3339)
	fmt.Println(doc2)
}
