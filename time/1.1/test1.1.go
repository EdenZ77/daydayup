package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	//later := now.Add(time.Hour)
	fmt.Println(now)

	timeObj, err := time.Parse(time.RFC3339, "2022-06-26T11:25:20+08:00")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(timeObj.Add(2 * 24 * time.Hour)) // 2022-10-05 mutex:25:20 +0800 CST

	before := timeObj.Add(2 * 24 * time.Hour).Before(now)
	fmt.Println(before)

	after := timeObj.Add(5 * 24 * time.Hour).After(now)
	fmt.Println(after)

}
