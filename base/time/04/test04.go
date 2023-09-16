package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now) // 2022-07-18 21:19:59.9636 +0800 CST m=+0.000069242

	loc, _ := time.LoadLocation("UTC")
	fmt.Println(now.In(loc)) // 2022-07-18 13:19:59.9636 +0000 UTC

	loc, _ = time.LoadLocation("Europe/Berlin")
	fmt.Println(now.In(loc)) // 2022-07-18 15:19:59.9636 +0200 CEST

	loc, _ = time.LoadLocation("America/New_York")
	fmt.Println(now.In(loc)) // 2022-07-18 09:19:59.9636 -0400 EDT

	loc, _ = time.LoadLocation("Asia/Dubai")
	fmt.Println(now.In(loc)) // 2022-07-18 17:19:59.9636 +0400 +04
}
