package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t)              // 2022-09-19 09:54:19.0834113 +0800 CST m=+0.002659901
	fmt.Println(t.Date())       // 2022 September 19
	fmt.Println(t.Year())       // 2022
	fmt.Println(t.Month())      // September
	fmt.Println(t.ISOWeek())    // 2022 38
	fmt.Println(t.Clock())      // 9 54 19
	fmt.Println(t.Day())        // 19
	fmt.Println(t.Weekday())    // Monday
	fmt.Println(t.Hour())       // 9
	fmt.Println(t.Minute())     // 54
	fmt.Println(t.Second())     // 19
	fmt.Println(t.Nanosecond()) // 83411300
	fmt.Println(t.YearDay())    // 262
}
